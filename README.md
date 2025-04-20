# Gateway de Pagamento - API Gateway (Go)

Este é o microsserviço da API Gateway desenvolvido em ***Go***.

## Sobre o Projeto

O Gateway de Pagamento é um sistema distribuído composto por:
- [Frontend em Next.js](https://github.com/patrick-cuppi/Gateway-Payment-Next.js)
- API Gateway em Go (este repositório) com **Apache Kafka** para comunicação assíncrona
- [Sistema de Antifraude em Nest.js](https://github.com/patrick-cuppi/Gateway-Payment-Nest.js)

### Componentes do Sistema

- **Frontend (Next.js)**
  - Interface do usuário para gerenciamento de contas e processamento de pagamentos
  - Desenvolvido com Next.js para garantir performance e boa experiência do usuário

- **Gateway (Go)**
  - Sistema principal de processamento de pagamentos
  - Gerencia contas, transações e coordena o fluxo de pagamentos
  - Publica eventos de transação no Kafka para análise de fraude

- **Apache Kafka**
  - Responsável pela comunicação assíncrona entre API Gateway e Antifraude
  - Garante confiabilidade na troca de mensagens entre os serviços
  - Tópicos específicos para transações e resultados de análise

- **Antifraude (Nest.js)**
  - Consome eventos de transação do Kafka
  - Realiza análise em tempo real para identificar possíveis fraudes
  - Publica resultados da análise de volta no Kafka

## Fluxo de Comunicação

1. Frontend realiza requisições para a API Gateway via REST
2. Gateway processa as requisições e publica eventos de transação no Kafka
3. Serviço Antifraude consome os eventos e realiza análise em tempo real
4. Resultados das análises são publicados de volta no Kafka
5. Gateway consome os resultados e finaliza o processamento da transação

## Ordem de Execução dos Serviços

Para executar o projeto completo, os serviços devem ser iniciados na seguinte ordem:

1. **API Gateway (Go)** - Deve ser executado primeiro pois configura a rede Docker
2. **Serviço Antifraude (Nest.js)** - Depende do Kafka configurado pelo Gateway
3. **Frontend (Next.js)** - Interface de usuário que se comunica com a API Gateway

## Instruções Detalhadas

Cada componente do sistema possui instruções específicas de instalação e configuração em seus repectivos repositórios:

- **Serviço Antifraude**: Consulte o README na pasta do projeto [(clique aqui)](https://github.com/patrick-cuppi/Gateway-Payment-Nest.js).
- **Frontend**: Consulte o README na pasta do projeto [(clique aqui)](https://github.com/patrick-cuppi/Gateway-Payment-Next.js).

> **Importante**: É fundamental seguir a ordem de execução mencionada acima, pois cada serviço depende dos anteriores para funcionar corretamente.

## Arquitetura da aplicação
[Visualize a arquitetura completa aqui](https://link.excalidraw.com/readonly/Nrz6WjyTrn7IY8ZkrZHy)

## Pré-requisitos

- [Go](https://golang.org/doc/install) 1.24 ou superior
- [Docker](https://www.docker.com/get-started)
  - Para Windows: [WSL2](https://docs.docker.com/desktop/windows/wsl/) é necessário
- [golang-migrate](https://github.com/golang-migrate/migrate)
  - Instalação: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- [Extensão REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) (opcional, para testes)

## Setup do Projeto

1. Clone o repositório:
```bash
git clone https://github.com/patrick-cuppi/Gateway-Payment-Golang
```

2. Configure as variáveis de ambiente:
```bash
cp .env.example .env
```

3. Inicie o banco de dados:
```bash
docker compose up -d
```

4. Execute as migrations:
```bash
migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/gateway?sslmode=disable" up
```

5. Execute a aplicação:
```bash
go run cmd/app/main.go
```

## API Endpoints

### Criar Conta
```http
POST /accounts
Content-Type: application/json

{
    "name": "John Doe",
    "email": "john@doe.com"
}
```
Retorna os dados da conta criada, incluindo o API Key para autenticação.

### Consultar Conta
```http
GET /accounts
X-API-Key: {api_key}
```
Retorna os dados da conta associada ao API Key.

### Criar Fatura
```http
POST /invoice
Content-Type: application/json
X-API-Key: {api_key}

{
    "amount": 100.50,
    "description": "Compra de produto",
    "payment_type": "credit_card",
    "card_number": "4111111111111111",
    "cvv": "123",
    "expiry_month": 12,
    "expiry_year": 2025,
    "cardholder_name": "John Doe"
}
```
Cria uma nova fatura e processa o pagamento. Faturas acima de R$ 10.000 ficam pendentes para análise manual.

### Consultar Fatura
```http
GET /invoice/{id}
X-API-Key: {api_key}
```
Retorna os dados de uma fatura específica.

### Listar Faturas
```http
GET /invoice
X-API-Key: {api_key}
```
Lista todas as faturas da conta.

## Testando a API

O projeto inclui um arquivo `test.http` que pode ser usado com a extensão REST Client do VS Code. Este arquivo contém:
- Variáveis globais pré-configuradas
- Exemplos de todas as requisições
- Captura automática do API Key após criação da conta

Para usar:
1. Instale a extensão REST Client no VS Code
2. Abra o arquivo `test.http`
3. Clique em "Send Request" acima de cada requisição 