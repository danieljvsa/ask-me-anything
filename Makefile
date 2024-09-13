postgres: 
	docker run --name postgres -p 5432:5432 -e - POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:latest

createdb: 
	docker exec -it postgres createdb --username=root --owner=root ama

dropdb: 
	docker exec -it postgres dropdb --username=root ama

migrateup: 
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/ama?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/ama?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

.PHONY:
	postgres createdb dropdb migrateup migratedown sqlc