
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build -o salary_api ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/salary_api .

EXPOSE 8081

CMD ["./salary_api"]