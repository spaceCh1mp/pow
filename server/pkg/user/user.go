package user

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/bson/primitive"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	db "github.com/spaceCh1mp/pow/server/db"
	grpc "google.golang.org/grpc"
)

type usersServer struct {
}

func msg() error {
	return fmt.Errorf("Not Implemented this method yet")
}

func isEmpty(v interface{}) bool {
	switch i := v.(type) {
	case string:
		if i == "" {
			return true
		}
		break
	default:
		if i == nil {
			return true
		}
		break
	}
	return false
}

//Implement gRPC crud methods.
func (v usersServer) Create(c context.Context, newUser *v1.NewUser) (*v1.ID, error) {
	//validate input data
	err := func() error {
		if (newUser == &v1.NewUser{}) {
			return fmt.Errorf("Empty Dataset: NewUser")
		}
		if isEmpty(newUser.FirstName) {
			return fmt.Errorf("Missing Field: FirstName")
		}
		if isEmpty(newUser.LastName) {
			return fmt.Errorf("Missing Field: LastName")
		}
		if isEmpty(newUser.Email) {
			return fmt.Errorf("Missing Field: Email")
		}
		if isEmpty(newUser.Password) {
			return fmt.Errorf("Missing Field: Password")
		}
		return nil
	}()
	if err != nil {
		return nil, err
	}

	client, err := db.Connect()
	if err != nil {
		//handle error properly
	}

	collection := client.Database(db.Name).Collection("user")

	res, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return nil, fmt.Errorf("Insert error")
	}

	b := res.InsertedID.(primitive.ObjectID)
	return &v1.ID{
		Id: b.Hex(),
	}, nil
}

func (v usersServer) Read(c context.Context, id *v1.ID) (*v1.ReadUser, error) {
	return &v1.ReadUser{}, msg()
}

func (v usersServer) ReadLog(c context.Context, id *v1.ID) (*v1.ReadUserLog, error) {
	return &v1.ReadUserLog{}, msg()
}

func (v usersServer) Update(c context.Context, u *v1.UpdateUser) (*v1.Result, error) {
	return &v1.Result{}, msg()
}

func (v usersServer) Delete(c context.Context, id *v1.ID) (*v1.Result, error) {
	return &v1.Result{}, msg()
}

//Config initialises the service
func Config() error {
	s := grpc.NewServer()
	var user usersServer
	v1.RegisterUsersServer(s, user)
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Couldn't listen on port. %v", err)
	}

	log.Printf("started listening on %v", l.Addr())
	return fmt.Errorf("%v", s.Serve(l))
}
