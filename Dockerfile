FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o salary-api ./cmd

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

RUN useradd -m appuser


RUN mkdir -p /home/jocelyn/salary_api_ver1/static

WORKDIR /home/appuser


COPY --from=builder /app/salary-api .


COPY --from=builder /app/static/index.html /home/jocelyn/salary_api_ver1/static/index.html

USER appuser

EXPOSE 8081

ENTRYPOINT ["./salary-api"]
