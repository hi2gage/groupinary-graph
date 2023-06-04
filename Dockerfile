FROM golang:1.20.4-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["go", "run", "server.go"]
