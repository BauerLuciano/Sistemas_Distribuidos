FROM golang:1.26-alpine

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o servidor_app cmd/servidor/main.go

EXPOSE 8080

CMD ["./servidor_app"]