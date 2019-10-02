package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//MockSession mocks the instance of the mongo server
// and dummies a response similar to a live one
type MockSession struct {
	C bool
}

//InsertUser ...
func (ms MockSession) InsertUser(document interface{}) (string, error) {
	return primitive.NewObjectID().Hex(), nil
}

//ReadUser ...
func (ms MockSession) ReadUser(filter bson.M) (Users, error) {
	m := filter["_id"].(string)

	//dummy reply for when the Find operation fails
	//possible cause would be a wrong ObjectId
	_, err := primitive.ObjectIDFromHex(m)
	if err != nil {
		r := mongo.SingleResult{}
		return Users{}, r.Err()
	}

	//dummy reply for when no document is returned
	if m == "4af9f070a466655aeb230cbd" {
		return Users{}, mongo.ErrNoDocuments
	}

	//If everything goes well
	return Users{
		FamilyID:  "",
		FirstName: "kene",
		LastName:  "agugua",
		UserName:  "spaceCh1mp",
		Email:     "Kenechukwuagugua@gmail.com",
	}, nil
}

//Connect Does nothing
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
