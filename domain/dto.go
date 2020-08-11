package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Owner           string             `json:"owner" bson:"owner"`
	Admins          []string           `json:"admins" bson:"admins"`
	Members         []string           `json:"members" bson:"members"`
	Private         bool               `json:"private" bson:"private"`
	DateCreated     int64              `json:"dateCreated" bson:"dateCreated"`
	Description     string             `json:"description" bson:"description"`
	GroupUniqueName string             `json:"groupUniqueName" bson:"groupUniqueName"`
}
