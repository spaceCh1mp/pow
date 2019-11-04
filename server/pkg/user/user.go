package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/bson"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	db "github.com/spaceCh1mp/pow/server/db"
	grpc "google.golang.org/grpc"
)

type usersServer struct {
}

var (
	errMsg = fmt.Errorf("Not Implemented this method yet")
	pool   db.MongoDB
)

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

func collection(str string) error {
	err := pool.SetCollection(str)
	if err != nil {
		//try to re-establish a connection
		//if it fails return the error
		err = pool.Connect()
		if err != nil {
			return errFC
		}
	}
	return nil
}

//Config initialises the service
func Config() {
	s := grpc.NewServer()
	var user usersServer

	v1.RegisterUsersServer(s, user)
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatalf("Couldn't listen on port. %v", err)
	}

	pool = &db.LiveSession{}
	//set collection to query
	if err = collection("user"); err != nil {
		return
	}
	defer pool.Disconnect()

	log.Printf("User: %v \n", s.Serve(l))
}

//Implement gRPC crud methods.
func (v usersServer) Create(c context.Context, u *v1.User) (*v1.ID, error) {
	//validate input data
	err := func() error {
		if (*u == v1.User{} || u == nil) {
			return errED
		}
		if isEmpty(u.GetFirstName()) {
			return errMF
		}
		if isEmpty(u.GetLastName()) {
			return errML
		}
		if isEmpty(u.GetUserName()) {
			return errUN
		}
		if isEmpty(u.GetEmail()) {
			return errME
		}
		if isEmpty(u.GetPassword()) {
			return errMP
		}
		return nil
	}()
	if err != nil {
		return nil, err
	}

	//make query
	resp, err := pool.Insert(u)
	if err != nil {
		return nil, err
	}

	return &v1.ID{
		Id: resp,
	}, nil
}

func (v usersServer) Read(c context.Context, id *v1.ID) (*v1.ReadUser, error) {

	b, err := json.Marshal(id)
	if err != nil {
		return nil, err
	}

	r, err := pool.Read(b)
	if err != nil {
		//if there's no response
		//if there's an err from the operation
		return nil, err
	}

	b, err = bson.Marshal(r)
	if err != nil {
		return nil, err
	}

	var z v1.User
	err = bson.Unmarshal(b, &z)
	if err != nil {
		return nil, err
	}

	return &v1.ReadUser{
		FirstName: z.GetFirstName(),
		Email:     z.GetEmail(),
		UserName:  z.GetUserName(),
	}, nil
}

func (v usersServer) Update(c context.Context, u *v1.UpdateUser) (*v1.Result, error) {

	b, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	err = pool.Update(b)
	if err != nil {
		return nil, err
	}

	return &v1.Result{
		Status: true,
	}, nil

}

func (v usersServer) Delete(c context.Context, id *v1.ID) (*v1.Result, error) {

	err := pool.Delete(id.GetId())
	if err != nil {
		return &v1.Result{Status: false}, nil
	}

	return &v1.Result{Status: true}, nil

}
