FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin . 

ENV APP_ENV=prod

ENTRYPOINT ["/app/bin"]