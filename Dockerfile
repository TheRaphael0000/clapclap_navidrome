FROM --platform=$BUILDPLATFORM alpine:latest

ARG TARGETOS
ARG TARGETARCH
ENV TINYGO_VERSION=0.41.1

WORKDIR /opt

RUN apk add wget make go zip
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo${TINYGO_VERSION}.${TARGETOS}-${TARGETARCH}.tar.gz
RUN tar xf tinygo${TINYGO_VERSION}.${TARGETOS}-${TARGETARCH}.tar.gz
ENV PATH="/opt/tinygo/bin:${PATH}"

WORKDIR /app/clap_dht_navidrome

CMD ["make"]