FROM golang:1.20.4-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod .
COPY go.sum .

RUN go mod download


# Create the /app/tmp directory
RUN mkdir /app/tmp

# Set permissions for the /app/tmp directory
RUN chmod 777 /app/tmp

CMD ["air", "-c", ".air.toml"]
