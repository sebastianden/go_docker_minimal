# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY *.go ./

RUN GO111MODULE=off GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-w -s" -o /server

##
## Deploy
##
FROM scratch

WORKDIR /

COPY --from=build /server /server
COPY *.html ./

ENTRYPOINT ["/server"]
