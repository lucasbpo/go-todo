package tags

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Tag struct {
	ID   bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name,omitempty"`
}
