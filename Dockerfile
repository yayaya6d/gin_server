# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /docker-gin-server

EXPOSE 8080

CMD [ "/docker-gin-server" ]
