run:
	go run main.go -dotenv
docker:
	docker build -t admintools .
test:
	go test ./... -cover
build:
	go build -o bin/admintools main.go
	cp .env bin
