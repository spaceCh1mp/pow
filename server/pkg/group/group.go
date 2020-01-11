package group

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/spaceCh1mp/pow/server/utils"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	db "github.com/spaceCh1mp/pow/server/db"
	grpc "google.golang.org/grpc"
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

	s := grpc.NewServer()
	pool = &db.LiveSession{}
	var group groupServer
	v1.RegisterGroupsServer(s, group)

	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("Couldn't listen on port (group) \n %v", err)
	}

	//set collection to query
	defer pool.Disconnect()
	if err := pool.SetCollection("family"); err != nil {
		return utils.ErrFC
	}

	log.Printf("Group: %v \n", s.Serve(l))

	return nil
}

//Create adds a new group to the database
func (g groupServer) Create(c context.Context, ng *v1.NewGroup) (*v1.ID, error) {

	//make query
	for {
		objectID := primitive.NewObjectIDFromTimestamp(time.Now().UTC())
		resp, err := pool.Insert(&db.Family{
			ID:          &objectID,
			Name:        ng.Name,
			Rate:        float64(ng.Rate),
			OrganiserID: "5d911f1f568f8287f5f8155e",
		})

		if err != nil {
			return nil, err
		}

		if _, err = primitive.ObjectIDFromHex(resp); err == nil {
			return &v1.ID{
				Id: resp,
			}, nil
		}
	}

}

//Read gets the group information
func (g groupServer) Read(c context.Context, id *v1.ID) (*v1.Group, error) {

	b, err := json.Marshal(id)
	if err != nil {
		return nil, err
	}

	//TODO
	// use db.Family
	resp, err := pool.Read(b, &v1.Group{})
	if err != nil {
		return nil, err
	}

	return resp.(*v1.Group), nil
}

//ReadGroupOrganiser fetches the head of the group
func (g groupServer) ReadGroupOrganiser(c context.Context, id *v1.ID) (*v1.ReadUser, error) {
	groupData, err := g.Read(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	tempGrpc, err := grpc.Dial(":9092", grpc.WithInsecure())
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

	return member(gm, "$Push")

}

//Delete gets rid of the group, I should start with this
func (g groupServer) Delete(c context.Context, id *v1.ID) (*v1.Result, error) {

	err := pool.Delete(id.GetId())
	if err != nil {
		return nil, err
	}

	return &v1.Result{Status: true}, nil
}

func member(gm *v1.GroupMember, op string) (*v1.Result, error) {

	res := &v1.Result{Id: gm.GetId(), Status: false}

	jsonByte, err := json.Marshal(gm)
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
