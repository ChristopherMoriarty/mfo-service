FROM postgres:17.5-alpine

RUN apk update && \
    apk upgrade && \
    rm -rf /var/cache/apk/*

RUN chown -R postgres:postgres /var/lib/postgresql && \
    chmod 700 /var/lib/postgresql/data

EXPOSE 5432

CMD ["postgres"]
