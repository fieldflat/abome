package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User Model
type User struct {
	gorm.Model
	UserID   string
	UserName string
	Email    string
	Password string
}

// generate password hash
func passwordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// check verification of password
func passwordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
