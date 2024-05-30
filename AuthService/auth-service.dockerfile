#base go image
FROM golang:1.22-alpine AS builder

RUN mkdir /app

COPY . /app


WORKDIR /app

RUN CGP_ENABLED=0 go build -o auth-service ./cmd/api

RUN go mod download

RUN chmod +x /app/auth-service


# build final image

FROM alpine

RUN mkdir /app

COPY --from=builder /app/auth-service /app


CMD ["/app/auth-service"]