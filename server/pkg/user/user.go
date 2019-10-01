package user

import (
	"context"
	"fmt"
	"log"
	"net"

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

	//set collection to query
	if err := pool.SetCollection("user"); err != nil {
		//handle collection error
	}

	//make query
	resp, err := pool.InsertOne(newUser)
	if err != nil {
		//handle err
	}

	return &v1.ID{
		Id: resp,
	}, nil
}

func (v usersServer) Read(c context.Context, id *v1.ID) (*v1.ReadUser, error) {
	return &v1.ReadUser{}, errMsg
}

func (v usersServer) ReadLog(c context.Context, id *v1.ID) (*v1.ReadUserLog, error) {
	return &v1.ReadUserLog{}, errMsg
}

func (v usersServer) Update(c context.Context, u *v1.UpdateUser) (*v1.Result, error) {
	return &v1.Result{}, errMsg
}

func (v usersServer) Delete(c context.Context, id *v1.ID) (*v1.Result, error) {
	return &v1.Result{}, errMsg
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

	pool = &db.LiveSession{}

	if err := pool.Connect(); err != nil {
		//handle connection error
	}

	log.Printf("started listening on %v", l.Addr())
	return fmt.Errorf("%v", s.Serve(l))
}
