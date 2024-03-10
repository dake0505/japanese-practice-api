FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./japanese-practice-api ./main.go
 
 
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/japanese-practice-api .
EXPOSE 8080
ENTRYPOINT ["./japanese-practice-api"]
