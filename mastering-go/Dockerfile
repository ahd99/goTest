FROM golang:alpine

RUN mkdir /files
COPY errorTest.go /files
WORKDIR /files
RUN go build -o /files/mytest errorTest.go
ENTRYPOINT ["/files/mytest"]