FROM golang:1.14.2-alpine3.11 as builder

WORKDIR /src

COPY go.mod /src/
COPY go.sum /src/
COPY cmd/ /src/cmd/
COPY internal/ /src/internal/

RUN go mod download \
    && GOOS=linux GOARCH=arm64 go build -v -o bin/rfxcom github/It-Alex/go-rfxcom-command/cmd/rfxcom

FROM alpine:3.11.5

COPY --from=builder /src/bin/rfxcom /usr/bin/go-rfxcom

ENV GO_RFXCON_PORT=1323

EXPOSE 1323

CMD ["/usr/bin/go-rfxcom"]
