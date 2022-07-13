#！/bin/bash

# 等待依赖项初始化成功
/app/wait

# 修改配置文件中的连接地址
sed -i 's@localhost:389@openldap:389@g' /app/config.yml
sed -i 's@host: localhost@host: mysql@g'  /app/config.yml

# 启动服务
/app/go-ldap-admin