package main

import (
	"log"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	c := v1.NewUsersClient(cc)

	//Do what you want with the usersClient here
	log.Println(c)
}
