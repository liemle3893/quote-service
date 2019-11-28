FROM golang:1.13.0-stretch AS builder

ENV GO111MODULE=on
# Build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -installsuffix cgo -mod=vendor -o /app main.go

# Main image
FROM scratch
EXPOSE 1323
COPY --from=builder /app /
ENTRYPOINT ["/app"]




