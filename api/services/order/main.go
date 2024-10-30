package main

import (
	"database/sql"
	"log"
	"net/http"

	"example.com/app/domain/orderapp"
	"example.com/app/sdk/mux"
	"example.com/business/domain/orderbus"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

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
	orderapp.RegisterRoutes(router, app)

	// Define o tipo Order para GraphQL
	orderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.Int},
			"item":     &graphql.Field{Type: graphql.String},
			"quantity": &graphql.Field{Type: graphql.Int},
			"price":    &graphql.Field{Type: graphql.Float},
		},
	})

	// Define a consulta para listar ordens
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"listOrders": &graphql.Field{
				Type: graphql.NewList(orderType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					orders, err := app.Service.ListOrders()
					if err != nil {
						log.Printf("Error querying orders: %v", err)
						return nil, err
					}
					log.Printf("Retrieved %d orders", len(orders))
					return orders, nil
				},
			},
		},
	})

	// Cria o schema GraphQL
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: query})
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// Configura o handler para o endpoint GraphQL
	router.Handle("/graphql", handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}))

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
