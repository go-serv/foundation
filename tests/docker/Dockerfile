FROM golang:1.19.0-buster

#RUN set -xe \
#    && apt-get update -y \
#    && apt-get install wget -y \
#      inotify-tools \
#      netcat \
#      git \
#      vim

RUN mkdir -p /app/target && mkdir -p /vendor && mkdir -p /ssl-certs
ADD ./src /app
COPY ./assets/go.mod /app
#
#RUN set -e && (\
#      cd /vendor && \
#      git clone https://github.com/go-serv/grpc-go.git && \
#      git clone https://github.com/AgentCoop/go-work.git && \
#      git clone https://github.com/go-serv/foundation.git \
#    )
RUN set -e && (\
      cd /app && \
      go get github.com/AgentCoop/go-work && \
      go get github.com/go-serv/foundation && \
      go get github.com/go-serv/grpc-go \
    )
#RUN cd /app && go mod tidy

WORKDIR /app
#ENTRYPOINT ["go", "run", "/app/target/main.go"]
ENTRYPOINT ["tail", "-f", "/dev/stdout"]