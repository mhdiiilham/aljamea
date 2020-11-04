FROM golang:1.15-alpine

RUN mkdir /server

ADD . /server

WORKDIR /server

RUN go build

EXPOSE 8080

CMD [ "/server/dummy-server" ]
