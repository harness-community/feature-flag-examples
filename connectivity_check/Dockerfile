FROM golang:1.22 as base

WORKDIR /

COPY . ./

RUN go build -o app *.go

CMD ./app

FROM debian:12-slim

COPY --from=base /app ./

EXPOSE 8080

CMD ./app