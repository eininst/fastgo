FROM registry.cn-zhangjiakou.aliyuncs.com/eininst/go:base_v1

WORKDIR /app
COPY . /app

ARG APP=""
RUN go build -o run $APP

ENTRYPOINT ["./run"]