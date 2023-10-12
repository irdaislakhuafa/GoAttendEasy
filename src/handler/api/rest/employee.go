package rest

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/model/rest/response"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/employee"
	"github.com/labstack/echo/v4"
)

type EmployeeInterface interface {
	Create(ctx context.Context) func(c echo.Context) error
	GetList(ctx context.Context) func(c echo.Context) error
	Get(ctx context.Context) func(c echo.Context) error
	Update(ctx context.Context) func(c echo.Context) error
	Delete(ctx context.Context) func(c echo.Context) error
}
type restEmployee struct {
	rest *Rest
}

func NewEmployee(rest *Rest, ctx context.Context) EmployeeInterface {
	employee := &restEmployee{
		rest: rest,
	}

	// TODO: use jwt middleware
	// TODO: added query parameter for get one by id
	rest.echo.POST("/api/employees", employee.Create(ctx))
	rest.echo.GET("/api/employees", employee.GetList(ctx))
	rest.echo.PUT("/api/employees", employee.Update(ctx))
	rest.echo.DELETE("/api/employees", employee.Delete(ctx))
	return employee
}

func (e *restEmployee) Create(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Employee]{}
		body := new(struct {
			UserID     string `validate:"required"`
			DivisionID string `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := e.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}
		defer tx.Rollback()

		employee, err := tx.Employee.Create().
			SetUserID(body.UserID).
			SetDivisionID(body.DivisionID).
			SetCreatedAt(time.Now()).
			SetCreatedBy(c.Get("user_id").(string)).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Data = employee
		return c.JSON(http.StatusOK, result)
	}
}

func (e *restEmployee) GetList(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[[]*generated.Employee]{}
		body := new(struct {
			IsDeleted bool `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		listEmployee, err := e.rest.client.Employee.Query().Where(employee.IsDeleted(body.IsDeleted)).All(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = listEmployee
		return c.JSON(http.StatusOK, result)
	}
}

func (e *restEmployee) Get(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Employee]{}
		body := new(struct {
			ID string `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		employee, err := e.rest.client.Employee.Get(ctx, body.ID)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = employee
		return c.JSON(http.StatusOK, result)
	}
}

func (e *restEmployee) Update(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Employee]{}
		body := new(struct {
			ID         string `validate:"required"`
			UserID     string `validate:"required"`
			DivisionID string `validate:"required"`
		})

		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := e.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}
		defer tx.Rollback()

		employee, err := tx.Employee.UpdateOneID(body.ID).
			SetUserID(body.UserID).
			SetDivisionID(body.DivisionID).
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

		result.Data = employee
		return c.JSON(http.StatusOK, result)
	}
}

func (e *restEmployee) Delete(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Employee]{}
		body := new(struct {
			ID string `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := e.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}
		defer tx.Rollback()

		employee, err := tx.Employee.UpdateOneID(body.ID).
			SetDeletedAt(time.Now()).
			SetDeletedBy(c.Get("'user_id").(string)).
			SetIsDeleted(true).
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

		result.Data = employee
		return c.JSON(http.StatusOK, result)
	}
}
