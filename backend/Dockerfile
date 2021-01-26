FROM golang:1.14.2 AS build-env

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /go/src/github.com/p1ass/midare

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build .

FROM alpine:3.11.6

RUN apk add --no-cache bash ca-certificates curl

COPY --from=build-env /go/src/github.com/p1ass/midare/midare /midare
RUN chmod a+x /midare

EXPOSE 8080
CMD ["/midare"]