
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./japanese-practice-api ./main.go
 
 
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/japanese-practice-api .
ENV DB_HOST=ep-super-sound-a4roud6p.us-east-1.pg.koyeb.app
ENV DB_USER=root
ENV DB_PASSWORD=P1xV9lytwFam
ENV DB_NAME=koyebdb
ENV DB_PORT=5432
EXPOSE 8080
ENTRYPOINT ["./japanese-practice-api"]

