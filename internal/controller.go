package internal

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Create(c echo.Context) error {
	var req CreateRequest
	c.Bind(&req)
	res, err := controller.service.Create(req)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusCreated, res)
}

func (controller *Controller) GetByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}
