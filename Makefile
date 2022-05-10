postgres:
	docker run --name golang_monolithic_boilerplate_image -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it golang_monolithic_boilerplate_image createdb --username=root --owner=root golang_monolithic_boilerplate

dropdb:
	docker exec -it golang_monolithic_boilerplate_image dropdb golang_monolithic_boilerplate

migrateup:
	migrate -path database/schema -database "postgresql://root:secret@localhost:5432/golang_monolithic_boilerplate?sslmode=disable" -verbose up

migratedown:
	migrate -path database/schema -database "postgresql://root:root@localhost:5432/golang_monolithic_boilerplate?sslmode=disable" -verbose down

redis:
	docker run --name golang_monolithic_bilerplate_redis -p 6379:6379 -d redis

test:
	go test -v -cover ./...

server:
	go run main.go


swagger:
	docker run --rm -it --env GOPATH=/go -v GOPATH/go/src -w /go/src quay.io/goswagger/swagger

.PHONY: postgres createdb dropdb migrateup migratedown test server redis swagger