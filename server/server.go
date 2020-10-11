package server

import (
	"log"
	"net/http"
	"os"

	user "github.com/fieldflat/abome/controller"
	"github.com/fieldflat/abome/entity"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// User model
type User entity.User

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	ctrl := user.Controller{}

	// root url
	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		log.Println(session.Get("UserID"))
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"login_name": session.Get("UserName"),
			"login_id":   session.Get("UserID"),
		})
	})

	// signup and login
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.tmpl.html", nil)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl.html", nil)
	})
	router.POST("/signup", ctrl.Create)
	router.POST("/login", ctrl.Login)
	router.GET("/logout", ctrl.Logout)

	// users
	u := router.Group("/users")
	{
		ctrl := user.Controller{}
		u.GET("", ctrl.IndexJSON)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return router
}
