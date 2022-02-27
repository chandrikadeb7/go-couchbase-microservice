FROM golang:1.16-alpine3.13 as build-env
RUN apk add --no-cache git gcc
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build 
CMD ["./billing-cycle-ms-app"]