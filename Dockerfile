FROM alpine:3.22

WORKDIR /app

COPY ./bin/goapp .

CMD ["/app/goapp"]