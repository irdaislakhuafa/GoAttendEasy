FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o build/main src/cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/build /app/build
COPY --from=builder /app/etc /app/etc
EXPOSE 8080
CMD [ "/app/build/main", "-env", "prod" ]