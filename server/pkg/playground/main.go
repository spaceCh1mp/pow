package main

import (
	"log"
	"context"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":9091", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	c := v1.NewGroupsClient(cc)

	//Do what you want with the usersClient here
	// r, err := c.Create(context.Background(), &v1.User{
	// 	FirstName: "Favour",
	// 	LastName:  "George",
	// 	UserName:  "spaceCh1mp",
	// 	Email:     "Kene@gmail.com",
	// 	Password:  "KEnrujivm",
	// // })

	// r, err := c.Update(context.Background(), &v1.UpdateUser{
	// 	Id:        "5db7437d8fc0d3cfd36b14cd",
	// 	FirstName: "Hadouken",
	// })

	gr, err := c.Read(context.TODO(), &v1.ID{
		Id: "5e131acdd441d4705893fa1a",
	})

	// gr, err := c.Create(context.TODO(), &v1.NewGroup{
	// 	Name: "Spc",
	// 	Rate: 32.09,
	// })
	// r, err := c.Read(context.TODO(), &v1.ID{Id: "5db7437d8fc0d3cfd36b14cd"})

	if err != nil {
		log.Println(err)
	}

	log.Println(gr)
}
