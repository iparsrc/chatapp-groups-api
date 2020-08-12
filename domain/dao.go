package domain

import (
	"context"
	"time"

	"github.com/parsaakbari1209/ChatApp-groups-api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Create(group *Group) (*Group, *utils.RestErr) {
	groupsC := db.Collection("groups")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := groupsC.InsertOne(ctx, bson.M{
		"_id":         group.ID,
		"name":        group.Name,
		"owner":       group.Owner,
		"admins":      group.Admins,
		"members":     group.Members,
		"private":     group.Private,
		"dateCreated": group.DateCreated,
		"description": group.Description,
	})
	if err != nil {
		return nil, utils.BadRequest("can't create user.")
	}
	return group, nil
}
