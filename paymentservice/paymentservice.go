package paymentservice

import (
	"context"

	ecommerce "ecommerce/generated"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type PaymentServiceServer struct {
	ecommerce.UnimplementedPaymentServiceServer
}

func (s *PaymentServiceServer) GenerateBill(ctx context.Context, req *ecommerce.GenerateBillRequest) (*ecommerce.GenerateBillResponse, error) {
	var totalAmount int32 = 0
	for _, item := range req.Items {
		totalAmount += item.Amount * 50
	}

	return &ecommerce.GenerateBillResponse{
		BillId:      uuid.NewString(),
		TotalAmount: totalAmount,
	}, nil
}

func RegisterPaymentService(paymentServer *grpc.Server) {
	ecommerce.RegisterPaymentServiceServer(paymentServer, &PaymentServiceServer{})
}
