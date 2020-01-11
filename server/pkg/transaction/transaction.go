package main

import (
	"context"
	"fmt"
	"log"
	"net"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	grpc "google.golang.org/grpc"
)

type transactionsServer struct{}

func main() {
	log.Fatalln(config())
}

//Config initialises the Transactions service
func config() error {
	s := grpc.NewServer()
	var transaction transactionsServer
	v1.RegisterTransactionsServer(s, transaction)

	l, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("Couldn't listen on port \n %v", err)
	}

	return fmt.Errorf("%v", s.Serve(l))
}

func (t transactionsServer) Create(c context.Context, nt *v1.NewTransaction) (*v1.Result, error) {
	return &v1.Result{}, nil
}
func (t transactionsServer) ReadTransaction(c context.Context, td *v1.ID) (*v1.Transaction, error) {
	return &v1.Transaction{}, nil
}
