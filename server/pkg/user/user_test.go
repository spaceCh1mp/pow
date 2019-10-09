package user

import (
	"testing"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/spaceCh1mp/pow/server/db"

	context "golang.org/x/net/context"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
)

//global interface for usersServer
var usersServerTest usersServer

func TestCreate(t *testing.T) {
	//Define mockSession
	pool = db.MockSession{C: true}

	// test table containing test cases for the create user method
	tt := []struct {
		name string
		tc   *v1.NewUser
		err  error
	}{
		{"Ok", &v1.NewUser{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			Email:     "Kenechukwuagugua@gmail.com",
			Password:  "123455788",
		},
			nil},
		{"No FirstName", &v1.NewUser{
			LastName: "Agugua",
			Email:    "Kenechukwuagugua@gmail.com",
			Password: "123455788",
		},
			errMF},
		{"No LastName", &v1.NewUser{
			FirstName: "Kenechukwu",
			Email:     "Kenechukwuagugua@gmail.com",
			Password:  "123455788",
		},
			errML},
		{"No Password", &v1.NewUser{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			Email:     "Kenechukwuagugua@gmail.com",
		},
			errMP},
		{"No Email", &v1.NewUser{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			Password:  "123455788",
		},
			errME},
		{"No value", &v1.NewUser{}, errED},
		{"Failed to Set Collection", &v1.NewUser{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			Email:     "Keneca@gmail.com",
			Password:  "123455788",
		},
			errFC,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			//set different connection data to test
			if v.name == "Failed to Set Collection" {
				pool = db.MockSession{C: false}
			}

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

	//reply
	r := mongo.SingleResult{}
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
			mongo.ErrNoDocuments,
		}, {
			"Invalid ObjectId", &v1.ID{Id: "g17t3jfnkkjrik"},
			r.Err(),
		}, {
			"Failed to Set Collection", &v1.ID{Id: primitive.NewObjectID().Hex()},
			errFC,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			//set different connection data to test
			if v.name == "Failed to Set Collection" {
				pool = db.MockSession{C: false}
			}

			_, e := usersServerTest.Read(context.Background(), v.in)
			if e != v.out {
				t.Fatalf("Expected: %v \n got: %v", v.out, e)
			}

		})
	}
}

func TestDelete(t *testing.T) {
	pool = &db.MockSession{C: true}

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
		}, {
			"Failed to Set Collection", &v1.ID{Id: "howdy"},
			nil, errFC,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {

			//set different connection data to test
			if v.name == "Failed to Set Collection" {
				pool = db.MockSession{C: false}
			}

			resp, err := usersServerTest.Delete(context.Background(), v.in)
			if err != v.out || resp.GetStatus() != v.res.GetStatus() {
				t.Fatalf("Expected{\nresult:%v\nerr:%v\n}\n Got{\nresult:%v\nerr:%v\n}",
					v.res.GetStatus(), v.out, resp.GetStatus(), err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	//Define mockSession
	pool = db.MockSession{C: true}

	tt := []struct {
		name   string
		u      *v1.UpdateUser
		status bool
		err    error
	}{
		{"Ok", &v1.UpdateUser{
			Id:        primitive.NewObjectID().Hex(),
			FirstName: "Kenechukwu",
			Email:     "Kenechukwuagugua@gmail.com",
		},
			true, nil},
		{"No Id", &v1.UpdateUser{
			FirstName: "Kenechukwu",
			Email:     "Kenechukwuagugua@gmail.com",
			Password:  "123455788",
		},
			false, errID},
		{"Invalid ID", &v1.UpdateUser{
			Id:        "ojmovm0emi3m03mic",
			FirstName: "Kenechukwu",
			Email:     "Kenechukwuagugua@gmail.com",
		},
			false, errID},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			got, err := usersServerTest.Update(context.Background(), v.u)
			if err != v.err || got.Status != v.status {
				t.Fatalf("Expected{\nresult:%v\nerr:%v\n}\n Got{\nresult:%v\nerr:%v\n}",
					v.status, v.err, got.GetStatus(), err)
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
