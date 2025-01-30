package users

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID       bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name,omitempty"`
	Email    string        `json:"email" bson:"email,omitempty"`
	Password string        `json:"password" bson:"password,omitempty"`
}
