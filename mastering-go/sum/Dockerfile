FROM golang:alpine

RUN mkdir /app
COPY argsum.go /app
WORKDIR /app
RUN go build -o /app/argsum argsum.go
ENTRYPOINT [ "/app/argsum" ]