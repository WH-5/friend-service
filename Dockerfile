FROM golang:1.23 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        netbase && \
    rm -rf /var/lib/apt/lists/* && \
    apt-get autoremove -y && \
    apt-get autoclean -y

COPY --from=builder /src/bin /app


#用这个文件里的配置文件
COPY data/conf /app/configs



WORKDIR /app

EXPOSE 8001
EXPOSE 9001
VOLUME /data/conf
#RUN apt-get update && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*
#
#COPY entrypoint.sh /entrypoint.sh
#RUN chmod +x /entrypoint.sh
#ENTRYPOINT ["/entrypoint.sh"]
CMD ["./friend-service", "-conf", "/app/configs"]
