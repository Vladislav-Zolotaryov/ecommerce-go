package main

import (
	"log"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"ecommerce/etrace"
	"ecommerce/orderservice"
)

func main() {
	etrace.OtelSetup("order-service", "http://jaeger:14268/api/traces")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	paymentConn, err := grpc.Dial("payment_service_server:50052",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		log.Fatalf("failed to connect to payment service: %v", err)
	}
	defer paymentConn.Close()

	orderServer := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	orderservice.RegisterOrderService(orderServer, paymentConn)

	log.Println("Order Service listening on port 50051")
	if err := orderServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
