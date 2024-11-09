FROM golang:1.22.9-bullseye as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tonx_backend main.go

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/tonx_backend .

COPY ./scripts/flight.json ./scripts/flight.json
COPY .env .env

EXPOSE 8888

CMD ["./tonx_backend"]
