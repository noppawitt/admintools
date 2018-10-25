dev:
	fresh
prod:
	go run main.go -env=production
build:
	go build -o bin/MACOSX/admintools main.go
	cp config.dev.json bin/MACOSX/config.dev.json
	cp config.prod.json bin/MACOSX/config.prod.json
clean:
	rm -rf bin
windows:
	GOOS=windows GOARCH=amd64 go build -o bin/WINDOWS/adminntools.exe main.go
	cp config.dev.json bin/WINDOWS/config.dev.json
	cp config.prod.json bin/WINDOWS/config.prod.json
