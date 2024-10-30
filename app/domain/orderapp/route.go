package orderapp

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/business/domain/orderbus"
	"github.com/gorilla/mux"
)

// RegisterRoutes registra as rotas REST para o app OrderApp.
func RegisterRoutes(router *mux.Router, app *OrderApp) {
	router.HandleFunc("/order", app.ListOrdersHandler).Methods("GET")   // Rota para listar ordens
	router.HandleFunc("/order", app.CreateOrderHandler).Methods("POST") // Rota para criar uma nova ordem
}

// ListOrdersHandler lida com a requisição REST para listar orders.
func (app *OrderApp) ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := app.Service.ListOrders()
	if err != nil {
		log.Printf("Error querying orders: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

// CreateOrderHandler lida com a requisição REST para criar um novo order.
func (app *OrderApp) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order orderbus.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		log.Printf("Error decoding order: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := app.Service.CreateOrder(&order)
	if err != nil {
		log.Printf("Error creating order: %v", err)
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
