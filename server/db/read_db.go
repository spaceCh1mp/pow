package db

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	context "golang.org/x/net/context"
)

var v interface{}

func parseRead(b []byte) (bson.D, error) {

	_, bm, err := parse(b, "r")

	return bm, err
}

//Read gets the record from a collection, where filter is the string representation of
//the document ObjectID.
//Read returns an error if the ID is not valid, ErrNoMatchedDocument is returned if there are no results
func (ls *LiveSession) Read(b []byte) (interface{}, error) {

	filter, err := parseRead(b)
	if err != nil {
		return nil, err
	}

	resp := ls.Collection.FindOne(context.Background(), filter)

	if err := resp.Decode(&v); err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, ErrNoMatchedDocument
		}

		return nil, err
	}

	return v, nil

}

//Read (mock) dummies the read operation. Used for testing
func (ms MockSession) Read(b []byte) (interface{}, error) {

	var filter map[string]string

	_ = json.Unmarshal(b, &filter)

	//dummy reply for when the Find operation fails
	//possible cause would be a wrong ObjectId
	_, err := primitive.ObjectIDFromHex(filter["id"])
	if err != nil {

		return nil, ErrInvalidHex

	}

	//dummy reply for when no document is returned
	if filter["id"] == "4af9f070a466655aeb230cbd" {
		return nil, ErrNoMatchedDocument
	}

	err = json.Unmarshal(b, &v)

	if err != nil {
		return nil, err
	}

	return v, nil

}
