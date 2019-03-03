# Using golang image to build

FROM golang:latest as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/loogo/wms/consignment-service

COPY . .

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# Run-time image

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/loogo/wms/consignment-service/consignment-service .

CMD ["./consignment-service"]