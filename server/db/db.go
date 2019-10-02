package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const name = "Pow"

//Users model
type Users struct {
	FamilyID string `json:"family_id" bson:"family_id"`

	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	UserName  string `json:"userName" bson:"userName"`

	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`

	Log Transactions `json:"log" bson:"log"`
}

//MongoDB implements methods for either a live or mock session
type MongoDB interface {
	Connect() error
	Disconnect() error

	SetCollection(string) error

	InsertUser(interface{}) (string, error)
	ReadUser(bson.M) (Users, error)
}

//LiveSession handles an instance of the mongo server
//it makes live database operations
type LiveSession struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

//InsertUser ...
func (ls *LiveSession) InsertUser(document interface{}) (string, error) {
	resp, err := ls.Collection.InsertOne(context.Background(), document)
	if err != nil {
		return "", err
	}

	r := resp.InsertedID.(primitive.ObjectID)
	return r.Hex(), nil
}

//ReadUser ...
func (ls *LiveSession) ReadUser(filter bson.M) (Users, error) {
	var u Users
	resp := ls.Collection.FindOne(context.Background(), filter)

	if err := resp.Decode(&u); err != nil {
		return Users{}, err
	}
	return u, nil
}

//Transactions model
type Transactions struct {
	Date   time.Time `json:"log" bson:"log"` // change to time stamp
	Amount uint      `json:"amount" bson:"amount"`
}

//Family model
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
//it returns an error if the client is been disconnected
func (ls *LiveSession) SetCollection(cName string) error {
	if ls.Client == nil {
		return mongo.ErrClientDisconnected
	}

	ls.Collection = ls.Client.Database(name).Collection(cName)
	return nil
}
