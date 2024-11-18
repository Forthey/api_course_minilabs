FROM golang:alpine3.20
LABEL authors="Forthey"

WORKDIR /

COPY server server
COPY .env .

WORKDIR /server

RUN go mod tidy

RUN go build

ENTRYPOINT ["go", "run", "."]

EXPOSE 8080