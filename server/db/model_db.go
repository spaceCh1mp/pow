package db

import "go.mongodb.org/mongo-driver/bson/primitive"

//Users model
type Users struct {
	ID       *primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	FamilyID string              `json:"familyID,omitempty" bson:"family_id"`

	FirstName string `json:"firstName,omitempty" bson:"firstname"`
	LastName  string `json:"lastName,omitempty" bson:"lastname"`
	UserName  string `json:"userName,omitempty" bson:"username"`

	Email    string `json:"email,omitempty" bson:"email"`
	Password string `json:"password,omitempty" bson:"password"`
}

//Transactions model
type Transactions struct {
	ID     *primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UserID string              `json:"userid" bson:"user_id"`
	Date   string              `json:"log" bson:"log"` // change to time stamp
	Amount uint                `json:"amount" bson:"amount"`
}

//Family model
type Family struct {
	ID          *primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string              `json:"name" bson:"name"`
	Rate        float64             `json:"rate" bson:"rate"`
	OrganiserID string              `json:"orgID" bson:"organiser_id"`
	MembersID   []string            `json:"memID,omitempty" bson:"members_id"`
}
