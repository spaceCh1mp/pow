package main

import (
	"context"
	"fmt"
	"log"
	"net"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	grpc "google.golang.org/grpc"
)

type groupServer struct{}

func main() {
	log.Fatalln(config())
}

//Config initialises the Transactions service
func config() error {
	s := grpc.NewServer()
	var group groupServer
	v1.RegisterGroupsServer(s, group)

	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Couldn't listen on port (group) \n %v", err)
	}

	return fmt.Errorf("%v", s.Serve(l))
}

func (g groupServer) Create(c context.Context, ng *v1.NewGroup) (*v1.ID, error) {
	return &v1.ID{}, nil
}
func (g groupServer) Read(c context.Context, id *v1.ID) (*v1.Group, error) {
	return &v1.Group{}, nil
}
func (g groupServer) ReadGroupMembers(c context.Context, id *v1.ID) (*v1.Members, error) {
	return &v1.Members{}, nil
}
func (g groupServer) ReadGroupOrganiser(c context.Context, id *v1.ID) (*v1.ReadUser, error) {
	return &v1.ReadUser{}, nil
}
func (g groupServer) Update(c context.Context, ug *v1.UpdateGroup) (*v1.Result, error) {
	return &v1.Result{}, nil
}
func (g groupServer) UpdateMember(c context.Context, gm *v1.GroupMember) (*v1.Result, error) {
	return &v1.Result{}, nil
}
func (g groupServer) Delete(c context.Context, id *v1.ID) (*v1.Result, error) {
	return &v1.Result{}, nil
}
func (g groupServer) DeleteMember(c context.Context, mg *v1.GroupMember) (*v1.Result, error) {
	return &v1.Result{}, nil
}
