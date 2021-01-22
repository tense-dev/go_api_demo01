FROM golang:1.15.7-alpine3.13 AS builder
WORKDIR /src
RUN go get -d -v golang.org/x/net/html  
COPY ..
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/main.go

FROM alpine:3.13  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app .
CMD ["./app"]