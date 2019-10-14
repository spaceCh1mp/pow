package group

import (
	"context"
	"fmt"
	"log"
	"net"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	db "github.com/spaceCh1mp/pow/server/db"
	grpc "google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

type groupServer struct{}

var (
	errMsg = fmt.Errorf("Not Implemented this method yet")
	pool   db.MongoDB
)

//Config initialises the Transactions service
func Config() {

	s, pool := grpc.NewServer(), &db.LiveSession{}
	defer pool.Disconnect()

	var group groupServer
	v1.RegisterGroupsServer(s, group)

	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("Couldn't listen on port (group) \n %v", err)
	}

	log.Printf("Group: %v \n", s.Serve(l))
}

//Create adds a new group to the database
func (g groupServer) Create(c context.Context, ng *v1.NewGroup) (*v1.ID, error) {

	return &v1.ID{}, nil
}

//Read gets the group information
func (g groupServer) Read(c context.Context, id *v1.ID) (*v1.Group, error) {

	defer pool.Disconnect()
	if err := pool.SetCollection("group"); err != nil {
		return nil, err
	}

	resp, err := pool.Read(id.GetId())
	if err != nil {
		return nil, err
	}

	b, err := bson.Marshal(resp)
	if err != nil {
		return nil, err
	}

	var gr v1.Group
	err = bson.Unmarshal(b, &gr)
	if err != nil {
		return nil, err
	}

	return &gr, nil
}

//ReadGroupMembers should be called inside Read tbh
func (g groupServer) ReadGroupMembers(c context.Context, id *v1.ID) (*v1.Members, error) {
	return &v1.Members{}, nil
}

//ReadGroupOrganiser fetches the head of the group
func (g groupServer) ReadGroupOrganiser(c context.Context, id *v1.ID) (*v1.ReadUser, error) {
	return &v1.ReadUser{}, nil
}

//Update.
func (g groupServer) Update(c context.Context, ug *v1.UpdateGroup) (*v1.Result, error) {
	return &v1.Result{}, nil
}

//UpdateMember makes no sense now that i think about it
func (g groupServer) UpdateMember(c context.Context, gm *v1.GroupMember) (*v1.Result, error) {
	return &v1.Result{}, nil
}

//Delete gets rid of the group, I should start with this
func (g groupServer) Delete(c context.Context, id *v1.ID) (*v1.Result, error) {

	err := pool.SetCollection("group")
	if err != nil {
		//TODO properly handle error
		return nil, err
	}

	defer pool.Disconnect()
	err = pool.Delete(id.GetId())
	if err != nil {
		return nil, err
	}

	return &v1.Result{Status: true}, nil
}

//DeleteMember gets rid of a user in the group
func (g groupServer) DeleteMember(c context.Context, mg *v1.GroupMember) (*v1.Result, error) {
	return &v1.Result{}, nil
}
