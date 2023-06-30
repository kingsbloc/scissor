FROM golang:1.18-bullseye as base

FROM base as dev

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install --no-install-recommends -y libvips-dev && apt-get clean

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

ENV APP_HOME /app
RUN mkdir "$APP_HOME"


ADD . "$APP_HOME"

WORKDIR "$APP_HOME"

COPY . "$APP_HOME"

# ADD docker-entrypoint.sh "$APP_HOME"

# RUN go mod tidy

EXPOSE 5000

# ENTRYPOINT ["sh", "/app/docker-entrypoint.sh"]