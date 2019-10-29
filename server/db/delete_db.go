package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	context "golang.org/x/net/context"
)

//Delete takes a valid ObjectId Hex (string) and deletes the document in the collection that matches it.package db
// ErrInvalidHex is returned if the string provided is an invalid objectId Hex,
// an error is returned if the operation fails and
// ErrNoMatchedDocument is returned if no documents matched the filter.
func (ls *LiveSession) Delete(filter string) error {

	i, err := primitive.ObjectIDFromHex(filter)
	if err != nil {
		return ErrInvalidHex
	}

	resp, err := ls.Collection.DeleteOne(context.TODO(), bson.M{"_id": i})
	if err != nil {
		return err
	}

	if resp.DeletedCount < int64(1) {
		return ErrNoMatchedDocument
	}

	return nil
}

//Delete (mock) dummies the delete operation. Used for testing
func (ms MockSession) Delete(filter string) error {
	//dummy reply for when the Delete operation fails.
	//possible cause would be an invalid ObjectId
	err := func(f string) error {

		_, err := primitive.ObjectIDFromHex(f)

		return err

	}(filter)

	if err != nil {
		return err
	}

	return nil
}
