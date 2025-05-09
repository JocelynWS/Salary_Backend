FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o salary-api ./cmd

# --- Runtime image ---
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

RUN useradd -m appuser

# ⚠️ Tạo đúng đường dẫn mà code yêu cầu
RUN mkdir -p /home/jocelyn/salary_api_ver1/frontend

WORKDIR /home/appuser

# Copy binary
COPY --from=builder /app/salary-api .

# ⚠️ Copy index.html vào đúng vị trí hardcoded
COPY --from=builder /app/frontend/index.html /home/jocelyn/salary_api_ver1/frontend/index.html

USER appuser

EXPOSE 8081

ENTRYPOINT ["./salary-api"]
