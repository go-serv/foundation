FROM golang:1.19.0-buster

#RUN set -xe \
#    && apt-get update -y \
#    && apt-get install wget -y \
#      inotify-tools \
#      netcat \
#      git \
#      vim

RUN mkdir -p /app/target && mkdir -p /vendor
ADD ./src /app
COPY ./assets/go.mod /app

RUN set -e && (\
      cd /vendor && \
      git clone https://github.com/go-serv/grpc-go.git && \
      git clone https://github.com/AgentCoop/go-work.git && \
      git clone https://github.com/go-serv/foundation.git \
    )
RUN cd /app && go mod tidy

RUN chmod -R 775 /share
WORKDIR /app
ENTRYPOINT ["go", "run", "/app/target/main.go"]