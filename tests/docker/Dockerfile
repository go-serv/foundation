FROM golang:1.19.0-buster

COPY ./assets/certs /share
RUN mkdir -p /app
ADD ./src /app
RUN cd /app && go mod tidy
WORKDIR /app
RUN ["go", "run", "main.go"]