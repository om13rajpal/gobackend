package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/om13rajpal/gobackend/internal/database"
	"github.com/om13rajpal/gobackend/internal/models"
	"github.com/om13rajpal/gobackend/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Signup(c *gin.Context) {
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err = database.UserCollection.InsertOne(ctx, user)

	if err != nil {
		log.Fatal("Could not save the data")
	}

	err = utils.SendMail(user.Email, "Registered ;)", "<h1>Hello "+user.Username+"!</h1>\nYour account has been successfully created")

	if err != nil {
		fmt.Println("Email not sent", err)
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		log.Fatal("Could not generate a token")
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "user inserted successfully",
		"token": token,
		"user":    user,
	})
}

func Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var fetchedUser models.User
	database.UserCollection.FindOne(ctx, bson.M{
		"username": user.Username,
	}).Decode(&fetchedUser)

	validPassword := utils.CheckPassword(user.Password, fetchedUser.Password)

	if !validPassword {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid password",
		})
		return
	}

	token, err := utils.GenerateToken(fetchedUser)

	if err != nil {
		log.Fatal("Could not generate a login token", err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": true,
		"message": "user logged in successfully",
		"token":  token,
		"user":   fetchedUser,
	})
}
