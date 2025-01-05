# this is base image for build the app
FROM golang:1.23.3-alpine as builder

# Set environment variables
ENV APP_ENV=dev
ENV DB_USER=avnadmin
ENV DB_HOST=pg-1db5e002-sanjaygupta07054-136d.g.aivencloud.com
ENV DB_PORT=16033
ENV DB_PASSWORD=AVNS_XicXqMVQz6Oi8qvbEQH
ENV DB_NAME=defaultdb
ENV DB_SSLMODE=require
ENV PORT=8090

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY config/env /app/config/env
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o carveoApp ./cmd
RUN chmod +x /app/carveoApp

# This is the final image
FROM alpine:3.10

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/carveoApp /app/carveoApp

EXPOSE 8090

CMD ["/app/carveoApp"]

