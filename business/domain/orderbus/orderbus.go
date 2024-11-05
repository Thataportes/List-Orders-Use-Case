package orderbus

import (
	"database/sql"
	"fmt"
)

// OrderService gerencia as operações de listagem de ordens.
type OrderService struct {
	DB *sql.DB
}

// TODO: usar context nas queries.
// ListOrders retorna uma lista de todas as orders do banco de dados.
func (s *OrderService) ListOrders() ([]Order, error) {
	rows, err := s.DB.Query("SELECT id, item, quantity, price FROM orders")
	if err != nil {
		return nil, fmt.Errorf("error querying orders: %w", err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.Item, &o.Quantity, &o.Price); err != nil {
			return nil, fmt.Errorf("error scanning order row: %w", err)
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (s *OrderService) CreateOrder(order *Order) error {
	query := "INSERT INTO orders (item, quantity, price) VALUES ($1, $2, $3)"
	_, err := s.DB.Exec(query, order.Item, order.Quantity, order.Price)
	return err
}
