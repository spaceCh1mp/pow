package group

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/spaceCh1mp/pow/server/utils"

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

/*
	TODO
		implement graceful shutdown for this service
*/

//Config initialises the Group service
func Config() error {

	s, pool := grpc.NewServer(), &db.LiveSession{}

	var group groupServer
	v1.RegisterGroupsServer(s, group)

	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("Couldn't listen on port (group) \n %v", err)
	}

	//set collection to query
	defer pool.Disconnect()
	if err := pool.SetCollection("group"); err != nil {
		return utils.ErrFC
	}

	log.Printf("Group: %v \n", s.Serve(l))

	return nil
}

//Create adds a new group to the database
func (g groupServer) Create(c context.Context, ng *v1.NewGroup) (*v1.ID, error) {

	return &v1.ID{}, nil
}

//Read gets the group information
func (g groupServer) Read(c context.Context, id *v1.ID) (*v1.Group, error) {

	resp, err := pool.Read([]byte(id.GetId()))
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

//ReadGroupOrganiser fetches the head of the group
func (g groupServer) ReadGroupOrganiser(c context.Context, id *v1.ID) (*v1.ReadUser, error) {
	groupData, err := g.Read(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	tempGrpc, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		return nil, utils.ErrInternalServerError
	}

	us := v1.NewUsersClient(tempGrpc)

	return us.Read(context.TODO(), &v1.ID{Id: groupData.GetOrgID()})
}

//Update.
func (g groupServer) Update(c context.Context, ug *v1.UpdateGroup) (*v1.Result, error) {

	res := &v1.Result{Id: ug.GetId(), Status: false}

	jsonByte, err := json.Marshal(ug)
	if err != nil {
		return res, err
	}

	err = pool.Update(jsonByte)
	if err != nil {
		return res, err
	}

	res.Status = true
	return res, nil
}

//UpdateMember makes no sense now that i think about it
func (g groupServer) UpdateMember(c context.Context, gm *v1.GroupMember) (*v1.Result, error) {

	return member("$Push")

}

//Delete gets rid of the group, I should start with this
func (g groupServer) Delete(c context.Context, id *v1.ID) (*v1.Result, error) {

	err := pool.Delete(id.GetId())
	if err != nil {
		return nil, err
	}

	return &v1.Result{Status: true}, nil
}

//DeleteMember gets rid of a user in the group
func (g groupServer) DeleteMember(c context.Context, mg *v1.GroupMember) (*v1.Result, error) {

	return member("$Pull")

}

func member(op string) (*v1.Result, error) {

	res := &v1.Result{Id: ug.GetId(), Status: false}

	jsonByte, err := json.Marshal(ug)
	if err != nil {
		return res, err
	}

	err = pool.UpdateArray(jsonByte, op)
	if err != nil {
		return res, err
	}

	res.Status = true
	return res, nil

}
