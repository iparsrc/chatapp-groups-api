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

func TestRetrive(t *testing.T) {
	if db == nil {
		ConnectDB("mongodb://localhost:27017")
	}
	group, restErr := Retrive("A")
	if restErr != nil {
		t.Error(restErr.Message)
	}
	fmt.Println("    ", *group)
}

func TestUpdate(t *testing.T) {
	restErr := Update("A", "Python-Gang", "This is for python gangs.", false)
	if restErr != nil {
		t.Error(restErr.Message)
	}
	restErr = Update("A", "Python-Gang", "This is for python gangs.", false)
	if restErr == nil {
		t.Error("up-to-date group must give a rest error not a nil rest error.")
	}
	group, restErr := Retrive("A")
	if restErr != nil {
		t.Error(restErr.Message)
	}
	fmt.Println("    ", *group)
}

func TestDelete(t *testing.T) { // Complete.
	if db == nil {
		ConnectDB("mongodb://localhost:27017")
	}
	restErr := Delete("A")
	if restErr != nil {
		t.Error(restErr.Message)
	}
	restErr = Delete("A")
	if restErr == nil {
		t.Error("Deleting a group that doesn't exists must not give a nil err.")
	}
	restErr = Delete("B")
	if restErr != nil {
		t.Error(restErr.Message)
	}
}
