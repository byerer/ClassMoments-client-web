FROM golang1.21 as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go run