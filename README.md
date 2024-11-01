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
   ```bash
   git clone https://github.com/Thataportes/List-Orders-Use-Case
   cd List-Orders-Use-Case

## Portas
- Servidor HTTP (REST): 8080
- gRPC: 50051
- GraphQL: 8080/graphql (Rota definida no handler do servidor HTTP.)