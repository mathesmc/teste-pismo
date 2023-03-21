DB_URL=postgresql://root:secret@localhost:5432/payments?sslmode=disable


create_postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER="root" -e POSTGRES_PASSWORD="secret" -d postgres:15-alpine

start_postgres:
	docker start postgres

createdb:
	docker exec -it postgres createdb --username="root" --owner="root" mini_bank

dropdb:
	docker exec -it postgres dropdb mini_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq "$(NAME)"

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

run:
	go run main.go

.PHONY:  start_postgres run_postgres createdb dropdb migrateup migratedown  new_migration sqlc test

