FROM golang:1.13.6-alpine AS build-env

WORKDIR /build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk --no-cache add git=~2

COPY *.go go.mod go.sum /build/

RUN go version
RUN go build

FROM scratch

ENV PATH=/

COPY --from=build-env /build/healthy /healthy

ENTRYPOINT ["healthy"]
