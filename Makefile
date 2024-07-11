postgres:
	docker run --name todo-db -p 5432:5432 -e POSTGRES_USER=todo-user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=todo-db -d postgres:latest

createdb:
	docker exec -it todo-db createdb --username=todo-user --owner=todo-user todo-db

dropdb:
	docker exec -it todo-db dropdb --username=todo-user todo-db

migrateup:
	migrate -path db/migrations -database "postgresql://todo-user:password@localhost:5432/todo-db?sslmode=disable" -verbose up

# migrateup1:
# 	migrate -path db/migrations -database "postgresql://todo-user:password@localhost:5432/todo-db?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://todo-user:password@localhost:5432/todo-db?sslmode=disable" -verbose down

# migratedown1:
# 	migrate -path db/migrations -database "postgresql://todo-user:password@localhost:5432/todo-db?sslmode=disable" -verbose down 1

up:
	docker compose up -d

down:
	docker compose down
	docker rmi todo-list

restart: down up

sqlc:
	sqlc generate

test-html:
	@echo "Creation of UI for tests..."
	@cd db/sqlc && go test -coverprofile=cover.txt ./...
	@cd db/sqlc && go tool cover -html=cover.txt

server:
	go run main.go

coverfile:
	go test -coverprofile=c.out
	go tool cover -html="c.out"

swagger:
	swag init -g internal/handlers/http/task.go

.PHONY: postgres createdb dropdb migrateup migrateup migratedown migratedown1 sqlc tests server mock storetest coverfile