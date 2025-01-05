# this is base image for build the app
FROM golang:1.23.3-alpine as builder

# Set environment variables
ENV DB_USER=${DB_USER}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}
ENV DB_SSLMODE=${DB_SSLMODE}
ENV PORT=${PORT}

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY config/env /app/config/env
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o carveoApp ./cmd
RUN chmod +x /app/carveoApp

# This is the final image\
FROM alpine:3.10

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/carveoApp /app/carveoApp

EXPOSE 8090

CMD ["/app/carveoApp"]

