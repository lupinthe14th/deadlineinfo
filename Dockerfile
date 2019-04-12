FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/lupinthe14th/deadlineinfo
COPY . /usr/local/go/src/github.com/lupinthe14th/deadlineinfo
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/deadlineinfo ./deadlineinfo


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/lupinthe14th/deadlineinfo/build/deadlineinfo /bin/deadlineinfo
CMD ["deadlineinfo", "up"]
