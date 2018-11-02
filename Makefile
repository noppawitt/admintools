run:
	go run main.go -dotenv
docker:
	docker build -t admintools .
test:
	go test ./... -cover
