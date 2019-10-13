package db

import (
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	context "golang.org/x/net/context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const name = "pow"

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

//MongoDB implements methods for either a live or mock session
type MongoDB interface {
	Connect() error
	Disconnect() error

	SetCollection(string) error

	InsertUser(interface{}) (string, error)
	Read(string) (interface{}, error)
	UpdateUser(string, []byte) error
	Delete(string) error
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
func (ls *LiveSession) Read(filter string) (interface{}, error) {

	var v interface{}

	i, err := primitive.ObjectIDFromHex(filter)
	if err != nil {
		return nil, err
	}

	resp := ls.Collection.FindOne(context.Background(), bson.M{"_id": i})

	if err := resp.Decode(&v); err != nil {
		return nil, err
	}

	return v, nil

}

//UpdateUser Doesn't work check https://github.com/spaceCh1mp/pow/issues/8#issue-504882414 for the issue
func (ls *LiveSession) UpdateUser(id string, b []byte) error {

	// i, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return err
	// }

	// log.Println("cow")
	// var update interface{}
	// err = bson.UnmarshalExtJSON(b, false, &update)
	// if err != nil {
	// 	return err
	// }

	// filter := bson.M{
	// 	"_id": bson.M{
	// 		"$eq": i, // check if bool field has value of 'false'
	// 	},
	// }
	// updat := bson.M{
	// 	"$set": bson.M{
	// 	},
	// }
	// log.Println("cow")
	// _, err = ls.Collection.UpdateOne(context.Background(), filter, updat) //problem

	// log.Println("cow")

	// if err != nil {
	// 	return err
	// }

	// return nil

	panic(fmt.Errorf("UpdateUser method doesn't work, Check https://github.com/spaceCh1mp/pow/issues/8#issue-504882414 for the issue "))
}

//Delete ...
func (ls *LiveSession) Delete(filter string) error {

	i, err := primitive.ObjectIDFromHex(filter)
	if err != nil {
		return err
	}

	resp, err := ls.Collection.DeleteOne(context.TODO(), bson.M{"_id": i})
	if err != nil {
		return err
	}

	if resp.DeletedCount < int64(1) {
		return errors.New("")
	}

	return nil
}

//Transactions model
type Transactions struct {
	ID     *primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UserID string              `json: "userid" bson:"user_id"`
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
	if err := ls.Connect(); err != nil {
		//TODO handle error
		return err
	}

	ls.Collection = ls.Client.Database(name).Collection(cName)
	return nil
}
