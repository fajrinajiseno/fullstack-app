# =========================
# 1) Build stage
# =========================
FROM golang:1.24-bullseye AS builder

WORKDIR /app

ENV OPENAPIYAML_LOCATION=openapi.yaml

RUN apt-get update && apt-get install -y ca-certificates
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libsqlite3-dev \
    make \
    && rm -rf /var/lib/apt/lists/*


COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
COPY openapi.yaml ./openapi.yaml

# Install openapi tools (oapi-codegen, etc.)
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# Generate OpenAPI server/client/types
RUN oapi-codegen -generate "types,chi-server,spec" -package openapigen -o internal/openapigen/openapi.gen.go openapi.yaml

# Download / tidy deps
RUN go mod tidy
RUN go mod vendor

# Generate secrets 
# RUN set -x && make gen-secret
RUN go run ./script/gen-secret/main.go

# Build application binary
# RUN set -x && make build
RUN go build -o bin/mygolangapp main.go

# =========================
# 2) Runtime stage
# =========================
FROM debian:bullseye-slim

RUN useradd -m -u 10001 appuser

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    sqlite3 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

ENV OPENAPIYAML_LOCATION=/app/openapi.yaml

COPY --from=builder /app/bin/mygolangapp /app/bin/mygolangapp
COPY --from=builder /app/openapi.yaml /app/openapi.yaml

VOLUME ["/app/data"]

EXPOSE 8080

USER appuser

CMD ["/app/bin/mygolangapp"]
