package mux

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter configura e retorna uma nova instância do roteador mux.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	return router
}

// RegisterRoutes adiciona as rotas ao roteador principal.
func RegisterRoutes(router *mux.Router, routeConfig func(r *mux.Router)) {
	routeConfig(router)
}

// loggingMiddleware é um middleware opcional para logar as requisições.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request - Method: %s, URI: %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
