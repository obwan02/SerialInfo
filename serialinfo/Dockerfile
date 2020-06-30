FROM python:3
WORKDIR /usr/local/src/serialinfo
COPY ./main.py ./main.py
COPY ./requirements.txt ./requirements.txt
RUN python -m pip install -r ./requirements.txt
ENTRYPOINT [ "python", "/usr/local/src/serialinfo/main.py" ]