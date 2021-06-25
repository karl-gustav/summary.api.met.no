# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.15-alpine as builder

RUN apk update
# Add required certificates to be able to call HTTPS endpoints.
RUN apk add --no-cache ca-certificates git tzdata

# Copy local code to the container image.
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .

# CGO_ENABLED=0 == Don't depend on libc (bigger but more independent binary)
# installsuffix == Cache dir for non cgo build files
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -installsuffix 'static' -o main

FROM scratch
WORKDIR /app

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo/Europe/Oslo /usr/share/zoneinfo/Europe/Oslo

COPY --from=builder /app/main .

CMD ["./main"]
