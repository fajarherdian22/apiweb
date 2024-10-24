mysql:
	docker run --name mysql -e MYSQL_ROOT_PASSWORD=admin1234 -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin1234 -p 3306:3306 -d mysql:latest
start_mysql:
	docker start mysql
stop_mysql:
	docker stop mysql
createdb:
	docker exec -it mysql createdb --username=admin --owner=admin simple_bank
dropdb:
	docker exec -it mysql dropdb --username=admin
migrateup:
	migrate -path db/migration -database "mysql://admin:admin1234@localhost:3306/data?sslmode=disable" --verbose up
migratedown:
	migrate -path db/migration -database "mysql://admin:admin1234@localhost:3306/data?sslmode=disable" --verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
clearredis:
	docker exec -it redis redis-cli flushall
server:
	go run main.go

.PHONY: mysql start_mysql stop_mysql createdb dropdb migrateup migratedown sqlc test clearredis server