# =========================
# Build stage
# =========================
FROM node:24.11.0-alpine AS builder

WORKDIR /app

ENV NODE_OPTIONS="--max-old-space-size=4096"
ENV NODE_ENV=production

RUN apk add --no-cache openjdk17-jre-headless

COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN corepack enable && pnpm install --frozen-lockfile

ENV BACKEND_API_BASE=http://backend:8080

COPY frontend/ .
COPY openapi.yaml ./openapi.yaml

RUN pnpm openapigen-docker
RUN pnpm build

EXPOSE 3000
CMD ["pnpm", "preview"]

# =========================
# Runtime
# =========================
# FROM node:24.11.0-alpine

# WORKDIR /app
# ENV NODE_ENV=production

# COPY --from=builder /app/.output ./.output
# COPY --from=builder /app/package.json ./

# EXPOSE 3000
# CMD ["node", ".output/server/index.mjs"]
