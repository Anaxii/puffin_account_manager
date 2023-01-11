FROM golang:1.19
#FROM ethereum/client-go:v1.10.1

WORKDIR /go/src/app

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD go run ./cmd/main