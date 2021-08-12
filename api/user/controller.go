package user

import (
	"MyTodoList/bussiness/user"
	"MyTodoList/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service user.Service
}

func NewController(service user.Service) Controller {
	return Controller{service}
}

func (ctrl *Controller) RegisterUser(c echo.Context) error {
	payload := user.CreateUserSpec{}

	c.Bind(&payload)

	fmt.Println(payload)

	result, err := ctrl.service.RegisterUser(payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{Status: false, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}
