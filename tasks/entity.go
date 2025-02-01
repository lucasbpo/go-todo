package tasks

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Task struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title,omitempty"`
	Description string        `json:"description" bson:"description,omitempty"`
	Tags        []string      `json:"tags" bson:"tags,omitempty"`
	Assign      bson.ObjectID `json:"assign" bson:"assign,omitempty"`
	Done        bool          `json:"done" bson:"done,omitempty"`
}
