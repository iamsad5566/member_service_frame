# Step 1: Test
FROM golang:latest AS tester

ARG GITHUB_TOKEN

ENV GOPRIVATE=github.com/iamsad5566/*

RUN apt-get update && apt-get install -y git
RUN echo "machine github.com login $GITHUB_TOKEN password x-oauth-basic" > ~/.netrc

WORKDIR /app
RUN git clone https://github.com/iamsad5566/setconf.git
RUN ls ./setconf
RUN cp ./setconf/config.yml .
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./...

# Step 2: Build
FROM golang:latest AS builder

ARG GITHUB_TOKEN
ENV GOPRIVATE=github.com/iamsad5566/*
ENV CGO_ENABLED=0

RUN echo "machine github.com login $GITHUB_TOKEN password x-oauth-basic" > ~/.netrc

WORKDIR /app
COPY --from=tester /app/go.mod /app/go.sum ./
RUN go mod download
COPY --from=tester /app/ .
RUN go build -o member_service

# Step 3: Run
FROM alpine:latest
ARG LATEST_SETCONF_VERSION
ENV ENVIRONMENT=prod

WORKDIR /app
RUN mkdir -p /go/pkg/mod/github.com/iamsad5566/setconf@$LATEST_SETCONF_VERSION
COPY --from=builder /app/member_service .
COPY --from=builder /app/config.yml /go/pkg/mod/github.com/iamsad5566/setconf@$LATEST_SETCONF_VERSION/.

EXPOSE 8080
CMD ["./member_service"]