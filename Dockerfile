FROM golang:1.21
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/philtree
CMD ["./main"]