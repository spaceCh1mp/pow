package db

import (
	"encoding/json"
	"fmt"
	"log"

	context "golang.org/x/net/context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const name = "pow"

func parse(b []byte, opt string) (interface{}, bson.D, error) {

	var (
		m       map[string]string
		bsonMap bson.D
		i       interface{}
	)

	//unmarshal data to map (m)
	//return err if operation fails
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, nil, err
	}

	//Loop through map
	for k, v := range m {

		i = v

		//get Document ID
		//return err if there was a problem converting the string to an ObjectId type
		if k == "id" {

			i, err = primitive.ObjectIDFromHex(string(v))
			if err != nil {
				return nil, nil, ErrInvalidHex
			}

			//check opt to know what operation called parse
			// continue if it's the update operation
			//set key to "_id" if it's the read operation.
			if opt == "u" {
				continue
			} else if opt == "r" {
				k = "_id"
			}

		}

		//add entries
		bsonMap = append(bsonMap, bson.E{Key: k, Value: i})

	}

	return m["id"], bsonMap, nil

}

//MongoDB implements methods for either a live or mock session
type MongoDB interface {
	Connect() error
	Disconnect() error

	SetCollection(string) error

	Insert(interface{}) (string, error)
	Read([]byte, interface{}) (interface{}, error)
	Update([]byte) error
	UpdateArray([]byte, string) error
	Delete(string) error
}

//LiveSession handles an instance of the mongo server
//it makes live database operations
type LiveSession struct {
	Client     *mongo.Client
	Collection *mongo.Collection
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

//MockSession mocks the instance of the mongo server
// and dummies a response similar to a live one
type MockSession struct {
	C bool
}

//Connect pings an empty client and returns an error if mockSession is set to false else
//it just returns nil.
func (ms MockSession) Connect() error {

	if !ms.C {
		var c = mongo.Client{}
		return c.Ping(context.TODO(), nil)
	}

	return nil
}

//Disconnect equally Does nothing
func (ms MockSession) Disconnect() error { return nil }

//SetCollection is the worst cause it takes an argument and literally does nothing with it,
// please stop bugging me
func (ms MockSession) SetCollection(string) error {
	if !ms.C {
		return mongo.ErrClientDisconnected
	}
	return nil
}
