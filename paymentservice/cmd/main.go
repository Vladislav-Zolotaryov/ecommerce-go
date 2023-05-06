package main

import (
	"log"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"ecommerce/etrace"
	"ecommerce/paymentservice"
)

func main() {
	etrace.OtelSetup("payment-service", "http://jaeger:14268/api/traces")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	paymentServer := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	paymentservice.RegisterPaymentService(paymentServer)

	log.Println("Payment Service listening on port 50052")
	if err := paymentServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
