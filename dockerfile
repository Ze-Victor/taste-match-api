FROM golang:1.24-alpine as builder 

RUN apk add --no-cache git
RUN apk add build-base

WORKDIR /
COPY .netrc root/.netrc
WORKDIR /build
COPY . .
RUN go build -o ./taste-match-api

FROM alpine:latest
WORKDIR /
COPY --from=builder /build/taste-match-api /opt/taste-match-api
ENTRYPOINT ["/opt/taste-match-api"]