package user

import (
	"fmt"
	"log"
	"testing"

	context "golang.org/x/net/context"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
)

//global interface for usersServer
var usersServerTest usersServer

func TestCreate(t *testing.T) {
	// test table containing test cases for the create user method
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
			fmt.Errorf("Missing Field: FirstName")},
		{"No LastName", v1.NewUser{
			FirstName: "Kenechukwu",
			Email:     "Kenechukwuagugua@gmail.com",
			Password:  "123455788",
		},
			fmt.Errorf("Missing Field: LastName")},
		{"No Email", v1.NewUser{
			FirstName: "Kenechukwu",
			LastName:  "Agugua",
			Password:  "123455788",
		},
			fmt.Errorf("Missing Field: Email")},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			_, err := usersServerTest.Create(context.Background(), &v.tc)
			if err != v.err {
				t.Fatalf("Expected: %v \n Got: %v", v.err, err)
			}
		})
	}
}

func TestGenRandString(t *testing.T) {
	//declare string array for a multiple of five instances
	var str [5]string
	//assign values to each instance
	//(i.e) index of the array from the function to test
	for i := range str {
		str[i] = genRandString()
	}

	//iterates through instances
	for i, v := range str {
		for _, k := range str[i+1:] {
			if v == k {
				t.Fail()
				log.Println("strings should not match")
			}
		}
	}
}

// func TestCow(t *testing.T) {
// 	got := Cow()
// 	expected := "Co"
// 	if got != expected {
// 		t.Fatalf("%v | %v", got, expected)
// 	}
// }
