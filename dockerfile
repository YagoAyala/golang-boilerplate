FROM golang:1.21-alpine3.21 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

FROM alpine:3.21

WORKDIR /root/

COPY --from=build /app/server .
EXPOSE 8080

CMD ["./server"]
