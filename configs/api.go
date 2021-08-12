package configs

import (
	"MyTodoList/api"

	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitApi(db *gorm.DB) {
	e := api.InitRoutes(db)

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8080"))
}
