# ──────────────────────────────────────────────────────────────────────────────
# Multi-stage Dockerfile for nasha
#
# Stage 1 (node-builder): builds the React frontend.
# Stage 2 (go-builder):   compiles the Go binary.
# Stage 3 (runtime):      minimal final image.
# ──────────────────────────────────────────────────────────────────────────────

# ── Stage 1: frontend build ───────────────────────────────────────────────────
FROM node:22-alpine AS node-builder

WORKDIR /app/web

COPY web/package*.json ./
RUN npm ci

COPY web/ ./
RUN npm run build

# ── Stage 2: Go build ─────────────────────────────────────────────────────────
FROM golang:1.25-alpine AS go-builder

# CGO is needed by the SQLite driver (mattn/go-sqlite3).
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Copy the pre-built frontend bundle so it can be embedded or served at runtime.
COPY --from=node-builder /app/web/dist ./web/dist

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o /nasha ./cmd/server

# ── Stage 3: runtime image ────────────────────────────────────────────────────
FROM alpine:3.21

# ca-certificates needed for HTTPS calls to S3, WebDAV, etc.
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=go-builder /nasha /usr/local/bin/nasha
COPY --from=node-builder /app/web/dist ./web/dist
COPY config.yaml ./config.yaml

EXPOSE 8080

ENV NASHA_ADDR=":8080"

ENTRYPOINT ["nasha"]
