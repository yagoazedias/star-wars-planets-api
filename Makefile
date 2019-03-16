setup:
	docker-compose up -d

run:
	make setup && go run main.go

test:
	make setup && go test ./...

build:
	go build main.go
