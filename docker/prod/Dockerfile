FROM golang:1.22.6-alpine3.19 as builder
RUN apk update
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o strange-air ./cmd/api

FROM alpine:3.19 as server
WORKDIR /var/www
COPY --from=builder /build/strange-air /var/www
COPY --from=builder /build/.env /var/www

CMD ["./strange-air"]