package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"ecommerce/etrace"
	ecommerce "ecommerce/generated"
)

func randomItem() *ecommerce.Item {
	items := []string{"Item A", "Item B", "Item C", "Item D", "Item E"}

	return &ecommerce.Item{
		Id:          uuid.NewString(),
		Amount:      int32(rand.Intn(100) + 1),
		Description: items[rand.Intn(len(items))],
	}
}

func main() {
	etrace.OtelSetup("client", "http://localhost:14268/api/traces")

	conn, err := grpc.Dial("localhost:50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		log.Fatalf("failed to connect to order service: %v", err)
	}
	defer conn.Close()

	orderClient := ecommerce.NewOrderServiceClient(conn)

	for i := 1; i <= 10; i++ {
		orderReq := &ecommerce.CreateOrderRequest{
			Items: []*ecommerce.Item{
				randomItem(),
				randomItem(),
				randomItem(),
			},
		}

		orderResp, err := orderClient.CreateOrder(context.Background(), orderReq)
		if err != nil {
			log.Printf("Error creating order: %v\n", err)
			continue
		}

		log.Printf("Order %d created with ID: %s\n", i, orderResp.OrderId)
		time.Sleep(1 * time.Second)
	}
}
