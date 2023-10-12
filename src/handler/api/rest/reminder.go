package rest

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/model/rest/response"
	"github.com/irdaislakhuafa/GoAttendEasy/src/middleware"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/reminder"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/customvalidator"
	"github.com/labstack/echo/v4"
)

type ReminderInterface interface {
	Create(ctx context.Context) func(c echo.Context) error
	GetList(ctx context.Context) func(c echo.Context) error
	Get(ctx context.Context) func(c echo.Context) error
	Update(ctx context.Context) func(c echo.Context) error
	Delete(ctx context.Context) func(c echo.Context) error
}

type restReminder struct {
	rest *Rest
}

func NewReminder(rest *Rest, ctx context.Context) ReminderInterface {
	reminder := &restReminder{
		rest: rest,
	}

	rest.echo.POST("/api/reminders", reminder.Create(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))
	rest.echo.GET("/api/reminders", reminder.GetList(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{}))
	rest.echo.GET("/api/reminders/:id", reminder.Get(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{}))
	rest.echo.PUT("/api/reminders", reminder.Update(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))
	rest.echo.DELETE("/api/reminders", reminder.Delete(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))

	return reminder
}

func (r *restReminder) Create(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Reminder]{}
		body := new(struct {
			IN  string `validate:"required"`
			OUT string `validate:"required"`
			Day int64  `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := r.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}
		defer tx.Rollback()

		layout := "15:04:05"
		in, err := time.Parse(layout, body.IN)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			result.Error = append(result.Error, map[string]string{"message": "IN: use time with format HH:mm:ss"})
			return c.JSON(http.StatusBadRequest, result)
		}

		out, err := time.Parse(layout, body.OUT)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			result.Error = append(result.Error, map[string]string{"message": "OUT: use time with format HH:mm:ss"})
			return c.JSON(http.StatusBadRequest, result)
		}

		if day := body.Day; day <= 0 || day > 7 {
			result.Error = append(result.Error, map[string]string{"message": fmt.Sprintf("value of day is days in weeks, thats means 1 to 7 with 1 is a Sunday, but your input is %v", day)})
			return c.JSON(http.StatusBadRequest, result)
		}

		reminder, err := tx.Reminder.Create().
			SetID(uuid.NewString()).
			SetIn(in).
			SetOut(out).
			SetDay(int(body.Day)).
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

		result.Data = reminder
		return c.JSON(http.StatusOK, result)
	}
}

func (r *restReminder) GetList(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[[]*generated.Reminder]{}
		body := new(struct {
			IsDeleted bool
		})

		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		listReminder, err := r.rest.client.Reminder.Query().
			Where(reminder.IsDeleted(body.IsDeleted)).
			All(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = listReminder
		return c.JSON(http.StatusOK, result)
	}
}

func (r *restReminder) Get(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Reminder]{}
		id := c.Param("id")

		reminder, err := r.rest.client.Reminder.Get(ctx, id)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = reminder
		return c.JSON(http.StatusOK, result)
	}
}

func (r *restReminder) Update(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Reminder]{}
		body := new(struct {
			ID  string `validate:"required"`
			IN  string `validate:"required"`
			OUT string `validate:"required"`
			Day int64  `validate:"required"`
		})

		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := r.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}
		defer tx.Rollback()

		layout := "15:04:05"
		in, err := time.Parse(layout, body.IN)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			result.Error = append(result.Error, map[string]string{"message": "IN: use time with format HH:mm:ss"})
			return c.JSON(http.StatusBadRequest, result)
		}

		out, err := time.Parse(layout, body.OUT)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			result.Error = append(result.Error, map[string]string{"message": "IN: use time with format HH:mm:ss"})
			return c.JSON(http.StatusBadRequest, result)
		}

		reminder, err := tx.Reminder.UpdateOneID(body.ID).
			SetIn(in).
			SetOut(out).
			SetUpdatedAt(time.Now()).
			SetUpdatedBy(c.Get("user_id").(string)).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = reminder
		return c.JSON(http.StatusOK, result)
	}
}

func (r *restReminder) Delete(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Reminder]{}
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

		tx, err := r.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}
		defer tx.Rollback()

		reminder, err := tx.Reminder.UpdateOneID(body.ID).
			SetIsDeleted(true).
			SetDeletedAt(time.Now()).
			SetDeletedBy(c.Get("user_id").(string)).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = reminder
		return c.JSON(http.StatusOK, result)
	}
}
