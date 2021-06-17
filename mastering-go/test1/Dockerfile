FROM golang:alpine

RUN mkdir /files
COPY hello.go /files
WORKDIR /files

RUN go build -o /files/hello hello.go
ENTRYPOINT ["/files/hello"]

