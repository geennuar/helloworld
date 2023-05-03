FROM alpine:latest

ENV TZ=Asia/Shanghai

RUN apk update \
    && apk add tzdata \
    && echo "${TZ}" > /etc/timezone \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && rm /var/cache/apk/*

ADD helloworld /helloworld
ENTRYPOINT ["/helloworld"]