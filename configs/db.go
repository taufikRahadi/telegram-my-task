package configs

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"MyTodoList/bussiness/task"
	"MyTodoList/bussiness/user"
)

var db *gorm.DB
var err error

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second,   // Slow SQL threshold
		LogLevel:                  logger.Silent, // Log level
		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,         // Disable color
	},
)

func InitDB() {
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "sql4430487:FzTVSLFwe4@tcp(sql4.freemysqlhosting.net:3306)/sql4430487?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(`DB Connection Error`)
	}

	// automigrate here
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&task.Task{})

}

func DbManager() *gorm.DB {
	return db
}
