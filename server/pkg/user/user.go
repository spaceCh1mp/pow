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

//Config initialises the service
func Config() {
	s := grpc.NewServer()
	var user usersServer

	v1.RegisterUsersServer(s, user)
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Couldn't listen on port. %v", err)
	}

	pool = &db.LiveSession{}

	log.Printf("User: %v \n", s.Serve(l))
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

//Implement gRPC crud methods.
func (v usersServer) Create(c context.Context, newUser *v1.NewUser) (*v1.ID, error) {
	//validate input data
	err := func() error {
		if (*newUser == v1.NewUser{} || newUser == nil) {
			return errED
		}
		if isEmpty(newUser.FirstName) {
			return errMF
		}
		if isEmpty(newUser.LastName) {
			return errML
		}
		if isEmpty(newUser.Email) {
			return errME
		}
		if isEmpty(newUser.Password) {
			return errMP
		}
		return nil
	}()
	if err != nil {
		return nil, err
	}

	//set collection to query
	if err := collection("user"); err != nil {
		return nil, err
	}

	//make query
	resp, err := pool.InsertUser(newUser)
	if err != nil {
		return nil, err
	}

	return &v1.ID{
		Id: resp,
	}, nil
}

func (v usersServer) Read(c context.Context, id *v1.ID) (*v1.ReadUser, error) {

	err := collection("user")
	if err != nil {
		return nil, err
	}

	r, err := pool.Read(id.GetId())
	if err != nil {
		//if there's no response
		//if there's an err from the operation
		return nil, err
	}

	b, err := bson.Marshal(r)
	if err != nil {
		return nil, err
	}

	var z v1.NewUser
	err = bson.Unmarshal(b, &z)
	if err != nil {
		return nil, err
	}

	return &v1.ReadUser{
		Name:     z.GetFirstName(),
		Email:    z.GetEmail(),
		Username: z.GetLastName(),
	}, nil
}

func (v usersServer) ReadLog(c context.Context, id *v1.ID) (*v1.ReadUserLog, error) {
	return &v1.ReadUserLog{}, errMsg
}

func (v usersServer) Update(c context.Context, u *v1.UpdateUser) (*v1.Result, error) {
	id := u.GetId()
	u.Id = "" //empty the userId

	b, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	err = pool.UpdateUser(id, b)
	if err != nil {
		return nil, err
	}

	return &v1.Result{
		Status: true,
	}, nil

}

func (v usersServer) Delete(c context.Context, id *v1.ID) (*v1.Result, error) {

	e := collection("user")

	if e != nil {
		return nil, e
	}

	err := pool.Delete(id.GetId())
	if err != nil {
		return &v1.Result{Status: false}, nil
	}

	return &v1.Result{Status: true}, nil
}
