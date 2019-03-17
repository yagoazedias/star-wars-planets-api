setup:
	(docker-compose up -d; sleep 5s)

run:
	(make setup; go run main.go)

test:
	(make setup; go test ./...)

build:
	go build main.go

install:
	go get ./...
