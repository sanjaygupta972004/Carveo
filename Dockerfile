FROM golang:1.23.3-alpine AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download   

COPY /env /app/env

COPY . .

RUN go clean -cache -modcache -i

RUN CGO_ENABLED=0 GOOS=linux go build -o carveoApp ./main.go   

RUN chmod +x /app/carveoApp

# This is the final image
FROM alpine:3.10

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/carveoApp /app/carveoApp
COPY --from=builder /app/env /app/env

EXPOSE 8080

CMD ["/app/carveoApp"]

