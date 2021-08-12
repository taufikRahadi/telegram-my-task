package task

import (
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) CreateTask(payload Task) (Task, error) {
	task := Task{Title: payload.Title, Username: payload.Username, Description: payload.Description, Status: payload.Status, DueDate: payload.DueDate}
	result := r.db.Create(&task)

	return task, result.Error
}

func (r *repository) GetAllMyTask(username string) ([]Task, error) {
	myTasks := []Task{}

	result := r.db.Find(&myTasks, "username = ?", username)

	fmt.Println(result)

	return myTasks, result.Error
}
