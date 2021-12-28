FROM golang:1.17.5
COPY . /app
WORKDIR /app
RUN wget -O /etc/apt/sources.list http://mirrors.cloud.tencent.com/repo/ubuntu20_sources.list \
	&& apt-key adv --recv-keys --keyserver keyserver.ubuntu.com 871920D1991BC93C \
	&& apt-key adv --recv-keys --keyserver keyserver.ubuntu.com 3B4FE6ACC0B21F32 \
	&& apt-key adv --recv-keys --keyserver keyserver.ubuntu.com 871920D1991BC93C \
	&& apt-get update && apt-get install libwebp-dev \
	&& cd /app && go env -w GO111MODULE=on \
	&& go env -w GOPROXY=https://goproxy.cn,direct \
	&& go mod tidy \
	&& go build -o app
ENV GO_ENV prod
ENTRYPOINT ["/app/app"]