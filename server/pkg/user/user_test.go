package user

import (
	"testing"

	"github.com/spaceCh1mp/pow/server/db"

	context "golang.org/x/net/context"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
)

//global interface for usersServer
var usersServerTest usersServer

func TestCreate(t *testing.T) {
	//Define mockSession
	pool = db.MockSession{}

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
		{"No Email", &v1.NewUser{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			Password:  "123455788",
		},
			errME},
		{"No value", &v1.NewUser{}, errED},
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

// func TestCow(t *testing.T) {
// 	got := Cow()
// 	expected := "Co"
// 	if got != expected {
// 		t.Fatalf("%v | %v", got, expected)
// 	}
// }
