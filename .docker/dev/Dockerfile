FROM golang:1.22.6-alpine3.19
WORKDIR /var/www

# Set time zone for São Paulo
RUN apk add --no-cache tzdata &&  \
    cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime &&  \
    echo "America/Sao_Paulo" > /etc/timezone

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest
CMD ["air"]
