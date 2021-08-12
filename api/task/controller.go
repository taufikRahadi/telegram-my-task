package task

import (
	"MyTodoList/bussiness/task"
	"MyTodoList/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service task.Service
}

func NewController(service task.Service) Controller {
	return Controller{
		service,
	}
}

func (ctrl *Controller) CreateTask(c echo.Context) error {
	payload := task.CreateTaskSpec{}

	c.Bind(&payload)

	fmt.Println(payload)
	result, err := ctrl.service.CreateTask(&payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{Status: false, Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func (ctrl *Controller) GetAllMyTask(c echo.Context) error {
	username := c.Param("username")
	tasks, err := ctrl.service.GetAllMyTask(username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{Status: false, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, tasks)
}
