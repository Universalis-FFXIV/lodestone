FROM golang:1.17-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download

ENV GIN_MODE=release

RUN go build ./cmd/lodestone

EXPOSE 3999

CMD [ "/app/lodestone" ]
