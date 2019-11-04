package main

import (
	"context"
	"log"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":9092", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	c := v1.NewUsersClient(cc)

	//Do what you want with the usersClient here
	// r, err := c.Create(context.Background(), &v1.User{
	// 	FirstName: "Favour",
	// 	LastName:  "George",
	// 	UserName:  "spaceCh1mp",
	// 	Email:     "Kene@gmail.com",
	// 	Password:  "KEnrujivm",
	// })

	r, err := c.Update(context.Background(), &v1.UpdateUser{
		Id:        "5db7437d8fc0d3cfd36b14cd",
		FirstName: "Hadouken",
	})

	// r, err := c.Read(context.TODO(), &v1.ID{Id: "5db7437d8fc0d3cfd36b14cd"})

	if err != nil {
		log.Println(err)
	}

	log.Println(r)
}
