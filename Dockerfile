# Build Stage
FROM golang:1.17.6-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz


# Run Stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh  .
COPY resource/example.json ./resource/example.json
COPY resource/useCase.json ./resource/useCase.json
COPY resource/word.json ./resource/word.json

RUN ["chmod", "+x", "./start.sh"]
COPY wait-for.sh .
RUN ["chmod", "+x", "./wait-for.sh"]
COPY db/migration ./migration

 
EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]