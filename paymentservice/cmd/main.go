package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"ecommerce/paymentservice"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	paymentServer := grpc.NewServer()
	paymentservice.RegisterPaymentService(paymentServer)

	log.Println("Payment Service listening on port 50052")
	if err := paymentServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
