FROM alpine:edge

WORKDIR /app

ARG PORT=3000

COPY ./main .

ENV ADDR=":${PORT}"
ENV PORT=${PORT}

EXPOSE ${PORT}

CMD [ "/app/main" ]