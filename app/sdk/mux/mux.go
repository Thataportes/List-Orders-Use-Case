package mux

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter configura e retorna uma nova instância do roteador mux.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware) // Middleware opcional para log de requisições
	return router
}

// RegisterRoutes adiciona as rotas ao roteador principal.
func RegisterRoutes(router *mux.Router, routeConfig func(r *mux.Router)) {
	routeConfig(router)
}

// loggingMiddleware é um middleware opcional para logar as requisições.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log de exemplo; pode ser personalizado ou removido
		fmt.Printf("Request - Method: %s, URI: %s\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
