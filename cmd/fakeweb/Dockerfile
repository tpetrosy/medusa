
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN ls -la
RUN go version
RUN ls -la ./internal
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir -p /opt/fakeweb
COPY --from=builder /go/bin/app /opt/fakeweb/app
COPY internal/templates /opt/fakeweb/internal/templates
WORKDIR /opt/fakeweb/
ENTRYPOINT /opt/fakeweb/app
LABEL Name=fakeweb Version=0.0.1
EXPOSE 8080
