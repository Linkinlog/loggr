FROM golang:1.22 AS builder

LABEL authors="log"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/loggr/main.go

FROM alpine:latest

ARG IMG_BB_KEY
ENV IMG_BB_KEY=$IMG_BB_KEY

ARG SQLITE_ADDR
ENV SQLITE_ADDR=$SQLITE_ADDR

ARG SENDGRID_API_KEY
ENV SENDGRID_API_KEY=$SENDGRID_API_KEY

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
