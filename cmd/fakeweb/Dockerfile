
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN ls -la
RUN go version
RUN go get -d -v ./...
RUN go install -v ./...
RUN ls -la /go/bin

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir -p /opt/app
COPY --from=builder /go/bin/fakeweb /opt/app/fakeweb
COPY internal/templates /opt/app/internal/templates
WORKDIR /opt/app/
ENTRYPOINT /opt/app/fakeweb
LABEL Name=fakeweb Version=0.0.1
EXPOSE 8080
