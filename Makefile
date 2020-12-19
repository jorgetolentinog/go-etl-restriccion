build_window:
	GOOS=windows GOARCH=amd64 go build -trimpath -o bin/app.exe

build_linux:
	GOOS=linux GOARCH=amd64 go build -trimpath -o bin/app

test:
	go test ./...
