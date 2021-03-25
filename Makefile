postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13.2-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres13 dropdb bank

migrateup:
	migrate -path C:/Users/joaog/go/projeto/api/db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path C:/Users/joaog/go/projeto/api/db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown