package rest

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/GoAttendEasy/src/entity"
	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/model/rest/response"
	"github.com/irdaislakhuafa/GoAttendEasy/src/middleware"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/attendance"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/predicate"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/user"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/customvalidator"
	"github.com/labstack/echo/v4"
)

type AttendanceInterface interface {
	In(ctx context.Context) func(c echo.Context) error
	GetList(ctx context.Context) func(c echo.Context) error
	Get(ctx context.Context) func(c echo.Context) error
	Out(ctx context.Context) func(c echo.Context) error
	Delete(ctx context.Context) func(c echo.Context) error
}

type restAttendance struct {
	rest *Rest
}

func NewAttendance(rest *Rest, ctx context.Context) AttendanceInterface {
	attendance := &restAttendance{
		rest: rest,
	}

	rest.echo.POST("/api/attendances/in", attendance.In(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin", "employee"}}))
	rest.echo.GET("/api/attendances", attendance.GetList(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{}}))
	rest.echo.GET("/api/attendances/:id", attendance.Get(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin", "employee"}}))
	rest.echo.PUT("/api/attendances/out", attendance.Out(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))
	rest.echo.DELETE("/api/attendances", attendance.Delete(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))

	return attendance
}

func (a *restAttendance) In(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Attendance]{}
		body := new(struct {
			UserID    string `validate:"required"`
			IN        string `validate:"required"`
			IsPresent bool   `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		layout := "15:04:05"
		in, err := time.Parse(layout, body.IN)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			result.Error = append(result.Error, map[string]string{"message": "IN: use time with format HH:mm:ss"})
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := a.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}
		defer tx.Rollback()

		now := time.Now()
		conditions := []predicate.Attendance{
			attendance.CreatedAtLTE(now),
			attendance.CreatedAtGTE(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)),
			attendance.UserID(body.UserID),
		}
		isExist, err := tx.Attendance.Query().Where(conditions...).Exist(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		} else if isExist {
			result.Error = append(result.Error, map[string]string{"message": "You have sent the attendance list today!"})
			return c.JSON(http.StatusBadRequest, result)
		}

		user, err := tx.User.Query().Where(user.ID(body.UserID)).Select(user.FieldIsDeleted).First(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if user.IsDeleted {
			result.Error = append(result.Error, map[string]string{"message": fmt.Sprintf("user with %v is deleted", body.UserID)})
			return c.JSON(http.StatusBadRequest, result)
		}

		attendance, err := tx.Attendance.Create().
			SetID(uuid.NewString()).
			SetIn(in).
			SetUserID(user.ID).
			SetIsPresent(body.IsPresent).
			SetCreatedAt(time.Now()).
			SetCreatedBy(c.Get("user_id").(string)).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Data = attendance
		return c.JSON(http.StatusOK, result)
	}
}

func (a *restAttendance) GetList(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[[]*generated.Attendance]{}
		body := new(struct {
			IsDeleted bool
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			return c.JSON(http.StatusBadRequest, result)
		}
		now := time.Now()
		conditions := []predicate.Attendance{
			attendance.IsDeleted(body.IsDeleted),
		}

		if c, isOk := c.Get("claims").(*entity.Claims); isOk {
			if strings.EqualFold(c.RoleName, "employee") {
				conditions = append(conditions,
					attendance.CreatedAtLTE(now),
					attendance.CreatedAtGTE(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)),
					attendance.UserID(c.UserID),
				)
			}
		}

		listAttendance, err := a.rest.client.Attendance.Query().Where(conditions...).All(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = listAttendance
		return c.JSON(http.StatusOK, result)
	}
}

func (a *restAttendance) Get(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Attendance]{}
		id := c.Param("id")

		attendance, err := a.rest.client.Attendance.Get(ctx, id)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = attendance
		return c.JSON(http.StatusOK, result)
	}
}

func (a *restAttendance) Out(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Attendance]{}
		body := new(struct {
			ID  string `validate:"required"`
			OUT string `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := a.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}
		defer tx.Rollback()

		layout := "15:04:05"
		out, err := time.Parse(layout, body.OUT)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			result.Error = append(result.Error, map[string]string{"message": fmt.Sprintf("OUT: use time with format HH:mm:ss")})
			return c.JSON(http.StatusBadRequest, result)
		}

		attendance, err := tx.Attendance.Get(ctx, body.ID)
		if err != nil {
		}

		attendance, err = tx.Attendance.UpdateOneID(body.ID).
			SetOut(out).
			SetUpdatedAt(time.Now()).
			SetUpdatedBy(c.Get("user_id").(string)).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Data = attendance
		return c.JSON(http.StatusOK, result)
	}
}

func (a *restAttendance) Delete(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Attendance]{}
		body := new(struct {
			ID string `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := a.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}
		defer tx.Rollback()

		attendance, err := tx.Attendance.UpdateOneID(body.ID).
			SetIsDeleted(true).
			SetDeletedAt(time.Now()).
			SetDeletedBy(c.Get("user_id").(string)).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Data = attendance
		return c.JSON(http.StatusOK, result)
	}
}
