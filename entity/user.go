package entity

import "github.com/jinzhu/gorm"

// User Model
type User struct {
	gorm.Model
	UserID   string
	UserName string
	Email    string
	Password string
}
