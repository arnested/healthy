FROM golang:1.12-alpine AS build-env

WORKDIR /build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk --no-cache add git=~2.20

COPY *.go go.mod go.sum /build/

RUN go version
RUN go build

FROM scratch

ENV PATH=/

COPY --from=build-env /go/src/arnested.dk/go/healthy/healthy /healthy

ENTRYPOINT ["healthy"]
