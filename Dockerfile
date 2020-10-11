FROM golang:latest

RUN mkdir /code
ADD . /code/
WORKDIR /code

RUN go mod download
RUN go build -o main ./

CMD ["/code/main"]