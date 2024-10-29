package orderapp

import (
	"example.com/business/domain/orderbus"
)

type OrderApp struct {
	Service *orderbus.OrderService
}
