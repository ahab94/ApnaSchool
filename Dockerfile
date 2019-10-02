FROM alpine:3.6

RUN apk add --no-cache \
        ca-certificates \
        bash \
    && rm -f /var/cache/apk/*

COPY bin/apnaschool /usr/local/bin/apnaschool

CMD ["/usr/local/bin/apnaschool"]
