FROM golang:1.23-alpine AS build
WORKDIR /src
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /server ./cmd/server

FROM alpine:3.21
RUN apk add --no-cache wget && addgroup -S app && adduser -S -G app app
COPY --from=build /server /app/server
USER app
EXPOSE 8080
ENTRYPOINT ["/app/server"]
