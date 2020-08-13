package domain

import (
	"fmt"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	if db == nil {
		ConnectDB("mongodb://localhost:27017")
	}
	groupOne := Group{
		ID:          "A",
		Name:        "LetsPair",
		Owner:       "1",
		Members:     []string{"1"},
		Admins:      []string{},
		Private:     true,
		DateCreated: time.Now().Unix(),
		Description: "Hey guys lets pair together.",
	}
	groupTwo := Group{
		ID:          "B",
		Name:        "GoLang-Gang",
		Owner:       "2",
		Members:     []string{"2", "1"},
		Admins:      []string{"1"},
		Private:     false,
		DateCreated: time.Now().Unix(),
		Description: "This group is for golang gang yeah.",
	}
	group, restErr := Create(&groupOne)
	if restErr != nil {
		t.Error(restErr.Message)
	}
	fmt.Println("    ", *group)
	group, restErr = Create(&groupTwo)
	if restErr != nil {
		t.Error(restErr.Message)
	}
	fmt.Println("    ", *group)
}

func TestRetrive(t *testing.T) { // Complete.
	if db == nil {
		ConnectDB("mongodb://localhost:27017")
	}
	// Try to retrive a group that exists.
	group, restErr := Retrive("A")
	if restErr != nil {
		t.Error(restErr.Message)
	}
	if group == nil {
		t.Error("Retriving a group that exist, must not give a nil group.")
	}
	fmt.Println("    ", *group)
	// Try to retrive a group that doesn't exist.
	group, restErr = Retrive("C")
	if restErr == nil {
		t.Error("Retriving a group that doesn't exsit, must not give a nil err.")
	}
	if group != nil {
		t.Error("Retriving a group that doesn't exist, must give a nil group.")
	}
}

func TestUpdate(t *testing.T) { // Complete.
	if db == nil {
		ConnectDB("mongodb://localhost:27017")
	}
	// Try to update a group that exists and is not up-to-date.
	if restErr := Update("A", "Python-Gang", "This is for python gang.", false); restErr != nil {
		t.Error(restErr.Message)
	}
	// Try to update a group that is already up-to-date(nothing is chagned).
	if restErr := Update("A", "Python-Gang", "This is for python gang.", false); restErr == nil {
		t.Error("up-to-date group must give a rest error not a nil rest error.")
	}
	// Try to update a group that doesn't exsit.
	if restErr := Update("C", "New-Group", "This is a new group.", false); restErr == nil {
		t.Error("Updating a group that doesn't exist, must not give a nil err.")
	}
}

func TestAddAdmin(t *testing.T) { // Complete.
	if db == nil {
		ConnectDB("mongodb://localhost:27017")
	}
	if restErr := AddAdmin("A", "1"); restErr != nil { // Try to add a new admin to the group.
		t.Error(restErr.Message)
	}
	if restErr := AddAdmin("A", "1"); restErr == nil { // Try to add an admin that is already an admin.
		t.Error("Adding admin that already is an admin of the group, must not give a nil err.")
	}
	if restErr := AddAdmin("C", "1"); restErr == nil { // Try to add an admin to a group that doesn't exist.
		t.Error("Adding admin to a group that doesn't exist, must not give a nil err.")
	}
}

func TestDelete(t *testing.T) { // Complete.
	if db == nil {
		ConnectDB("mongodb://localhost:27017")
	}
	if restErr := Delete("A"); restErr != nil { // Try to delete a group that exists.
		t.Error(restErr.Message)
	}
	if restErr := Delete("A"); restErr == nil { // Try to delete a group that doesn't exist.
		t.Error("Deleting a group that doesn't exists, must not give a nil err.")
	}
	if restErr := Delete("B"); restErr != nil { // Try to delete a group that exist.
		t.Error(restErr.Message)
	}
}
