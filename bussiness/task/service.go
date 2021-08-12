package task

import (
	"fmt"
	"time"
)

type CreateTaskSpec struct {
	Title       string     `json:"title"`
	Username    string     `json:"username"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	DueDate     string     `json:"due_date"`
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateTask(payload *CreateTaskSpec) (Task, error) {
	fmt.Println(payload)
	task := Task{
		Title:       payload.Title,
		Description: payload.Description,
		Username:    payload.Username,
		Status:      payload.Status,
		DueDate:     time.Now(),
	}
	return s.repository.CreateTask(task)
}

func (s *service) GetAllMyTask(username string) ([]Task, error) {
	return s.repository.GetAllMyTask(username)
}
