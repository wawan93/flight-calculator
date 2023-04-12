FROM golang:alpine as builder

WORKDIR /app

ADD go.mod go.mod
ADD go.sum go.sum

RUN go mod download

ADD . .

RUN go build ./cmd/server

FROM scratch

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

EXPOSE 8080

ENTRYPOINT ["/app/server"]