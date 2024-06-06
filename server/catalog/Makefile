PROJECT_NAME := catalog

build:
	go build -o $(PROJECT_NAME) main.go

run:
	go run ./cmd/api/

clean:
	rm -f $(PROJECT_NAME)

migrateup:
	migrate -path migrations -database "postgres://postgres:Zbekxzz3@localhost:5432/catalog-service?sslmode=disable" up

migratedown:
	migrate -path migrations -database "postgres://postgres:Zbekxzz3@localhost:5432/catalog-service?sslmode=disable" down