# Terceiro desafio Pós GO Expert (Clean Architecture)
---
Criar um UseCase de listagem de Pedidos.

Esta listagem precisa ser feita com:

* Endpoint REST (GET /order)
* Service ListOrders com GRPC
* Query ListOrders GraphQL

## Requiuisitos
  * Golang versão 1.23 ou superior
  * Docker e docker-compose

## Como executar
Clone o repositório com o comando:
```bash
git clone https://github.com/robinsonrenato/robinsonrenato/goexpert-clean-architecture.git
```

Atualizar todas as dependências do projeto:
```bash
go mod tidy
```

Inicie o banco de dados e o rabbitmq:
```bash
docker-compose up -d
```

Inicie o servidor:
```bash
go run cmd/main.go cmd/wire_gen.go
```

Interaja com  a API que está na pasta /api:
```bash
create_order.http
list_orders.http
```
