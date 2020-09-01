FROM alpine:edge

WORKDIR /app

ARG PORT=3000

COPY ./main .

ENV ADDR=":${PORT}"

EXPOSE ${PORT}

CMD [ "main" ]