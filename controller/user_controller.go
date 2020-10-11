package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fieldflat/abome/db"
	"github.com/fieldflat/abome/entity"
	user "github.com/fieldflat/abome/service"
	"github.com/gin-contrib/sessions"
)

// Controller is user controlller
type Controller struct{}

// User is alias of entity.User struct
type User entity.User

// Index action: GET /users
func (pc Controller) IndexJSON(c *gin.Context) {
	var s user.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /users
func (pc Controller) Create(c *gin.Context) {
	log.Println("[call] controller/user_controller.go | func Create")
	var s user.Service
	user, err := s.CreateModel(c)

	if err != nil {
		c.HTML(http.StatusOK, "signup.tmpl.html", gin.H{
			"err": err,
		})
		log.Println(err)
	} else {
		session := sessions.Default(c)
		if session.Get("UserID") != user.UserID {
			createSession(c, user)
		}
		log.Printf("%v\n", session.Get("UserID"))
		c.HTML(http.StatusTemporaryRedirect, "index.tmpl.html", gin.H{
			"login_name": session.Get("UserName"),
			"login_id":   session.Get("UserID"),
		})
	}
	log.Println("[call end] controller/user_controller.go | func Create")
}

// Show action: GET /users/:id
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /users/:id
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /users/:id
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

//
// private function
//

// userGet
func userGet(username, pass string) (int, User) {
	db := db.GetDB()
	var user User
	cnt := 0
	db.Where("user_name = ?", username).Find(&user).Count(&cnt)

	return cnt, user
}

// createSession
func createSession(c *gin.Context, user user.User) {
	session := sessions.Default(c)
	session.Set("UserID", user.UserID)
	session.Set("UserName", user.UserName)
	session.Save()
}
