package models

import (
	"time"
)

type User struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" binding:"required" gorm:"uniqueIndex;not null"`
	Email       string    `json:"email" binding:"required,email" gorm:"uniqueIndex;not null"`
	ActiveLogin string    `json:"-" gorm:"not null"` // El - evita que se muestre en JSON
	FullName    string    `json:"full_name"`
	Role        string    `json:"role" gorm:"default:user"`
	Active      bool      `json:"active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
