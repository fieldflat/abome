package entity

import "github.com/jinzhu/gorm"

// User Model
type User struct {
	gorm.Model
	UserID   string `json:"user_id" validate:"required,min=1,max=99"`
	UserName string `json:"user_name" validate:"required,min=1,max=99"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
