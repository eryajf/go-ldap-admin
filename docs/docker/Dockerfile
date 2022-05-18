FROM eryajf/openresty:1.21.4-centos7

ADD . /
ADD default.conf /etc/nginx/conf.d/default.conf
RUN  chmod +x /go-ldap-admin run.sh

EXPOSE 80
EXPOSE 8888

CMD /wait && /run.sh