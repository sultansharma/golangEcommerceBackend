package models

import (
	"time"
	_ "time"
)

type User struct {
	Id          uint      `json:"id"`
	UserType    string    `json:"user_type"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email" binding:"required,email" gorm:"unique;not null"`
	Phone       string    `json:"phone"`
	Password    []byte    `json:"-"`
	DeviceToken string    `json:"device_token"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
