package user

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/fieldflat/abome/db"
	"github.com/fieldflat/abome/entity"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

// Service procides user's behavior
type Service struct{}

// User is alias of entity.User struct
type User entity.User

// GetAll is get all User
func (s Service) GetAll() ([]User, error) {
	log.Println("[call] service/user_service.go | func GetAll")
	db := db.GetDB()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}
	log.Println("[call] service/user_service.go | func GetAll")
	return u, nil
}

// CreateModel creates User model
func (s Service) CreateModel(c *gin.Context) (User, error) {
	log.Println("[call] service/user_service.go | func CreateModel")
	db := db.GetDB()
	var user User

	user.UserID = c.PostForm("user_id")
	user.UserName = c.PostForm("user_name")
	user.Email = c.PostForm("email")
	password := c.PostForm("password")
	passwordConfirmation := c.PostForm("password_confirmation")

	// confirm password
	if password != passwordConfirmation {
		return user, errors.New("Password doesn't match. ")
	}
	if len(password) < 6 {
		return user, errors.New("Password minimum length is 6")
	}

	user.Password, _ = passwordHash(password)
	passwordConfirmation, _ = passwordHash(passwordConfirmation)

	// validation check
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return user, errors
	}

	db.Create(&user)
	log.Println("[call end] service/user_service.go | func CreateModel")
	return user, nil
}

// GetByID is get a User
func (s Service) GetByID(id string) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

// UpdateByID is update a User
func (s Service) UpdateByID(id string, c *gin.Context) (User, error) {
	log.Println("[call] controller/user_service.go | func UpdateByID")
	db := db.GetDB()
	var u User
	log.Println(id)

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	// if err := c.BindJSON(&u); err != nil {
	// 	log.Println(err)
	// 	return u, err
	// }

	u.UserID = c.PostForm("user_id")
	u.UserName = c.PostForm("user_name")

	db.Save(&u)
	log.Println(u)
	log.Println("[call end] controller/user_service.go | func UpdateByID")

	return u, nil
}

// DeleteByID is delete a User
func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}

// GetByEmailAndPassword is a function
func (s Service) GetByEmailAndPassword(email, password string) (User, error) {
	log.Println("[call] controller/user_service.go | func GetByEmailAndPassword")
	db := db.GetDB()
	var u User
	db.Where("email = ?", email).Find(&u)
	err := passwordVerify(u.Password, password)
	if err != nil {
		log.Println("[NG] password or email are incorrect.")
		log.Println(u)
		return u, err
	}
	log.Println(u)
	log.Println("[call end] controller/user_service.go | func GetByEmailAndPassword")
	return u, err
}

//
// private function
//

// generate password hash
func passwordHash(pw string) (string, error) {
	log.Println("[call] service/user_service.go | func passwordHash")
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	log.Println("[call] service/user_service.go | func passwordHash")
	return string(hash), err
}

// check verification of password
func passwordVerify(hash, pw string) error {
	log.Println("[call] service/user_service.go | func passwordVerify")
	log.Println("[call] service/user_service.go | func passwordVerify")
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
