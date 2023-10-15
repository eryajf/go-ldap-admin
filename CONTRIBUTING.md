# 贡献者指南

欢迎反馈、bug 报告和拉取请求，可点击[issue](https://github.com/eryajf/go-ldap-admin/issues) 提交.

如果你是第一次进行 GitHub 协作，可参阅： [协同开发流程](https://howtosos.eryajf.net/HowToStartOpenSource/01-basic-content/03-collaborative-development-process.html)

1. 项目使用`golangci-lint`进行检测，提交 pr 之前请在本地执行 `make lint` 并通过。

2. 如非必要，尽可能谨慎新增配置文件，以免造成升级时产生意料之外的问题。

3. 注意一些功能调整，如何涉及到前端页面调整，尽可能两个 pr 有所关联，否则不好合并。
