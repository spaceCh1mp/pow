package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func connectMongo() {
	//set Client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	//Connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Couldn't connect to the mongodb server. \n Response: %v", err)
	}

	//Check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatalf("Ping failed, Check connection. \n Response: %v", err)
	}

	log.Printf("Connected to the MongoDB server...")
}

func disconnectMongo(c *mongo.Client) error {

	if err := c.Disconnect(context.TODO()); err != nil {
		log.Printf("Could not close the connection.\nErr: %v", err)
		return err
	}

	return nil
}
