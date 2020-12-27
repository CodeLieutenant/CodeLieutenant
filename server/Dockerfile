FROM brossquad/fiber-dev:1.0.3 as builder
COPY . /app
WORKDIR /app
RUN task build-prod


FROM alpine:3.12 as production
ARG PORT=4000
COPY --from=builder /app/build /app
WORKDIR /app
EXPOSE ${PORT}
CMD ["./app"]

