package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	context "golang.org/x/net/context"
)

//Insert takes a single document and inserts it into a collection
//An error is returned if the operation fails
func (ls *LiveSession) Insert(document interface{}) (string, error) {

	//ToDo ensure data is in line with model to be inserted
	resp, err := ls.Collection.InsertOne(context.TODO(), document)
	if err != nil {
		return "", err
	}

	r := resp.InsertedID.(primitive.ObjectID)
	return r.Hex(), nil

}

//Insert (mock) dummies the insert operation. Used for testing
func (ms MockSession) Insert(document interface{}) (string, error) {

	return primitive.NewObjectID().Hex(), nil

}
