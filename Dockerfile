FROM golang:1.12-alpine AS build-env

WORKDIR /go/src/arnested.dk/go/healthy
COPY *.go /go/src/arnested.dk/go/healthy/
COPY vendor /go/src/arnested.dk/go/healthy/vendor/

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build

FROM scratch

ENV PATH=/

COPY --from=build-env /go/src/arnested.dk/go/healthy/healthy /healthy

ENTRYPOINT ["healthy"]
