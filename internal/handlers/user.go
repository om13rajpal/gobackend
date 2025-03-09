package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/om13rajpal/gobackend/internal/database"
	"github.com/om13rajpal/gobackend/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	cursor, err := database.UserCollection.Find(ctx, bson.M{})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "could not fetch users",
		})
		return
	}
	var users []bson.M

	for cursor.Next(ctx) {
		var document bson.M
		err := cursor.Decode(&document)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, document)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "users fetched successfully",
		"users":   users,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	objId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	user := database.UserCollection.FindOne(ctx, bson.M{
		"_id": objId,
	})

	var fetchedUser models.User

	err = user.Decode(&fetchedUser)

	if err != nil {
		log.Fatal("Could not decode data\n", err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "user found",
		"user":    fetchedUser,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	objId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	deleteUser := database.UserCollection.FindOneAndDelete(ctx, bson.M{
		"_id": objId,
	})

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":      true,
		"message":     "User deleted successfully",
		"deletedUser": deleteUser,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updatedValues models.User

	c.BindJSON(&updatedValues)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	objId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	var updatedUser models.User
	err = database.UserCollection.FindOneAndUpdate(ctx, bson.M{
		"_id": objId,
	}, bson.M{"$set": updatedValues}).Decode(&updatedUser)
	
	if err != nil {
		log.Fatal(err)
	}	

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":      true,
		"message":     "user updated successfully",
		"updatedUser": updatedUser,
	})
}
