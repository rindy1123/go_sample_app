FROM golang:1.22-alpine3.19
WORKDIR /app

RUN apk --no-cache add curl
RUN curl -sSf https://atlasgo.sh | sh

CMD ["go", "run", "./cmd/server.go"]
