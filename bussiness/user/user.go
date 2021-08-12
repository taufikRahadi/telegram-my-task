package user

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Firstname string    `gorm:"notNull" json:"firstname"`
	Lastname  string    `json:"lastname"`
	Username  string    `gorm:"unique" json:"username"`
	CreatedAt time.Time `gorm:"autoCreatedTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
