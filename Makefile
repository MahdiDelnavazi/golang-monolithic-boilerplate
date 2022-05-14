createdb:
	docker exec -it golang_monolithic_boilerplate_image createdb --username=root --owner=root golang_monolithic_boilerplate

dropdb:
	docker exec -it golang_monolithic_boilerplate_image dropdb golang_monolithic_boilerplate

redis:
	docker run --name golang_monolithic_bilerplate_redis -p 6379:6379 -d redis

test:
	go test -v -cover ./...

server:
	go run main.go

swagger:
	docker run --rm -it --env GOPATH=/go -v GOPATH/go/src -w /go/src quay.io/goswagger/swagger

.PHONY: createdb dropdb test server redis swagger