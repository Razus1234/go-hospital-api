# builder
FROM golang:alpine AS builder
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download
COPY . .
RUN go build -o /app/api ./

# final image
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/api /usr/local/bin/api
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/api"]