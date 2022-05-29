package test

import (
	"fmt"
	"testing"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/public/tools"
)

func TestUnGenPassword(t *testing.T) {
	InitConfig()
	pass := "$2a$10$FlzrnJeE3Ad8uokvSAl/gunkRZsdREwlFZZqPcwfkekXOc9oAa9KS"
	fmt.Printf("秘钥为：%s\n", config.Conf.System.RSAPrivateBytes)
	// 密码通过RSA解密
	decodeData, err := tools.RSADecrypt([]byte(pass), config.Conf.System.RSAPrivateBytes)
	if err != nil {
		fmt.Printf("密码解密失败：%s\n", err)
	}
	fmt.Printf("密码解密后为：%s\n", string(decodeData))
}
