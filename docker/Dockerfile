FROM golang:1.22-alpine3.19
WORKDIR /app

RUN apk --no-cache add curl
RUN curl -sSf https://atlasgo.sh | sh

RUN go install github.com/air-verse/air@latest

CMD "./docker/scripts/local.sh"
