DB_URL=postgresql://root:secret@localhost:5432/mini_bank?sslmode=disable


create_postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER="root" -e POSTGRES_PASSWORD="secret" -d postgres:15-alpine

start_postgres:
	docker start postgres

createdb:
	docker exec -it postgres createdb --username="root" --owner="root" mini_bank

dropdb:
	docker exec -it postgres dropdb mini_bank

migrateup:
	migrate -path database/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path database/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir database/migration -seq "$(NAME)"
db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

.PHONY:  start_postgres run_postgres createdb dropdb migrateup migratedown  new_migration db_docs db_schema sqlc test

