
FROM golang:1.22

WORKDIR /app

COPY go.mod .

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o main server.go

EXPOSE 50051

CMD ["./main"]
