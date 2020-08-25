FROM node:14.7-alpine

WORKDIR /app

ARG PORT=4000
ARG DIST=dist/dusanmalusev-frontend/server

COPY ./${DIST} .

ENV PORT=${PORT}

EXPOSE ${PORT}

CMD [ "node", "main.js" ]