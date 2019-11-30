package user

import (
	"testing"

	context "golang.org/x/net/context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	"github.com/spaceCh1mp/pow/server/db"
)

//global interface for usersServer
var usersServerTest usersServer

func TestCreate(t *testing.T) {
	//Define mockSession
	pool = db.MockSession{C: true}

	// test table containing test cases for the create user method
	tt := []struct {
		name string
		tc   *v1.User
		err  error
	}{
		{"Ok", &v1.User{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			UserName:  "space",
			Email:     "Kenechukwuagugua@gmail.com",
			Password:  "123455788",
		},
			nil},
		{"No FirstName", &v1.User{
			LastName: "Agugua",
			UserName: "space",
			Email:    "Kenechukwuagugua@gmail.com",
			Password: "123455788",
		},
			errMF},
		{"No LastName", &v1.User{
			FirstName: "Kenechukwu",
			UserName:  "space",
			Email:     "Kenechukwuagugua@gmail.com",
			Password:  "123455788",
		},
			errML},
		{"No Password", &v1.User{
			FirstName: "Kenechukwu",
			UserName:  "space",
			LastName:  "Agugua",
			Email:     "Kenechukwuagugua@gmail.com",
		},
			errMP},
		{"No Email", &v1.User{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			UserName:  "space",
			Password:  "123455788",
		},
			errME},
		{"No Username", &v1.User{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			Password:  "123455788",
		},
			errUN},
		{"No value", &v1.User{}, errED},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {

			_, e := usersServerTest.Create(context.Background(), v.tc)
			if v.err != e {
				t.Fatalf("Expected: %v \n Got: %v", v.err, e)
			}

		})
	}
}

func TestRead(t *testing.T) {
	//Define mockSession
	pool = db.MockSession{C: true}

	tt := []struct {
		name string
		in   *v1.ID
		out  error
	}{
		{
			"Ok", &v1.ID{Id: primitive.NewObjectID().Hex()},
			nil,
		}, {
			"No Result", &v1.ID{Id: "4af9f070a466655aeb230cbd"},
			db.ErrNoMatchedDocument,
		}, {
			"Invalid ObjectId", &v1.ID{Id: "g17t3jfnkkjrik"},
			db.ErrInvalidHex,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {

			_, e := usersServerTest.Read(context.Background(), v.in)
			if e != v.out {
				t.Fatalf("Expected: %v \n got: %v", v.out, e)
			}

		})
	}
}

func TestDelete(t *testing.T) {
	pool = db.MockSession{C: true}

	tt := []struct {
		name string
		in   *v1.ID
		res  *v1.Result
		out  error
	}{
		{
			"Ok", &v1.ID{Id: primitive.NewObjectID().Hex()},
			&v1.Result{Status: true}, nil,
		}, {
			"Failed to delete data", &v1.ID{Id: "4af9f070a466"},
			&v1.Result{Status: false}, nil,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {

			resp, err := usersServerTest.Delete(context.Background(), v.in)
			if err != v.out || resp.GetStatus() != v.res.GetStatus() {
				t.Fatalf("Expected{\nresult:%v\nerr:%v\n}\n Got{\nresult:%v\nerr:%v\n}",
					v.res.GetStatus(), v.out, resp.GetStatus(), err)
			}

		})
	}
}

func TestIsEmpty(t *testing.T) {
	tt := []struct {
		name string
		in   interface{}
		out  bool
	}{
		{"String", "Hello World!", false},
		{"String-Empty", "", true},
		{"Int", 0, false},
		{"Int-Empty", nil, true},
	}
	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			if got := isEmpty(v.in); got != v.out {
				t.Fatalf("Expected: %v \n Got: %v", v.out, got)
			}
		})
	}
}
