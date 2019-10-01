package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//MockSession mocks the instance of the mongo server
// and dummies a response similar to a live one
type MockSession struct{}

//InsertOne ...
func (ms MockSession) InsertOne(document interface{}) (string, error) {
	return primitive.NewObjectID().Hex(), nil
}

//Connect Does nothing
func (ms MockSession) Connect() error { return nil }

//Disconnect equally Does nothing
func (ms MockSession) Disconnect() error { return nil }

//SetCollection is the worst cause it takes an argument and literally does nothing with it,
// please stop bugging me
func (ms MockSession) SetCollection(string) error { return nil }
