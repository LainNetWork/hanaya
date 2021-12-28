FROM golang:1.17.5-alpine
COPY . /app
WORKDIR /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
	&& apk add build-base
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk add libwebp=1.2.1-r0 \
	&& cd /app && go env -w GO111MODULE=on \
	&& go env -w GOPROXY=https://goproxy.cn,direct \
	&& go mod tidy \
	&& go build -o app
ENV GO_ENV prod
ENTRYPOINT ["/app/app"]