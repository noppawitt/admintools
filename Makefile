run:
	go run main.go -dotenv
dev:
	fresh
build:
	go build -o bin/MACOSX/admintools main.go
test:
	go test ./... -cover
windows:
	GOOS=windows GOARCH=amd64 go build -o bin/WINDOWS/adminntools.exe main.go
clean:
	rm -rf bin
