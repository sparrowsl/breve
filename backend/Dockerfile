FROM golang:1.22-alpine

WORKDIR /app

COPY go.* .

RUN go mod tidy

COPY . .

EXPOSE 5000

CMD [ "go", "run", "./cmd/api/" ]
