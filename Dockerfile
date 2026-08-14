FROM alpine:latest

WORKDIR /opt

RUN apk add wget make go zip
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v0.41.1/tinygo0.41.1.linux-amd64.tar.gz
RUN tar xf tinygo0.41.1.linux-amd64.tar.gz
ENV PATH="/opt/tinygo/bin:${PATH}"

WORKDIR /app/clap_dht_navidrome

CMD ["make"]