FROM golang:1.18

LABEL maintainer="crazy_cat <ages521you@hotmail.com>"

ENV GOPROXY https://goproxy.cn,direct
ENV CODEPATH /usr/src/code
# 安装必要的软件包和依赖包
USER root
RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    curl \
    zip \
    unzip \
    git \
    vim

# 安装 goctl
USER root
# GO111MODULE=on
RUN  GOPROXY=https://goproxy.cn,direct go install github.com/tal-tech/go-zero/tools/goctl@cli
#github.com/tal-tech/go-zero/tools/goctl@latest

# $GOPATH/bin添加到环境变量中
ENV PATH $GOPATH/bin:$PATH

# 清理垃圾
USER root
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

# 设置工作目录
WORKDIR CODEPATH
#copyCode
COPY . .
#构建
RUN go build -o httpServer .
EXPOSE 8080
ENTRYPOINT ["./httpServer"]
#COPY ./httpserver /