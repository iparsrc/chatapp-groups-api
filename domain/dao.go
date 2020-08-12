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

func Retrive(id string) (*Group, *utils.RestErr) {
	groupsC := db.Collection("groups")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var group Group
	if err := groupsC.FindOne(ctx, bson.M{"_id": id}).Decode(&group); err != nil {
		return nil, utils.NotFound("group not found.")
	}
	return &group, nil
}

func Delete(id string) *utils.RestErr {
	groupsC := db.Collection("groups")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := groupsC.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return utils.InternalServerErr("can't operate delete functionality.")
	}
	if res.DeletedCount == 0 {
		return utils.NotFound("group not found.")
	}
	return nil
}

func Update(id, name, description string, private bool) *utils.RestErr {
	groupsC := db.Collection("groups")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	update := bson.M{
		"$set": bson.M{
			"name":        name,
			"private":     private,
			"description": description,
		},
	}
	res, err := groupsC.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return utils.InternalServerErr("can't operate update functionality.")
	}
	if res.MatchedCount == 0 {
		return utils.NotFound("group not found.")
	}
	if res.ModifiedCount == 0 {
		return utils.BadRequest("nothing to update, group is already up-to-date.")
	}
	return nil
}

func AddAdmin(groupID, userID string) *utils.RestErr {
	groupsC := db.Collection("groups")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	update := bson.M{
		"$addToSet": bson.M{
			"admins": userID,
		},
	}
	res, err := groupsC.UpdateOne(ctx, bson.M{"_id": groupID}, update)
	if err != nil {
		return utils.InternalServerErr("can't operate add admin functionality.")
	}
	if res.MatchedCount == 0 {
		return utils.NotFound("group not found")
	}
	if res.ModifiedCount == 0 {
		return utils.BadRequest("user is already an admin.")
	}
	return nil
}

func AddMember(groupID, userID string) *utils.RestErr {
	groupsC := db.Collection("groups")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	update := bson.M{
		"$addToSet": bson.M{
			"members": userID,
		},
	}
	res, err := groupsC.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return utils.InternalServerErr("can't operate add member functionality.")
	}
	if res.MatchedCount == 0 {
		return utils.NotFound("group not found.")
	}
	if res.ModifiedCount == 0 {
		return utils.BadRequest("user is already a member")
	}
	return nil
}
