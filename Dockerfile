FROM golang:1.17.10 AS builder

# ENV GOPROXY      https://goproxy.io

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o go-ldap-admin .

FROM centos:centos7
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/ .
RUN chmod +x wait go-ldap-admin docker-start.sh && yum -y install vim net-tools telnet wget curl && yum clean all

CMD [ "sh", "-c", "docker-start.sh" ]