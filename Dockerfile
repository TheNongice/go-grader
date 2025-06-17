# This Dockerfile is not fully complete. You can contribute to my repo with PR!
FROM golang:tip-20250613-bookworm
WORKDIR /src
# Isolate IOI
RUN curl https://www.ucw.cz/isolate/debian/signing-key.asc >/etc/apt/keyrings/isolate.asc
RUN apt-get update
RUN apt-get install -y git libcap-dev pkg-config build-essential libsystemd-dev
RUN git clone https://github.com/ioi/isolate.git

WORKDIR /src/isolate
RUN make install

# Go's Service
WORKDIR /src/app
RUN git clone https://github.com/TheNongice/go-grader.git .
RUN mkdir problem \
    && mkdir runner \
    && mkdir runner/isolate_logs \
    && mkdir runner/temp_code \
    && mkdir runner/temp_problem \
    && mkdir runner/temp_code/output \
    && mkdir runner/temp_code/cpp \
    && mkdir runner/temp_code/cpp/output

RUN echo "DIR_GRADER_PATH=/src/app/" >> .env \
    && echo "ISOLATE_PATH=/usr/local/bin/isolate" >> .env

RUN go mod tidy
RUN go build -o go-grader .
EXPOSE 8000
CMD ["/src/app/go-grader"]