FROM golang:1.21 AS builder
WORKDIR /go/src/app/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl tzdata
ENV TZ=Asia/Tokyo
WORKDIR /usr/local/bin/
COPY --from=builder --chown=nobody:nogroup /go/src/app/app .
USER nobody
CMD ["/usr/local/bin/app", "serve"]
