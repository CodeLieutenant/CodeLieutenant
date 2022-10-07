FROM rust:1.64-alpine3.16 as build

WORKDIR /opt/api

COPY . .

RUN apk update && \
    apk upgrade && \
    apk add --no-cache musl-dev gcc

RUN cargo build --profile dist

FROM alpine:3.16 as production

COPY --from=build /opt/api/target/dist/api /bin/api

RUN apk update && \
    apk upgrade && \
    apk add --no-cache musl

# TODO: Copy Config file

ENV RUST_LOG=api=info

EXPOSE 3000
EXPOSE 6669

ENTRYPOINT [ "/bin/api" ]
