FROM golang:1.21 AS builder

WORKDIR /app

COPY app/go.mod app/go.sum ./
RUN go mod tidy

COPY app/ ./

RUN go build -o /my_app ./main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /my_app ./

COPY app/templates ./templates
COPY app/static ./static

CMD ["/app/my_app"]

