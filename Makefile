postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root bank_simple

dropdb:
	docker exec -it postgres15 dropdb bank_simple

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_simple?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/bank_simple?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate	

test:
	go test -v ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test