package main

import (
	"github.com/Thashmi03/netxd_customer_connectors/constants"
	"context"
	"fmt"

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
		CustomerId: 107,
		Firstname:  "Abi",
		Lastname:   "S",
		BankId:     67890,
		Balance:    100000,
		IsActive:   false,
	})
	if err != nil {
		log.Fatalf("Failed to call: %v", err)
	}

	fmt.Printf("CustomerID: %d\nCreatedTime:%v\n", response.CustomerId,response.CreatedAt)
}