---
title: centos安装 Go-Ldap-Admin
category: /第一版记录/2022-08-12
renderNumberedHeading: true
grammar_cjkRuby: true
---



#### 安装数据库
1：安装MySQL：5.7数据库

一、安装YUM Repo
1、由于CentOS 的yum源中没有mysql，需要到mysql的官网下载yum repo配置文件。

	官网地址是：https://dev.mysql.com/downloads/repo/yum/
	
最新的源下载命令：

	wget https://dev.mysql.com/get/mysql80-community-release-el7-5.noarch.rpm


2、然后进行repo的安装：

	rpm -ivh mysql80-community-release-el7-5.noarch.rpm

执行完成后会可以查看本机可用数据库版本

	yum repolist all | grep mysql
	
可以根据自己的需要，关闭，或者开启想要的版本

	yum-config-manager --disable mysql80-community
	
	关闭8.0
	
	yum-config-manager --enable mysql57-community
	
	开启5.7
	
	
	
二、使用yum命令即可完成安装

1、安装命令：

	yum install mysql-server
	
2、启动msyql：

	systemctl start mysqld 

3、获取安装时的临时密码（在第一次登录时就是用这个密码）：

	grep 'temporary password' /var/log/mysqld.log
	
 登录之后，必须要改一次密码，密码必须设置强密码
 
 	alter user 'root'@'localhost' identified by 'Eryajf@123';
	

4、倘若没有获取临时密码，则删除原来安装过的mysql残留的数据

	rm -rf /var/lib/mysql

5、再启动mysql

	systemctl start mysqld #启动MySQL

6、创建数据库

	CREATE USER 'ldap'@'localhost' IDENTIFIED BY 'Eryajf@123';
	
	CREATE DATABASE IF NOT EXISTS `go_ldap_admin` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
	
	GRANT ALL PRIVILEGES on `go_ldap_admin`.* to 'ldap'@'localhost';
	
	FLUSH privileges;

#### 2：安装openLDAP



1：使用yum方式安装

	yum install openldap openldap-clients openldap-servers
	
2:复制一个默认配置到指定目录下,并授权，这一步一定要做，然后再启动服务，不然生产密码时会报错

	cp /usr/share/openldap-servers/DB_CONFIG.example /var/lib/ldap/DB_CONFIG

3:授权给ldap用户,此用户yum安装时便会自动创建

	chown -R ldap /var/lib/ldap/DB_CONFIG
	
4:启动服务，先启动服务，配置后面再进行修改

	systemctl start slapd
	
	systemctl enable slapd
 
查看状态，正常启动则ok

	systemctl status slapd
	
5:修改openldap配置

这里就是重点中的重点了，从openldap2.4.23版本开始，所有配置都保存在`/etc/openldap/slapd.d`目录下的`cn=config`文件夹内，不再使用`slapd.conf`作为配置文件。配置文件的后缀为 `ldif`，且每个配置文件都是通过命令自动生成的，任意打开一个配置文件，在开头都会有一行注释，说明此为自动生成的文件，请勿编辑，使用ldapmodify命令进行修改

---

6:安装openldap后，会有三个命令用于修改配置文件，分别为ldapadd, ldapmodify, ldapdelete，顾名思义就是添加，修改和删除。而需要修改或增加配置时，则需要先写一个ldif后缀的配置文件，然后通过命令将写的配置更新到slapd.d目录下的配置文件中去

---

7：初始化配置

生成管理员密码,记录下这个密码，{SSHA}这一串，后面需要用到

	slappasswd -s 123456
	
	{SSHA}LSgYPTUW4zjGtIVtuZ8cRUqqFRv1tWpE
 
新增修改密码文件,ldif为后缀，文件名随意，不要在/etc/openldap/slapd.d/目录下创建类似文件
生成的文件为需要通过命令去动态修改ldap现有配置，如下，我在家目录下，创建文件

	cd ~
	
	vim changepwd.ldif
	----------------------------------------------------------------------
	dn: olcDatabase={0}config,cn=config
	changetype: modify
	add: olcRootPW
	olcRootPW: {SSHA}LSgYPTUW4zjGtIVtuZ8cRUqqFRv1tWpE
	----------------------------------------------------------------------
	# 这里解释一下这个文件的内容：
	# 第一行执行配置文件，这里就表示指定为 cn=config/olcDatabase={0}config 文件。
	# 你到/etc/openldap/slapd.d/目录下就能找到此文件
	# 第二行 changetype 指定类型为修改
	# 第三行 add 表示添加 olcRootPW 配置项
	# 第四行指定 olcRootPW 配置项的值
	# 在执行下面的命令前，你可以先查看原本的olcDatabase={0}config文件，
	# 里面是没有olcRootPW这个项的，执行命令后，你再看就会新增了olcRootPW项，
	# 而且内容是我们文件中指定的值加密后的字符串
 
执行命令，修改ldap配置，通过-f执行文件

	ldapadd -Y EXTERNAL -H ldapi:/// -f changepwd.ldif
	

查看`olcDatabase={0}config`内容,新增了一个`olcRootPW`项。

	cat /etc/openldap/slapd.d/cn=config/olcDatabase={0}config.ldif

#### 切记不能直接修改/etc/openldap/slapd.d/目录下的配置。

---

我们需要向 LDAP 中导入一些基本的 Schema。这些 Schema 文件位于`/etc/openldap/schema/`目录中，schema控制着条目拥有哪些对象类和属性，可以自行选择需要的进行导入，

	依次执行下面的命令，导入基础的一些配置,我这里将所有的都导入一下，
	其中core.ldif是默认已经加载了的，不用导入
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/cosine.ldif

	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/nis.ldif
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/inetorgperson.ldif
	
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/collective.ldif
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/corba.ldif
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/duaconf.ldif
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/dyngroup.ldif
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/java.ldif
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/misc.ldif
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/openldap.ldif
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/pmi.ldif
	ldapadd -Y EXTERNAL -H ldapi:/// -f /etc/openldap/schema/ppolicy.ldif
	
#### 修改域名，新增con.ldif, 
这里我自定义的域名为 eryajf.net，管理员用户账号为admin。
如果要修改，则修改文件中相应的dc=eryajf,dc=net为自己的域名，密码切记改成上面重新生成的哪个。

	vim con.ldif
	
---
	
	dn: olcDatabase={1}monitor,cn=config
	changetype: modify
	replace: olcAccess
	olcAccess: {0}to * by dn.base="gidNumber=0+uidNumber=0,cn=peercred,cn=external,cn=auth" read by dn.base="cn=admin,dc=yaobili,dc=com" read by * none
 
	dn: olcDatabase={2}hdb,cn=config
	changetype: modify
	replace: olcSuffix
	olcSuffix: dc=eryajf,dc=net
 
	dn: olcDatabase={2}hdb,cn=config
	changetype: modify
	replace: olcRootDN
	olcRootDN: cn=admin,dc=eryajf,dc=net
 
	dn: olcDatabase={2}hdb,cn=config
	changetype: modify
	replace: olcRootPW
	olcRootPW: {SSHA}LSgYPTUW4zjGtIVtuZ8cRUqqFRv1tWpE
 
	dn: olcDatabase={2}hdb,cn=config
	changetype: modify
	add: olcAccess
	olcAccess: {0}to attrs=userPassword,shadowLastChange by dn="cn=admin,dc=eryajf,dc=net" write by anonymous auth by self write by * none
	olcAccess: {1}to dn.base="" by * read
	olcAccess: {2}to * by dn="cn=admin,dc=eryajf,dc=net" write by * read

执行命令，修改配置

	ldapmodify -Y EXTERNAL -H ldapi:/// -f con.ldif
	
最后这里有5个修改，所以执行会输出5行表示成功。

然后，启用memberof功能

新增add-memberof.ldif, #开启memberof支持并新增用户支持memberof配置

	# 新增add-memberof.ldif, #开启memberof支持并新增用户支持memberof配置
	
	vim add-memberof.ldif
-------------------------------------------------------------
	dn: cn=module{0},cn=config
	cn: modulle{0}
	objectClass: olcModuleList
	objectclass: top
	olcModuleload: memberof.la
	olcModulePath: /usr/lib64/openldap
 
	dn: olcOverlay={0}memberof,olcDatabase={2}hdb,cn=config
	objectClass: olcConfig
	objectClass: olcMemberOf
	objectClass: olcOverlayConfig
	objectClass: top
	olcOverlay: memberof
	olcMemberOfDangling: ignore
	olcMemberOfRefInt: TRUE
	olcMemberOfGroupOC: groupOfUniqueNames
	olcMemberOfMemberAD: uniqueMember
	olcMemberOfMemberOfAD: memberOf
-------------------------------------------------------------
	# 新增refint1.ldif文件
	vim refint1.ldif
-------------------------------------------------------------
	dn: cn=module{0},cn=config
	add: olcmoduleload
	olcmoduleload: refint
------------------------------------------------------------- 
 
	# 新增refint2.ldif文件
	vim refint2.ldif
-------------------------------------------------------------
	dn: olcOverlay=refint,olcDatabase={2}hdb,cn=config
	objectClass: olcConfig
	objectClass: olcOverlayConfig
	objectClass: olcRefintConfig
	objectClass: top
	olcOverlay: refint
	olcRefintAttribute: memberof uniqueMember  manager owner
-------------------------------------------------------------
 
 

#### 依次执行下面命令，加载配置，顺序不能错

	ldapadd -Q -Y EXTERNAL -H ldapi:/// -f add-memberof.ldif
---

	ldapmodify -Q -Y EXTERNAL -H ldapi:/// -f refint1.ldif
---
	
	ldapadd -Q -Y EXTERNAL -H ldapi:/// -f refint2.ldif
---
	
#### 创建一个组织
到此，配置修改完了,在上述基础上，我们来创建一个叫做 eryajf company 的组织，并在其下创建一个 admin 的组织角色（该组织角色内的用户具有管理整个 LDAP 的权限）和 People 和 Group 两个组织单元：

	# 新增配置文件
	
	vim base.ldif
----------------------------------------------------------
	dn: dc=eryajf,dc=net
	objectClass: top
	objectClass: dcObject
	objectClass: organization
	o: Yaobili Company
	dc: yaobili
 
	dn: cn=admin,dc=eryajf,dc=net
	objectClass: organizationalRole
	cn: admin
 
	dn: ou=People,dc=eryajf,dc=net
	objectClass: organizationalUnit
	ou: People
 
	dn: ou=Group,dc=eryajf,dc=net
	objectClass: organizationalRole
	cn: Group
----------------------------------------------------------
 
这里注意，可以把eryajf和net改成自己的 
 
 执行命令，添加配置, 这里要注意修改域名为自己配置的域名，然后需要输入上面我们生成的密码

	ldapadd -x -D cn=admin,dc=yaobili,dc=com -W -f base.ldif

#### 3：安装OpenResty

配置阿里和openresty的yum源：

	$ yum -y install yum-utils
	
	$ yum-config-manager --add-repo http://mirrors.aliyun.com/repo/Centos-7.repo
	
	$ yum-config-manager --add-repo https://openresty.org/package/centos/openresty.repo


安装

	yum install openresty

创建一些软链，便于维护或者规范

	ln -snf /usr/local/openresty/nginx/sbin/nginx /usr/sbin/nginx
	
	ln -snf /usr/local/openresty/nginx/conf /etc/nginx
	
启动服务

	systemctl start openresty
	
	systemctl status openresty

#### 4：安装go开发环境

先下载golang，下载地址：

	https://studygolang.com/dl

上传至`/usr/local`目录下并解压

添加环境变量

修改profile文件：

修 改`/etc/profile`（对所有用户都是有效的）

	vi /etc/profile

在里面加入:

	export GO_HOME=/usr/local/go
	export PATH=$PATH:$GO_HOME/bin

使配置生效

	source /etc/profile
	
#### 5:安装安装nodejs、npm、yarn

1：安装epel源

	yum install epel-release -y

2：安装nodejs和npm

	yum install nodejs npm -y
	
	启用Yarn存储库并导入存储库的GPG密钥，请执行以下命令：

	curl --silent --location https://dl.yarnpkg.com/rpm/yarn.repo | sudo tee /etc/yum.repos.d/yarn.repo
	
	
	sudo rpm --import https://dl.yarnpkg.com/rpm/pubkey.gpg
	
	
	yum install yarn
	
	4:设置淘宝镜像和安装淘宝cnpm

	npm config set registry https://registry.npm.taobao.org
	
	npm install -g cnpm --registry=https://registry.npm.taobao.org
	
#### 6：部署后端

检索项目到/usr/local

	git clone https://gitee.com/eryajf-world/go-ldap-admin.git
	
进入目录，编译项目
	
	$ make build-linux

更改配置，根据实际情况调整配置文件内容

	vim config.yml
	
修改数据库信息和LDAP连接信息


systemd管理，基于systemd进行管理

	$cat /usr/lib/systemd/system/go-ldap-admin.service
---
	[Unit]
	Description=Go Ldap Admin Service

	[Service]
	WorkingDirectory=/data/www/go-ldap-admin.eryajf.net/
	ExecStart=/data/www/go-ldap-admin.eryajf.net/go-ldap-admin

	[Install]
	WantedBy=multi-user.target


启动项目

	$ systemctl daemon-reload
	
	$ systemctl start go-ldap-admin
	
	$ systemctl status go-ldap-admin
	
#### 7：部署前端

检索项目到本地/usr/local

	git clone https://gitee.com/eryajf-world/go-ldap-admin-ui.git

前端通过OpenResty代理的方式进行部署

编译项目之前，需要将`.env.production`中的`VUE_APP_BASE_API`配置项，更改为正式部署环境的域名,或者`本机的IP地址`不含端口号。

	VUE_APP_BASE_API = 'http://demo-go-ldap-admin.eryajf.net/'
	
---
	编译项目,一些直接从GitHub拉取的依赖需要进行如下配置
---	

	git config --global url."https://".insteadOf git://
	
	npm install --registry=http://registry.npmmirror.com
	
	yarn build:prod
	
	
#### 8:在OpenResty中添加如下配置，代理本项目

	$ cat /etc/nginx/conf/nginx.conf
	
	
---
	worker_processes  1;
	events {
    worker_connections  1024;
	}

	http {
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
 
    keepalive_timeout  65;
    server {
    listen 80;
    server_name demo-go-ldap-admin.eryajf.net;
    root /usr/local/go-ldap-admin-ui/dist;
    location / {
        try_files $uri $uri/ /index.html;
        add_header Cache-Control 'no-store';
    }

    location /api/ {
        proxy_set_header Host $http_host;
        proxy_set_header  X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_pass http://127.0.0.1:8888;
    }
	}
	}

检查一下几个服务是否都开启


数据库

systemctl status mysqld 

LDAP

systemctl status slapd

后端

systemctl status go-ldap-admin

代理

systemctl status openresty

都启动之后，就可以打开web管理页面

http:ip:端口

默认用户名：admin，默认密码是LDAP设置的密码。