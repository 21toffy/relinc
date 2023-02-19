postgres:
	docker run -d --network relinc-network --name relinc --ip=127.0.0.1 -p 5439:5439 -e POSTGRES_USER=tofunmi -e POSTGRES_PASSWORD=toffy123 postgres:13
startpostgres:
	docker start relinc
checkpostgres:
	docker status relinc
createdb:
	docker exec -it relinc createdb --username=tofunmi --owner=tofunmi relinc_db
dropdb:
	docker exec -it relinc dropdb --username=tofunmi relinc_db
migrateupcicd:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/relinc_db?sslmode=disable" -verbose up
migrateup:
	migrate -path db/migration -database "postgresql://tofunmi:toffy123@172.17.0.3:5432/relinc_db?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://tofunmi:toffy123@172.17.0.3:5432/relinc_db?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migrateup sqlc startpostgres test checkpostgres server



# postgres:
# 	docker run -d --name relinc --ip=127.0.0.1 -p 5439:5439 -e POSTGRES_USER=tofunmi -e POSTGRES_PASSWORD=toffy123 postgres:13
# startpostgres:
# 	docker start relinc
# checkpostgres:
# 	docker status relinc
# createdb:
# 	docker exec -it relinc1 createdb --username=tofunmi --owner=tofunmi relinc_db
# dropdb:
# 	docker exec -it relinc1 dropdb relinc_db
# migrateupcicd:
# 	migrate -path db/migration -database "postgresql://root:root@localhost:5432/relinc_db?sslmode=disable" -verbose up
# migrateup:
# 	migrate -path db/migration -database "postgresql://tofunmi:toffy123@172.17.0.3:5432/relinc_db?sslmode=disable" -verbose up
# migratedown:
# 	migrate -path db/migration -database "postgresql://tofunmi:toffy123@172.17.0.3:5432/relinc_db?sslmode=disable" -verbose down
# sqlc:
# 	sqlc generate
# server:
# 	go run main.go
# test:
# 	go test -v -cover ./...server
	
# .PHONY: postgres createdb dropdb migrateup migrateup sqlc startpostgres test checkpostgres server

