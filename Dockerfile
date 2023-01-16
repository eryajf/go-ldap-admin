FROM golang:1.17.13-alpine3.15  AS builder

# ENV GOPROXY      https://goproxy.io

RUN mkdir /app && apk add --no-cache --virtual .build-deps \
    ca-certificates \
    gcc \
    g++

ADD . /app/

WORKDIR /app

RUN sed -i 's@localhost:389@openldap:389@g' /app/config.yml \
    && sed -i 's@host: localhost@host: mysql@g'  /app/config.yml && go build -o go-ldap-admin .

### build final image
FROM alpine:3.15

# we set the timezone `Asia/Shanghai` by default, you can be modified
# by `docker build --build-arg="TZ=Other_Timezone ..."`
ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/ .


RUN apk upgrade \
    && apk add bash tzdata \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone

RUN chmod +x wait go-ldap-admin

CMD ./wait && ./go-ldap-admin