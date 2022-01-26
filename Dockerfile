FROM golang:1.17-alpine as builder

WORKDIR /go/src/
RUN update-ca-certificates
COPY . .
RUN go build main.go

FROM golang:1.17-alpine
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/main /bin/server

CMD ["/bin/server"]
