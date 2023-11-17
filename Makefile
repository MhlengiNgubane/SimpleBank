postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root simple_bank

opendb:
	docker exec -it postgres13 psql -U root -d simple_bank

dropdb:
	docker exec -it postgres13 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up 1		

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down 1	

mysql:	
	docker run --name mysql8 -p 3306:3306 -e MYSQL_DATABASE=root -e MYSQL_ROOT_PASSWORD=password -d mysql:latest

createMySQLdb:
	docker exec -it mysql8 mysql -uroot -ppassword root	

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 mysql createMySQLdb sqlc test opendb

