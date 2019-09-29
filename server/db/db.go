package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Name exports the name of the database
const Name = "Pow"

//User field
type User struct {
	FamilyID string `json:"family_id" bson:"family_id"`

	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`

	Log Transaction `json:"log" bson:"log"`
}

//Transaction field
type Transaction struct {
	Date   time.Time `json:"log" bson:"log"` // change to time stamp
	Amount uint      `json:"amount" bson:"amount"`
}

//Family field
type Family struct {
	Name        string   `json:"name" bson:"name"`
	OrganiserID string   `json:"organiser_id" bson:"organiser_id"`
	MembersID   []string `json:"members_id" bson:"members_id"`
}

//Connect makes a new connection with the database
//it returns a mongo.client ibject if the operation was successful or an err if it failed
func Connect() (*mongo.Client, error) {
	//set Client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	//Connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to the mongodb server. \n Response: %v", err)
	}

	//Check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, fmt.Errorf("Ping failed, Check connection. \n Response: %v", err)
	}

	return client, nil
}

//Disconnect closes the connection with the database using the mongo.client object
//returns an error if the operation failed
func Disconnect(c *mongo.Client) error {

	if err := c.Disconnect(context.TODO()); err != nil {
		log.Printf("Could not close the connection.\nErr: %v", err)
		return err
	}

	return nil
}
