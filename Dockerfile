FROM golang:1.23.3-alpine AS builder

RUN mkdir /app

# set up environment variables
ENV APP_ENV=dev

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download     
COPY /env /app/env
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o carveoApp ./main.go      
RUN chmod +x /app/carveoApp

EXPOSE 8090

CMD [ "/app/carveoApp" ]

# This is the final image
FROM alpine:3.10

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/carveoApp /app/carveoApp
COPY --from=builder /app/env /app/env

EXPOSE 8090

CMD ["/app/carveoApp"]

