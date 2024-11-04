# List Orders Use Case

## Descrição

Este projeto implementa um use case para listar pedidos utilizando REST, gRPC e GraphQL.

## Estrutura

- `api/services/order/`: Ponto de entrada do servidor HTTP.
- `app/domain/orderapp/`: Lógica de manipulação de orders.
- `business/domain/orderbus/`: Lógica de negócio para consultar e registrar orders.
- `proto/`: Definições gRPC.
- `database/init.sql`: Migrações do banco de dados.

## Requisitos

- Docker
- Docker Compose

## Como Executar

1. Clone o repositório:

  - git clone https://github.com/Thataportes/List-Orders-Use-Case
  - cd List-Orders-Use-Case

2. Inicie os serviços com Docker Compose:
- make docker 
(docker-compose up --build)

## Portas
- Servidor HTTP (REST): 8080
- gRPC: 50051
- GraphQL: 8080/graphql (Rota definida no handler do servidor HTTP.)

As aplicações gRPC e GraphQL estão iniciadas nas seguintes portas:
- gRPC escuta na porta 50051
- GraphQL pode ser acessado em http://localhost:8080/graphql

## Testando as Requisições
API HTTP (api.http)

### Listar Ordens
GET http://localhost:8080/order HTTP/1.1
Accept: application/json

### Criar Ordem
POST http://localhost:8080/order HTTP/1.1
Content-Type: application/json

{
    "item": "Tablet",
    "quantity": 5,
    "price": 300.00
}
