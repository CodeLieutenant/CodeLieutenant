FROM golang:1.17

COPY . /build

RUN cd /build \
    && apt install gcc g++ make -y \
    && make -j build DOCKER=1 \
    && mkdir -p /var/log/app \
    && mkdir -p /etc/app \
    && rm -rf ./bin/migrations \
    && rm -rf ./bin/views \
    && mv config.yml /etc/app/config.yml \
    && cp -r ./bin /app

WORKDIR /app

ENV DEBUG=false
ENV LOG_LEVEL=debug

EXPOSE 8000

CMD ["./dusanmalusev", "--level", "${LOG_LEVEL}", "--config", "/etc/app/config.yml", "server", "--debug", "${DEBUG}", "--address", "0.0.0.0:8000"]
