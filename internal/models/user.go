package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty" validate:"email"`
	Username string `json:"username" bson:"username" validate:"min=3,max=20"`
	Password string `json:"password" bson:"password" validate:"min=6,max=30"`
}