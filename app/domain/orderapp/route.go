package orderapp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes registra as rotas REST para o app OrderApp.
func RegisterRoutes(router *mux.Router, app *OrderApp) {
	router.HandleFunc("/order", app.ListOrdersHandler).Methods("GET")

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
