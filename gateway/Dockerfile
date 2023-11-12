# Step 1: Modules caching
FROM golang:1.20.11-alpine3.17 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.20.11-alpine3.17 as builder
COPY --from=modules /go/pkg /go/pkg
COPY ./config /app/config
COPY ./proto /app/proto
COPY ./pkg /app/pkg
COPY ./api_gateway /app/api_gateway
COPY go.mod go.sum /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./api_gateway/cmd

# Step 3: Final
FROM scratch
COPY --from=builder /app/config /config
COPY --from=builder /bin/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app"]