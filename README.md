<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**目录**

- [Go-Ldap-Admin](#go-ldap-admin)
  - [缘起](#%E7%BC%98%E8%B5%B7)
  - [在线体验](#%E5%9C%A8%E7%BA%BF%E4%BD%93%E9%AA%8C)
  - [项目地址](#%E9%A1%B9%E7%9B%AE%E5%9C%B0%E5%9D%80)
  - [核心功能](#%E6%A0%B8%E5%BF%83%E5%8A%9F%E8%83%BD)
  - [快速开始](#%E5%BF%AB%E9%80%9F%E5%BC%80%E5%A7%8B)
  - [本地开发](#%E6%9C%AC%E5%9C%B0%E5%BC%80%E5%8F%91)
  - [生产部署](#%E7%94%9F%E4%BA%A7%E9%83%A8%E7%BD%B2)
  - [感谢](#%E6%84%9F%E8%B0%A2)
  - [另外](#%E5%8F%A6%E5%A4%96)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

<h1 align="center">Go-Ldap-Admin</h1>

<div align="center">
基于Go+Vue实现的openLDAP后台管理项目。
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/eryajf-world/go-ldap-admin" alt="Go version"/>
<img src="https://img.shields.io/badge/Gin-1.6.3-brightgreen" alt="Gin version"/>
<img src="https://img.shields.io/badge/Gorm-1.20.12-brightgreen" alt="Gorm version"/>
<img src="https://img.shields.io/github/license/eryajf-world/go-ldap-admin" alt="License"/>
</p>
</div>

## 缘起

我曾经经历的公司强依赖openLDAP来作为企业内部员工管理的平台，并通过openLDAP进行各平台的认证打通工作。

但成也萧何败也萧何，给运维省力的同时，ldap又是维护不够友好的。

在[godap](https://github.com/bradleypeabody/godap)项目中，作者这样描述对ldap的感受：

> The short version of the story goes like this: I hate LDAP. I used to love it. But I loved it for all the wrong reasons. LDAP is supported as an authentication solution by many different pieces of software. Aside from its de jure standard status, its wide deployment cements it as a de facto standard as well.
>
> However, just because it is a standard doesn't mean it is a great idea.
>
> I'll admit that given its age LDAP has had a good run. I'm sure its authors carefully considered how to construct the protocol and chose ASN.1 and its encoding with all of wellest of well meaning intentions.
>
> The trouble is that with today's Internet, LDAP is just a pain in the ass. You can't call it from your browser. It's not human readable or easy to debug. Tooling is often arcane and confusing. It's way more complicated than what is needed for most simple authentication-only uses. (Yes, I know there are many other uses than authentication - but it's often too complicated for those too.)
>
> Likely owing to the complexity of the protocol, there seems to be virtually no easy to use library to implement the server side of the LDAP protocol that isn't tied in with some complete directory server system; and certainly not in a language as easy to "make it work" as Go.

他说他对ldap又爱又恨，因为ldap出现的最早，许多的三方软件都兼容支持它，它成了这方面的一个标准。但问题在于，它对于维护者而言，又是复杂麻烦的。就算是有Phpldapadmin这样的平台能够在浏览器维护，但看到那样上古的界面，以及复杂的交互逻辑，仍旧能够把不少人劝退。

鉴于此，我开发了这个现代化的openLDAP管理后台。

## 在线体验

> admin / 123456

演示地址：[http://demo-go-ldap-admin.eryajf.net](http://demo-go-ldap-admin.eryajf.net)

## 项目地址

| 分类 |                        GitHub                        |                        Gitee                        |
| :--: | :--------------------------------------------------: | :-------------------------------------------------: |
| 后端 |  https://github.com/eryajf-world/go-ldap-admin.git   |  https://gitee.com/eryajf-world/go-ldap-admin.git   |
| 前端 | https://github.com/eryajf-world/go-ldap-admin-ui.git | https://gitee.com/eryajf-world/go-ldap-admin-ui.git |

## 核心功能

- 基于 GIN WEB API 框架，基于Casbin的 RBAC 访问控制模型，JWT 认证，Validator 参数校验
- 基于 GORM 的数据库存储
- 基于 go-ldap 库的主逻辑交互
- 用户管理
  - 用户的增删改查
- 分组管理
  - 分组的增删改查
  - 分组内成员的管理

## 快速开始

你可以通过docker-compose在本地快速拉起进行体验。

快速拉起的容器包括：MySQL-5.7，openLDAP-1.4.0，phpldapadmin-0.9.0，go-ldap-admin。

服务端口映射如下：

|    Service    |         Port          |
| :-----------: | :-------------------: |
|     MySQL     |      `3307:3306`      |
|   openLDAP    |       `389:389`       |
| phpldapadmin  |       `8091:80`       |
| go-ldap-admin | `8090:80`,`8888:8888` |

拉起之前确认是否有与本地端口冲突的情况。

```
$ git clone https://github.com/eryajf-world/go-ldap-admin.git

$ cd docs/docker-compose

$ docker-compose up -d
```

当看到容器都正常运行之后，可以在本地访问：http://localhost:8090，用户名/密码：admin/123456

`登录页：`

![](http://t.eryajf.net/imgs/2022/05/17dbe07a137c9b4c.png)

`首页：`

![](http://t.eryajf.net/imgs/2022/05/b18c5fbf5ba0e6af.png)

`用户管理：`

![](http://t.eryajf.net/imgs/2022/05/f3ae695b703c00c8.png)

`分组管理：`

![](http://t.eryajf.net/imgs/2022/05/cb7bcd851b2c972f.png)

`分组内成员管理：`

![](http://t.eryajf.net/imgs/2022/05/f1732540ce0632de.png)

## 本地开发

### 前言准备

前提是已准备好MySQL与openLDAP，本地开发建议直接通过docker拉起即可，可参考文档：[https://wiki.eryajf.net/pages/3a0d5f](https://wiki.eryajf.net/pages/3a0d5f)。

### 拉取代码

```
# 后端代码
$ git clone https://github.com/eryajf-world/go-ldap-admin.git

# 前端代码
$ git clone https://github.com/eryajf-world/go-ldap-admin-ui.git
```

后端目录结构：

```
├─config     # viper读取配置
├─controller # controller层，响应路由请求的方法
├─docs       # 一些物料信息
├─logic      # 主要的处理逻辑
├─middleware # 中间件
├─model      # 结构体模型
├─public     # 一些公共的，工具类的放在这里
├─routes     # 所有路由
├─service    # 整合与底层存储交互的方法
├─svc        # 定义入参出参的结构体
└─test       # 跑测试用的
```

### 更改配置

```
# 修改后端配置
$ cd go-ldap-admin
# 文件路径 config.yml
$ vim config.yml

# 根据自己本地的情况，调整数据库以及openLDAP的配置信息。
```

### 启动服务

```
# 启动后端
$ cd go-ldap-admin
$ go mod tidy
$ go run main.go
$ make run

# 启动前端
$ cd go-ldap-admin-ui
$ yarn
$ yarn dev
```

本地访问：http://localhost:8090，用户名/密码：admin/密码是配置文件中openLDAP中admin的密码。

## 生产部署

生产环境单独部署，通过Nginx代理服务，配置如下：

```nginx
server {
    listen 80;
    server_name go-ldap-admin.eryajf.net;

    root /data/www/web/dist;

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
```

## 感谢

感谢如下优秀的项目，没有这些项目，不可能会有go-ldap-admin：

- 后端技术栈
  - [Gin-v1.6.3](https://github.com/gin-gonic/gin)
  - [Gorm-v1.20.12](https://github.com/go-gorm/gorm)
  - [Go-ldap-v3.4.2](https://github.com/go-ldap/ldap)
  - [Casbin-v2.22.0](https://github.com/casbin/casbin)
- 前端技术栈
  - [element-ui](https://github.com/ElemeFE/element)
  - [axios](https://github.com/axios/axios)

- 另外感谢
  - [go-web-mini](https://github.com/gnimli/go-web-mini)：项目基于该项目重构而成，感谢作者的付出。

## 另外

- 如果觉得项目不错，麻烦动动小手点个⭐️star⭐️!
- 如果你还有其他想法或者需求，欢迎在issue中交流！
- 程序还有很多bug，欢迎各位朋友一起协同共建！