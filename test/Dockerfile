# This docker file is for being used with environment variables PORT and Duration (sec)
# docker run --rm -e PORT=80 -e DURATION=0.1 -p 8001:80 kunalvirwal/go_server

FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY ./test_servers.go .

RUN go mod init GoServer

RUN go mod tidy

RUN go build -o test_server .

CMD ["./test_server"]