package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HandleHome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "backend is up and working",
	})
}
