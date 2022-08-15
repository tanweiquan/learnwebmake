FROM golang:1.16
LABEL MAINTAINER="tanweiquan 1274681067@qq.com"
COPY . /app
WORKDIR /app
ENV GO111MODULE=on 
ENV GOPROXY="https://goproxy.cn" 
RUN go build -o webmake .
EXPOSE 8080
ENTRYPOINT ["./webmake"]
