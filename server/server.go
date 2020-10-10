package server

import (
	"log"
	"net/http"
	"os"

	"github.com/fieldflat/abome/entity"
	"github.com/fieldflat/abome/user"
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
	ctrl := user.Controller{}

	// root url
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	// signup and login
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.tmpl.html", nil)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl.html", nil)
	})
	router.POST("/signup", ctrl.Create)
	router.POST("/login", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		log.Println(c)
	})

	// u := r.Group("/users")
	// {
	// 	ctrl := user.Controller{}
	// 	u.GET("", ctrl.Index)
	// 	u.GET("/:id", ctrl.Show)
	// 	u.POST("", ctrl.Create)
	// 	u.PUT("/:id", ctrl.Update)
	// 	u.DELETE("/:id", ctrl.Delete)
	// }

	return router
}
