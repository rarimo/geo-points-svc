FROM golang:1.19.7-alpine as buildbase

RUN apk add git build-base ca-certificates

WORKDIR /go/src/github.com/rarimo/geo-points-svc
COPY . .
RUN go mod tidy && go mod vendor
RUN CGO_ENABLED=1 GO111MODULE=on GOOS=linux go build -o /usr/local/bin/geo-points-svc /go/src/github.com/rarimo/geo-points-svc

FROM scratch
COPY --from=alpine:3.9 /bin/sh /bin/sh
COPY --from=alpine:3.9 /usr /usr
COPY --from=alpine:3.9 /lib /lib

COPY --from=buildbase /usr/local/bin/geo-points-svc /usr/local/bin/geo-points-svc
COPY --from=buildbase /go/src/github.com/rarimo/geo-points-svc/proof_keys/passport.json /proof_keys/passport.json
COPY --from=buildbase /go/src/github.com/rarimo/geo-points-svc/proof_keys/poll.json /proof_keys/poll.json
COPY --from=buildbase /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["geo-points-svc"]