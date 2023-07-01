#start building this image from this source (OS, etc...)
FROM golang:1.19-alpine3.17

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

ENV APP_HOME /app

WORKDIR "$APP_HOME"

COPY go.mod ./

RUN go mod download

COPY . .

RUN ["go", "get", "github.com/swaggo/swag/cmd/swag"]
RUN ["go", "install", "github.com/swaggo/swag/cmd/swag@latest"]

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -polling -log-prefix=false -exclude-dir="./docs" -build="swag init -g ./cmd/main.go --output ./docs && go build -o ./tmp/main ./cmd" -command="./tmp/main"