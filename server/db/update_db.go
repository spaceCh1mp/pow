package db

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	context "golang.org/x/net/context"
)

func parseUpdate(b []byte) (*primitive.ObjectID, bson.D, error) {

	i, bm, err := parse(b, "u")

	id, err := primitive.ObjectIDFromHex(i.(string))
	if err != nil {
		return nil, nil, err
	}

	return &id, bm, err

}

//Update receives the json encoded value of the operation parameters containing the document id
//and fields to update.
//Update returns an error if there was an issue parsing the byte slice.
//Update returns an error if the update operation fails.
//Update returns ErrInvalidHex if the hex string is invalid, returns ErrNotModified if the operation completes without modifying any document and returns
//ErrNoMatchedDocument if no document matched the filter
func (ls *LiveSession) Update(b []byte) error {

	//Parse JSON encoded values and return values for Id and params
	id, params, err := parseUpdate(b)
	if err != nil {
		return err
	}

	//build filter
	filter := bson.D{
		{Key: "_id", Value: bson.D{
			{Key: "$eq", Value: *id},
		}},
	}

	//build update
	update := bson.D{
		{Key: "$set", Value: params},
	}

	//Run updateOne operation and pass necessary parameters
	res, err := ls.Collection.UpdateOne(context.Background(), filter, update) //problem
	if err != nil {
		return err
	}

	//Check if any document matched
	if res.MatchedCount == 0 {
		return ErrNoMatchedDocument
	}

	//Check if the document was modified
	if res.ModifiedCount == 0 {
		return ErrNotModified
	}

	return nil

}

//Update (mock) dummies the Update operation. Used for testing
func (ms MockSession) Update(b []byte) error {

	log.Println("hetr1")
	var d map[string]string
	err := json.Unmarshal(b, &d)
	if err != nil {
		return err
	}

	log.Println("hetr2")
	_, err = primitive.ObjectIDFromHex(d["id"])
	if err != nil {
		return ErrInvalidHex
	}

	log.Println("hetr3")
	if d["id"] == "4af9f070a466655aeb230cbd" {
		return ErrNoMatchedDocument
	}

	log.Println("hetr4")
	return nil

}
