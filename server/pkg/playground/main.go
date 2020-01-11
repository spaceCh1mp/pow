package main

import (
	"context"
	"log"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":9093", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	c := v1.NewTransactionsClient(cc)

	//User

	//Do what you want with the usersClient here
	// r, err := c.Create(context.Background(), &v1.User{
	// 	FirstName: "Favour",
	// 	LastName:  "George",
	// 	UserName:  "spaceCh1mp",
	// 	Email:     "Kene@gmail.com",
	// 	Password:  "KEnrujivm",
	// })

	// r, err := c.Update(context.Background(), &v1.UpdateUser{
	// 	Id:        "5e196fb9039bfda7b345b013",
	// 	FirstName: "Hadouken",
	// })

	// r, err := c.Read(context.TODO(), &v1.ID{
	// 	Id: "5e196fb9039bfda7b345b013",
	// })

	//Group
	// r, err := c.Create(context.TODO(), &v1.NewGroup{
	// 	Name: "Spc",
	// 	Rate: 32.09,
	// })

	// r, err := c.Read(context.TODO(), &v1.ID{Id: "5e19710b039bfda7b345b014"})

	//Transaction
	// r, err := c.Create(context.TODO(), &v1.NewTransaction{
	// 	Amount: 500,
	// 	Userid: "5e196fb9039bfda7b345b013",
	// 	Date:   time.Now().String(),
	// })

	r, err := c.ReadTransaction(context.TODO(), &v1.ID{
		Id: "5e1971ff039bfda7b345b015",
	})

	if err != nil {
		log.Println(err)
	}

	log.Println(r)
}
