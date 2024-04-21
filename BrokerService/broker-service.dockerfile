#base go image
FROM golang:1.22-alpine AS builder

RUN mkdir /app

COPY . /app


WORKDIR /app

RUN CGP_ENABLED=0 go build -o broker-service ./cmd/api

RUN go mod download

RUN chmod +x /app/broker-service


# build final image

FROM alpine

RUN mkdir /app

COPY --from=builder /app/broker-service /app


CMD ["/app/broker-service"]