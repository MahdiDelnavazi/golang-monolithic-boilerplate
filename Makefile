createdb:
	docker run --name mongodb -d -p 27017:27017 mongo:latest

redis:
	docker run --name golang_monolithic_bilerplate_redis -p 6379:6379 -d redis

test:
	go test -v -cover ./...

server:
	go run main.go

swagger:
	docker run --rm -it --env GOPATH=/go -v GOPATH/go/src -w /go/src quay.io/goswagger/swagger

.PHONY: createdb dropdb test server redis swagger