FROM alpine:3.3

COPY ./dist/http-fs /dist/http-fs

ENV ENV=PRODUCTION
WORKDIR /dist/
ENTRYPOINT ["/dist/http-fs"]