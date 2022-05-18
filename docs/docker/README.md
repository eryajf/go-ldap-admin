参考：https://wiki.eryajf.net/pages/95cf71/#%E5%90%AF%E7%94%A8-buildx-%E6%8F%92%E4%BB%B6

```sh
$ export DOCKER_CLI_EXPERIMENTAL=enabled

$ docker buildx create --use --name mybuilder

$ docker buildx inspect mybuilder --bootstrap

$ docker buildx build --no-cache -t eryajf/go-ldap-admin:v0.1 --platform=linux/arm64,linux/amd64 . --push
$ docker buildx build --no-cache -t eryajf/go-ldap-admin --platform=linux/arm64,linux/amd64 . --push
```