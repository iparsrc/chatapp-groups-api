package domain

import (
	"fmt"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	ConnectDB("mongodb://localhost:27017")
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
	ConnectDB("mongodb://localhost:27017")
	group, restErr := Retrive("A")
	if restErr != nil {
		t.Error(restErr.Message)
	}
	fmt.Println("    ", *group)
}
