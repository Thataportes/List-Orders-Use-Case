package main

import (
	"database/sql"
	"log"
	"net/http"

	"example.com/app/domain/orderapp"
	"example.com/app/sdk/mux" // Caminho ajustado para sdk/mux
	"example.com/business/domain/orderbus"

	_ "github.com/lib/pq"
)

func main() {
	// Conexão com o banco de dados PostgreSQL
	db, err := sql.Open("postgres", "postgres://user:password@db:5432/orderdb?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Configura o serviço de negócios e o app
	service := &orderbus.OrderService{DB: db}
	app := &orderapp.OrderApp{Service: service}

	// Configura e inicia o roteador com handlers registrados
	router := mux.NewRouter()
	orderapp.RegisterRoutes(router, app) // Passa o roteador e o app com os handlers

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
