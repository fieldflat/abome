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
	router.LoadHTMLGlob("templates/*/*")
	router.Static("/static", "static")
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	userCtrl := user.Controller{}

	// root url
	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		log.Println(session.Get("UserID"))
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"login_name": session.Get("UserName"),
			"login_uid":  session.Get("UserID"),
		})
	})
	// signup and login
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.tmpl.html", nil)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl.html", nil)
	})
	router.POST("/signup", userCtrl.Create)
	router.POST("/login", userCtrl.Login)
	router.GET("/logout", userCtrl.Logout)

	// users
	u := router.Group("/users")
	{
		userCtrl := user.Controller{}
		u.GET("/", userCtrl.Index)
		u.GET("/:id", userCtrl.Show)
		u.PUT("/:id", userCtrl.Update)
		u.DELETE("/:id", userCtrl.Delete)
	}

	router.GET("/user/edit/:id", userCtrl.Edit)
	router.POST("/user/update/:id", userCtrl.Update)

	return router
}
