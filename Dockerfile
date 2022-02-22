FROM golang:1.17-alpine as builder

WORKDIR /go/src/
RUN update-ca-certificates
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM scratch
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/main /bin/server

CMD ["/bin/server"]
