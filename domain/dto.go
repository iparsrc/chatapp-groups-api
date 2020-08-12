package domain

type Group struct {
	// ID must be a unique name reqresenting group.
	ID          string   `json:"id" bson:"_id"`
	Name        string   `json:"name" bson:"name"`
	Owner       string   `json:"owner" bson:"owner"`
	Admins      []string `json:"admins" bson:"admins"`
	Members     []string `json:"members" bson:"members"`
	Private     bool     `json:"private" bson:"private"`
	DateCreated int64    `json:"dateCreated" bson:"dateCreated"`
	Description string   `json:"description" bson:"description"`
}
