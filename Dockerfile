FROM golang:1.17.10 AS builder

# ENV GOPROXY      https://goproxy.io

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN sed -i 's@host: localhost@host: mysql@g'  config.yml \
    && sed -i 's@localhost:389@openldap:389@g' config.yml \
    && go build -o go-ldap-admin .

FROM centos:centos7
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/wait .
COPY --from=builder /app/ .
RUN chmod +x wait go-ldap-admin && yum -y install vim net-tools telnet wget curl && yum clean all
CMD ./wait && ./go-ldap-admin