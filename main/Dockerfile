FROM golang:1.13-alpine as builder
ARG VERSION=0.0.1
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
#ENV GOARCH=arm

# build
WORKDIR /go/src/network-scheduler/main
COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on go mod download
COPY . .
RUN go install -ldflags "-s -w -X main.version=$VERSION" network-scheduler

# runtime image
FROM gcr.io/google_containers/ubuntu-slim:0.14
COPY --from=builder /go/bin/network-scheduler /usr/bin/network-scheduler
EXPOSE 8100
ENTRYPOINT ["network-scheduler"]
