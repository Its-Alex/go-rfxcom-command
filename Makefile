deps:
	go mod download

build:
	go build -v .

build-nas:
	GOOS=linux GOARCH=arm64 go build -v .

watch:
	reflex -r "\.go$$" -s -- go run .
