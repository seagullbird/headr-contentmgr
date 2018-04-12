FROM alpine

ENV POSTGRESQL_ARGS="host=postgresql-postgresql port=5432 user=postgres dbname=postgres password=qBDXNlz276 sslmode=disable"

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY contentmgr /

ENTRYPOINT /contentmgr