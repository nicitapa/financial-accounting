POSTGRES_USER=postgres
POSTGRES_PASSWORD=1234
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DATABASE=financial_db

$env:DB_URL = "postgres://postgres:1234@localhost:5432/financial_db?sslmode=disable"

migrate-up:
	migrate -path ./migrations -database $(DB_URL) up

migrate-down:
	migrate -path ./migrations -database $(DB_URL) down

migrate-reset:
	migrate -path ./migrations -database $(DB_URL) down -all

migrate-create:
	migrate create -ext sql -dir ./migrations -seq -digits 4 $(name)
format:
	gofmt -w .
	goimports -w .

lint:
	golangci-lint run ./...