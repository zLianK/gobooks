FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add build-base
RUN GOOS=linux CGO_ENABLED=1 go build -ldflags="-w -s" -o server cmd/gobooks/main.go

FROM alpine
COPY --from=builder /app/server .
CMD ["/server"]