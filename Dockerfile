FROM gcc:10.1
WORKDIR /usr/local/src/macserial
COPY ./macserial ./macserial
RUN gcc -v macserial.c -o /usr/local/bin/macserial

FROM python:3
WORKDIR /usr/local/src/serialinfo
COPY ./main.py ./main.py
COPY ./requirements.txt ./requirements.txt
ENTRYPOINT [ "python", "/usr/local/src/serialinfo/main.py" ]