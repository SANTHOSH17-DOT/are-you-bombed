FROM golang:1.18

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go run main.go generate flat 100

RUN go run main.go generate nested 10