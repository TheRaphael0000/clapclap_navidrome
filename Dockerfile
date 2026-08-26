FROM --platform=$BUILDPLATFORM alpine:latest

ARG TARGETOS
ARG TARGETARCH
ENV TINYGO_VERSION=0.41.1

WORKDIR /opt

# make and go for building the app, zip for bundling the app
RUN apk add curl make go zip
RUN curl -sL https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo${TINYGO_VERSION}.${TARGETOS}-${TARGETARCH}.tar.gz | tar -xzf -
ENV PATH="/opt/tinygo/bin:${PATH}"

WORKDIR /app/clapclap_navidrome

CMD ["make"]