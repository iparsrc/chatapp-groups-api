package domain

import (
	"context"
	"time"

	"github.com/parsaakbari1209/ChatApp-groups-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(group *Group) (*Group, *utils.RestErr) {
	groupsC := db.Collection("groups")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	result, err := groupsC.InsertOne(ctx, bson.M{
		"name":            group.Name,
		"owner":           group.Owner,
		"admins":          group.Admins,
		"members":         group.Members,
		"private":         group.Private,
		"dateCreated":     group.DateCreated,
		"description":     group.Description,
		"groupUniqueName": group.GroupUniqueName,
	})
	if err != nil {
		return nil, utils.BadRequest("can't create user.")
	}
	group.ID = result.InsertedID.(primitive.ObjectID)
	return group, nil
}
