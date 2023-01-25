# syntax=docker/dockerfile:1

# Build Stage
FROM golang:1.19.5-alpine3.17 AS BuildStage
WORKDIR /build
COPY . .
RUN apk add build-base
RUN go mod download
RUN go build .

# Deploy Stage
FROM alpine:latest
WORKDIR /app
COPY --from=BuildStage /build/barber-shop .
EXPOSE 8080
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
RUN mkdir /data
RUN chown nonroot:nonroot /data
USER nonroot
ENTRYPOINT ["/app/barber-shop"]
