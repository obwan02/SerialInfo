FROM gcc:10.1 as macserial
WORKDIR /usr/src/macserial
COPY ./macserialsrc/* ./
RUN gcc macserial.c -o /usr/bin/macserial

FROM golang:1.14
WORKDIR /usr/src/macserial_server
COPY ./gosrc/* ./
RUN go build -o /usr/bin/macserial_server
COPY --from=macserial /usr/bin/macserial /usr/bin/macserial
ENTRYPOINT [ "/usr/bin/macserial_server" ]