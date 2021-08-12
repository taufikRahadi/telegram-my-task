package task

import (
	"time"
)

type TaskStatus string

const (
	Done       TaskStatus = "DONE"
	InProgress TaskStatus = "IN PROGRESS"
)

type Task struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string     `json:"title" gorm:"not null"`
	Username    string     `json:"username"`
	Description string     `json:"description" gorm:"text"`
	Status      TaskStatus `json:"status" gorm:"default:IN PROGRESS"`
	DueDate     time.Time  `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
