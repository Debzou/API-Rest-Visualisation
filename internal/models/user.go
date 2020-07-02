package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

)


type User struct {
	ID       primitive.ObjectID `form:"_id" json:"_id"`
	Username string             `form:"username" json:"username"`
	Password string             `form:"password" json:"password"`
	Status   string             `form:"status" json:"status"`
}

