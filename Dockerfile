FROM golang:1.17-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN GIN_MODE=release go build ./cmd/lodestone -o /lodestone

EXPOSE 3999

CMD [ "/lodestone" ]