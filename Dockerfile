# syntax=docker/dockerfile:1

# Node Build Stage
FROM node:19-alpine3.16 AS NodeBuildStage
WORKDIR /build
COPY ["react-app/package.json", "react-app/package-lock.json*", "./"]
RUN npm install --production
COPY react-app .
RUN npm run build

# Go Build Stage
FROM golang:1.19.5-alpine3.17 AS GoBuildStage
WORKDIR /build
RUN apk add build-base
COPY --from=NodeBuildStage /build/build react-app/build
COPY ["go.mod", "go.sum*", "./"]
RUN go mod download
COPY . .
RUN go build .

# Deploy Stage
FROM alpine:latest
WORKDIR /app
COPY --from=GoBuildStage /build/barber-shop .
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
RUN mkdir /data && chown nonroot:nonroot /data
USER nonroot
EXPOSE 8080
ENTRYPOINT ["/app/barber-shop"]
CMD [ "-db_filename=/data/barber.db" ]
