FROM golang:1.16.4-alpine AS build-env

WORKDIR /build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk --no-cache add git=~2 upx=~3

COPY *.go go.mod go.sum /build/

RUN go version
RUN go build -ldflags '-s -w'
RUN upx healthy

FROM scratch

COPY --from=build-env /build/healthy /healthy

ENTRYPOINT ["/healthy"]
