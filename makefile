run:
	go run cmd/main.go

swag:
	swag init -g api/api.go -o api/docs

migrate-up:
	migrate -path ./migration/postgres -database 'postgres://javohir:javohir1@localhost:5432/postgres_connect?sslmode=disable' up

migrate-down:
	migrate -path ./migration/postgres -database 'postgres://javohir:javohir1@localhost:5432/postgres_connect?sslmode=disable' down