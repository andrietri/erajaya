FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o erajaya

EXPOSE 8081

CMD ./erajaya
