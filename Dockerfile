FROM golang:1.18

WORKDIR  /go/src/github.com/SANTHOSH17-DOT/YOU-ARE-BOMBED

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

CMD bash docker-entry.sh