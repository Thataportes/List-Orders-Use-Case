package orderapp

import (
	"context"

	"example.com/business/domain/orderbus"
)

// Resolver implementa o resolver GraphQL para listar orders.
type Resolver struct {
	Service *orderbus.OrderService
}

// ListOrders resolve a consulta GraphQL para listar orders.
func (r *Resolver) ListOrders(ctx context.Context) ([]orderbus.Order, error) {
	return r.Service.ListOrders()
}
