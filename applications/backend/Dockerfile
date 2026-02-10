FROM golang:1.25.7-alpine

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]

EXPOSE 5000

