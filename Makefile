postgres:
	docker run --name postgres_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=467958 -d postgres

createdb:
	docker exec -it postgres_db createdb --username=root --owner=root microapp

dropdb:
	docker exec -it postgres_db dropdb microapp

microapp:
	docker exec -it postgres_db psql -U root microapp

migrateup:
	migrate -path db/migration -database "postgresql://root:467958@localhost:5432/microapp?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:467958@localhost:5432/microapp?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb microapp migrateup migratedown sqlc