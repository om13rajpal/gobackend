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
	r.POST("/signup", handlers.Signup)

	return r
}