package orderapp

import (
	"context"
	"log"

	"example.com/business/domain/orderbus"
	"example.com/proto"
)

// OrderServiceGRPC implementa o serviço gRPC para listar orders.
type OrderServiceGRPC struct {
	proto.UnimplementedOrderServiceServer
	Service *orderbus.OrderService
}

// ListOrders implementa o método gRPC para listar orders.
func (s *OrderServiceGRPC) ListOrders(ctx context.Context, req *proto.ListOrdersRequest) (*proto.ListOrdersResponse, error) {
	orders, err := s.Service.ListOrders()
	if err != nil {
		log.Printf("Error querying orders via gRPC: %v", err)
		return nil, err
	}

	var grpcOrders []*proto.Order
	for _, order := range orders {
		grpcOrders = append(grpcOrders, &proto.Order{
			Id:       order.ID,
			Item:     order.Item,
			Quantity: int32(order.Quantity),
			Price:    order.Price,
		})
	}
	return &proto.ListOrdersResponse{Orders: grpcOrders}, nil
}
