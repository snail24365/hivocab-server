postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root hivocab

dropdb:
	docker exec -it postgres12 dropdb hivocab

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/hivocab?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/hivocab?sslmode=disable" -verbose down

sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/snail24365/hivocab-server/db/sqlc Store


test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown test server mock sqlc