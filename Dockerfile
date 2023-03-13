FROM golang:1.20-alpine AS build_base
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN cd cmd/gpio-to-api && \
    go build -o /build/out/my-app .

# Start fresh from a smaller image
FROM alpine:3.17.2
RUN apk add ca-certificates
COPY --from=build_base /build/out/my-app /app/gpio-to-api
WORKDIR /app
CMD ["/app/gpio-to-api"]
