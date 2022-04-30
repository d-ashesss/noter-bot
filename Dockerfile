FROM golang:1.16-alpine AS builder

RUN apk add -q upx

COPY . /app
WORKDIR /app/cmd/noter

RUN go build -ldflags="-s -w" -o /app/main
RUN upx /app/main


FROM alpine:3

COPY --from=builder /app/main /app

CMD ["/app"]
