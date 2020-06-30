FROM gcc:10.1
WORKDIR /usr/local/src/macserial
COPY ./src/* ./
RUN gcc macserial.c -o /usr/local/bin/macserial

FROM golang:1.11
WORKDIR /usr/local/src/macserial_server
COPY ./server.go ./server.go
RUN go build server.go -o /usr/local/bin/macserial_server
ENTRYPOINT [ "/usr/local/bin/macserial_server" ]