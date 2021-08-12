package api

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	taskC "MyTodoList/api/task"
	userC "MyTodoList/api/user"

	"MyTodoList/bussiness/task"
	"MyTodoList/bussiness/user"
)

func InitRoutes(db *gorm.DB) *echo.Echo {
	e := echo.New()

	userservice := user.NewService(user.NewRepository(db))
	taskService := task.NewService(task.NewRepository(db))

	usercontroller := userC.NewController(userservice)
	taskController := taskC.NewController(taskService)

	// register routes here
	e.POST("/user", usercontroller.RegisterUser)

	e.POST("/task", taskController.CreateTask)
	e.GET("/tasks/:username", taskController.GetAllMyTask)

	return e
}
