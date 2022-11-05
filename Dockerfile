FROM golang:1.18

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .