FROM golang:1.22.6-alpine3.19

RUN apk update && apk add --no-cache git

WORKDIR /var/www

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

CMD ["air"]
