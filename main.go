package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fieldflat/abome/db"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

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
	router.POST("/signup", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		var user User
		user.UserID = c.PostForm("user_id")
		user.UserName = c.PostForm("user_name")
		user.Email = c.PostForm("email")
		user.Password = c.PostForm("password")
		log.Println(user)
	})
	router.POST("/login", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		log.Println(c)
	})

	db.Init()
	router.Run(":" + port)
	db.Close()
}
