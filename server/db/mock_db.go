package db

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	context "golang.org/x/net/context"
)

//MockSession mocks the instance of the mongo server
// and dummies a response similar to a live one
type MockSession struct {
	C bool
}

//InsertUser (mock) ...
func (ms MockSession) InsertUser(document interface{}) (string, error) {

	return primitive.NewObjectID().Hex(), nil

}

//ReadUser ...
func (ms MockSession) ReadUser(filter string) (*Users, error) {

	//dummy reply for when the Find operation fails
	//possible cause would be a wrong ObjectId
	_, err := primitive.ObjectIDFromHex(filter)
	if err != nil {
		r := mongo.SingleResult{}
		return nil, r.Err()
	}

	//dummy reply for when no document is returned
	if filter == "4af9f070a466655aeb230cbd" {
		return nil, mongo.ErrNoDocuments
	}

	//If everything goes well
	return &Users{
		FamilyID:  "",
		FirstName: "kene",
		LastName:  "agugua",
		UserName:  "spaceCh1mp",
		Email:     "Kenechukwuagugua@gmail.com",
	}, nil
}

func updateParams(b []byte) interface{} {
	var us Users
	_ = json.Unmarshal(b, &us)

	return us
}

//UpdateUser ...
func (ms MockSession) UpdateUser(id string, b []byte) error {
	//check if\
	var d interface{}
	err := bson.UnmarshalExtJSON(b, false, &d)
	if err != nil {

	}

	return nil
}

//DeleteUser ...
func (ms MockSession) DeleteUser(filter string) error {
	//dummy reply for when the Delete operation fails.
	//possible cause would be an invalid ObjectId
	err := valid(filter)
	if err != nil {
		return err
	}

	return nil
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

func valid(f string) error {
	_, err := primitive.ObjectIDFromHex(f)
	return err
}
