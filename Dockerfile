FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache bash ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/web

EXPOSE 8000

CMD ["./app"]
