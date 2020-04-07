FROM golang:latest
RUN go env -w GO111MODULE=on GOPROXY=https://goproxy.cn,direct
WORKDIR /gomod
COPY . .
EXPOSE 8888
CMD ["/bin/bash","build.sh"]