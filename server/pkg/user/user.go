package user

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	grpc "google.golang.org/grpc"
)

type usersServer struct {
}

func msg() error {
	return fmt.Errorf("Not Implemented this method yet")
}

func genRandString() string {
	v, pool := "", []string{"1", "2", "4", "5", "6", "8", "0", "f", "J", "o", "s", "W", "a", "7", "l", "m", "n", "b", "c", "C", "T", "r", "h"}
	for i := 0; i < 12; i++ {
		rand.Seed(time.Now().UnixNano())
		v += pool[rand.Intn(len(pool))]
	}
	return v
}

//Implement gRPC crud methods.
func (v usersServer) Create(c context.Context, newUser *v1.NewUser) (*v1.ID, error) {
	return &v1.ID{
		Id: genRandString(),
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
