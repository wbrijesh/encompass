FROM golang:1.20-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/main.go ./
RUN go build -o main

CMD ["./main"]