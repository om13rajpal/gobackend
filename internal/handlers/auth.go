package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/om13rajpal/gobackend/internal/database"
	"github.com/om13rajpal/gobackend/internal/models"
	"github.com/om13rajpal/gobackend/utils"
)

func Signup(c *gin.Context){
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal("Could not read the body")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal("Could not hash the password")
	}

	user.Password = hashedPassword

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	_, err = database.UserCollection.InsertOne(ctx, user)

	if err != nil {
		log.Fatal("Could not save the data")
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": true,
		"message": "user inserted successfully",
		"user": user,
	})
}