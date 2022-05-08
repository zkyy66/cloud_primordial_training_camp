FROM golang:1.18 AS build
LABEL maintainer="crazy_cat <ages521you@hotmail.com>"
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM busybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 8080
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]