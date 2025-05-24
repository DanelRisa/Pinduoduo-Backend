FROM golang:1.23.6

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY ./db/migrations /app/migrations

RUN go build -o main ./user-service

CMD ["./main"]