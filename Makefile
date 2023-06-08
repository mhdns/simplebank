postgres:
	docker run --name postgres14 -e POSTGRES_PASSWORD=12345678 -p 5432:5432 -v ${HOME}/db/postgres/data:/var/lib/postgresql/data -d postgres

stoppg:
	docker stop postgres14

startpg:
	docker start postgres14

createdb:
	docker exec -it postgres14 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres14 dropdb -U postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:12345678@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:12345678@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres stoppg startpg createdb dropdb migrateup migratedown sqlc
