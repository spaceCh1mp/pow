package user

import (
	"errors"
	"testing"

	context "golang.org/x/net/context"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
)

//Don't run any tests till the mongo server has been mocked
//global interface for usersServer
var usersServerTest usersServer

func TestCreate(t *testing.T) {
	// test table containing test cases for the create user method
	er := errors.New("")
	tt := []struct {
		name string
		tc   v1.NewUser
		err  error
	}{
		{"Ok", v1.NewUser{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			Email:     "Kenechukwuagugua@gmail.com",
			Password:  "123455788",
		},
			nil},
		{"No FirstName", v1.NewUser{
			LastName: "Agugua",
			Email:    "Kenechukwuagugua@gmail.com",
			Password: "123455788",
		},
			er},
		{"No LastName", v1.NewUser{
			FirstName: "Kenechukwu",
			Email:     "Kenechukwuagugua@gmail.com",
			Password:  "123455788",
		},
			er},
		{"No Email", v1.NewUser{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			Password:  "123455788",
		},
			er},
		{"No value", v1.NewUser{}, er},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			_, e := usersServerTest.Create(context.Background(), &v.tc)
			if v.err == nil && e != v.err {
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
