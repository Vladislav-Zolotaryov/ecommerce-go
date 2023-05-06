package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"ecommerce/orderservice"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	paymentConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to payment service: %v", err)
	}
	defer paymentConn.Close()

	orderServer := grpc.NewServer()

	orderservice.RegisterOrderService(orderServer, paymentConn)

	log.Println("Order Service listening on port 50051")
	if err := orderServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
