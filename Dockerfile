FROM golang:1.21 AS builder
ARG GITHUB_ACCESS_TOKEN
WORKDIR /go/src/app/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl tzdata
ENV TZ=Asia/Tokyo
WORKDIR /usr/local/bin/
COPY --from=builder /go/src/app/appctl .
CMD ["/usr/local/bin/appctl", "serve"]
