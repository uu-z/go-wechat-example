
FROM golang:alpine AS builder

WORKDIR /go/src/builder

ADD . /go/src/builder

RUN  go build -o main

FROM alpine

WORKDIR /app

COPY --from=builder /go/src/builder /app

EXPOSE 8080

ENTRYPOINT [ "./main" ]