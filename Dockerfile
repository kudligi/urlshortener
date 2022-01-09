
FROM golang:1.16-alpine

WORKDIR $GOPATH/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

ENV SHORT_URL_SIZE=8
ENV APP_DOMAIN="http://localhost:9090/"


COPY ./ ./

RUN go build -o /server

EXPOSE 9090

CMD ["/server"]
