# syntax=docker/dockerfile:1
ARG BUILDPLATFORM=linux/amd64
ARG TARGETPLATFORM=linux/arm64

## Build
FROM --platform=${BUILDPLATFORM} golang:1.20-bullseye AS build

ARG GO_BUILDER_GITHUB_TOKEN=missing_token
RUN git config --global "url.https://$GO_BUILDER_GITHUB_TOKEN@github.com/".insteadOf https://github.com/
ENV GOPRIVATE=github.com/derek-test

WORKDIR $GOPATH/src/app

ENV GO111MODULE=on

COPY . .

RUN mkdir /app \
    && CGO_ENABLED=1 go build -race -o /app/main.amd64 src/cmd/main.go \
    && GOARCH=arm64 go build -o /app/main src/cmd/main.go

FROM --platform=${TARGETPLATFORM} 196233775518.dkr.ecr.us-west-2.amazonaws.com/go-base:v1
WORKDIR /app

COPY --from=build /app/main .

EXPOSE 3003
EXPOSE 50051

USER app:app
HEALTHCHECK CMD curl --fail http://localhost:3003 || exit 1   
ENTRYPOINT ["./main"]
