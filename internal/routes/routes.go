package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/om13rajpal/gobackend/internal/handlers"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/", handlers.HandleHome)
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	r.DELETE("/users/:id", handlers.DeleteUser)

	r.PATCH("/users/:id", handlers.UpdateUser)

	return r
}