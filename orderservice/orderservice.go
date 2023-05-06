package orderservice

import (
	"context"
	"fmt"

	ecommerce "ecommerce/generated"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type OrderServiceServer struct {
	ecommerce.UnimplementedOrderServiceServer
	paymentServiceClient ecommerce.PaymentServiceClient
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *ecommerce.CreateOrderRequest) (*ecommerce.CreateOrderResponse, error) {
	billReq := &ecommerce.GenerateBillRequest{
		Items: req.Items,
	}
	_, err := s.paymentServiceClient.GenerateBill(ctx, billReq)
	if err != nil {
		return nil, fmt.Errorf("failed to generate bill: %v", err)
	}

	return &ecommerce.CreateOrderResponse{
		OrderId: uuid.NewString(),
	}, nil
}

func RegisterOrderService(orderServer *grpc.Server, paymentConn *grpc.ClientConn) {
	ecommerce.RegisterOrderServiceServer(orderServer, &OrderServiceServer{
		paymentServiceClient: ecommerce.NewPaymentServiceClient(paymentConn),
	})
}
