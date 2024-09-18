FROM golang:1.22.6-alpine3.19 as build
RUN apk update
WORKDIR /var/www
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o strange-air ./cmd/api

FROM alpine:3.19 as app
WORKDIR /app
COPY --from=build /var/www/strange-air /app
COPY --from=build /var/www/.env /app

CMD ["./strange-air"]