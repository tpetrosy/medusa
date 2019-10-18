
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir -p /opt/medusaweb
COPY --from=builder /go/bin/app /opt/medusaweb/app
COPY internal/templates /opt/medusaweb/templates
WORKDIR /opt/medusaweb/
ENTRYPOINT /opt/medusaweb/app
LABEL Name=medusaweb Version=0.0.1
EXPOSE 8080
