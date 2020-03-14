deps:
	go mod download

build-amd64:
	mkdir -p bin/linux/amd64
	GOOS=linux GOARCH=amd64 go build -v -o bin/rfxcom github/It-Alex/go-rfxcom-command/cmd/rfxcom

build-arm64:
	mkdir -p bin/linux/arm64
	GOOS=linux GOARCH=arm64 go build -v -o bin/linux/arm64/nas-rfxcom github/It-Alex/go-rfxcom-command/cmd/rfxcom

watch:
	reflex -r "\.go$$" -s -- go run github/It-Alex/go-rfxcom-command/cmd/rfxcom
