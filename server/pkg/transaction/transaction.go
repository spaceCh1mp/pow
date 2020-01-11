package transaction

import (
	"context"
	"encoding/json"
	"log"
	"net"

	v1 "github.com/spaceCh1mp/pow/server/api/proto/v1"
	"github.com/spaceCh1mp/pow/server/db"
	grpc "google.golang.org/grpc"
)

type transactionsServer struct{}

var pool db.MongoDB

func collection(str string) error {
	err := pool.SetCollection(str)
	if err != nil {
		//implement error on failed connection
		return err
	}

	return nil
}

//Config initialises the service
func Config() error {

	s := grpc.NewServer()
	var transactions transactionsServer

	v1.RegisterTransactionsServer(s, transactions)
	l, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Fatalf("Couldn't listen on port. %v", err)
	}

	pool = &db.LiveSession{}

	//set collection to query
	if err = collection("Transactions"); err != nil {
		return err
	}

	defer pool.Disconnect()

	log.Printf("Transactions: %v \n", s.Serve(l))

	return nil
}

func (t transactionsServer) Create(c context.Context, nt *v1.NewTransaction) (*v1.Transaction, error) {
	resp, err := pool.Insert(nt)
	if err != nil {
		return nil, err
	}

	return &v1.Transaction{
		Id:     resp,
		Amount: nt.Amount,
		Userid: nt.Userid,
		Date:   nt.Date,
	}, nil
}

func (t transactionsServer) ReadTransaction(c context.Context, td *v1.ID) (*v1.Transaction, error) {

	b, err := json.Marshal(td)
	if err != nil {
		return nil, err
	}

	resp, err := pool.Read(b, &v1.Transaction{})
	if err != nil {
		return nil, err
	}

	return resp.(*v1.Transaction), nil

}
