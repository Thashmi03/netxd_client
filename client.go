package main

import (
	"context"
	"fmt"

	"github.com/Thashmi03/netxd_customer_connectors/constants"

	"log"

	pb "github.com/Thashmi03/netxd_customer"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(constants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)


	response, err := client.CreateCustomer(context.Background(), &pb.Details{
		CustomerId: 117,
		Firstname:  "Abi",
		Lastname:   "S",
		BankId:     67890,
		Balance:    100000,
		IsActive:   false,
	})
	if err != nil {
		log.Fatalf("Failed to call: %v", err)
	}

	res, err := client.Transfer(context.Background(), &pb.Request{
			TransactionId: 00001,
			FromAccount:   317,
			ToAccount:     318,
			Amount:         100,
		
	})
	if err != nil {
		log.Fatalf("Failed to call: %v", err)
	}
	fmt.Printf("CustomerID: %d\nCreatedTime:%v\n", response.CustomerId,response.CreatedAt)
	fmt.Printf("CustomerID: %d\n", res.TransactionId)
}