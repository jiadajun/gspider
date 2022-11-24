FROM golang:1.16

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /code

RUN go env -w GO111MODULE=on

RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

ADD go.mod .

RUN go mod download

COPY . .

CMD "./run.sh"