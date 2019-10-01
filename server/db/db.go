package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const name = "Pow"

//MongoDB implements methods for either a live or mock session
type MongoDB interface {
	Connect() error
	Disconnect() error

	SetCollection(string) error

	InsertOne(interface{}) (string, error)
}

//LiveSession handles an instance of the mongo server
//it makes live database operations
type LiveSession struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

//InsertOne ...
func (ls *LiveSession) InsertOne(document interface{}) (string, error) {

	resp, err := ls.Collection.InsertOne(context.Background(), document)
	if err != nil {
		return "", err
	}

	r := resp.InsertedID.(primitive.ObjectID)
	return r.Hex(), nil
}

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
//it returns a mongo.client object if the operation was successful or an err if it failed
func (ls *LiveSession) Connect() error {
	//set Client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	//Connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return fmt.Errorf("Couldn't connect to the mongodb server. \n Response: %v", err)
	}

	//Check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return fmt.Errorf("Ping failed, Check connection. \n Response: %v", err)
	}

	ls.Client = client
	return nil
}

//Disconnect closes the connection with the database using the mongo.client object
//returns an error if the operation failed
func (ls *LiveSession) Disconnect() error {

	if err := ls.Client.Disconnect(context.TODO()); err != nil {
		log.Printf("Could not close the connection.\nErr: %v", err)
		return err
	}

	return nil
}

//SetCollection sets the name of the collection to be queried
func (ls *LiveSession) SetCollection(cName string) error {
	if (ls.Client == &mongo.Client{}) {
		err := ls.Connect()
		if err != nil {
			return err
		}
	}

	ls.Collection = ls.Client.Database(name).Collection(cName)
	return nil
}
