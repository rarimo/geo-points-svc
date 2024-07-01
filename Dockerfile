FROM golang:1.22-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/rarimo/geo-points-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/geo-points-svc /go/src/github.com/rarimo/geo-points-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/geo-points-svc /usr/local/bin/geo-points-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["geo-points-svc"]
