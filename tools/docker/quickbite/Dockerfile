FROM golang:1.23-alpine

ARG TARGETPLATFORM
ARG APP_NAME=quickbite

WORKDIR /app

COPY ./bin/${TARGETPLATFORM}/${APP_NAME} /app/${APP_NAME}

RUN chmod +x /app/${APP_NAME}

ENTRYPOINT ["/app/quickbite", "server"]
