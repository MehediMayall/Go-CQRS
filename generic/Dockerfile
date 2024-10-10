# Build go project

FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .


RUN apk add -u -t build-tools curl git && \
go build -o server ./cmd/*.go && \
apk del build-tools && \
rm -rf /var/cache/apk/*

# Runtime container

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates
COPY --from=builder /app/server /app/server

EXPOSE 3500

ENTRYPOINT [ "/app/server" ]
