FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o golang-products

EXPOSE 9090

CMD ./golang-products