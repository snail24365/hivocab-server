# Build Stage
FROM golang:1.16.13-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz


# Run Stage
FROM alpine:3.15
ENV TZ=Asia/Seoul
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

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

ADD zoneinfo.zip /
ENV ZONEINFO /zoneinfo.zip

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]