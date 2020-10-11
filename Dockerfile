FROM golang:latest

RUN mkdir /code
ADD . /code/
WORKDIR /code

RUN GOPROXY="https://goproxy.cn/" GO111MODULE=on go build -o main ./

CMD ["/code/main"]