FROM golang:{{.GoVersion}} as build-env

ENV GO111MODULE=on  \
    CGO_ENABLED=0   \
    GOOS=linux      \
    GOARCH=amd64

WORKDIR /build
COPY cmd/ cmd/
COPY go.mod .
COPY go.sum .
ARG VERSION
RUN go build -o app -ldflags "-X main.AppVersion=${VERSION}" cmd/*.go

FROM alpine
WORKDIR /go
WORKDIR /go/bin
COPY --from=build-env /build/app .
COPY application.yaml /go/bin/application.yaml
RUN chmod +x app

ENTRYPOINT ["/go/bin/app"] 
