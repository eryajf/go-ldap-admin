package common

import (
	"fmt"
	"net"
	"time"

	"github.com/eryajf-world/go-ldap-admin/config"

	ldap "github.com/go-ldap/ldap/v3"
)

// 全局mysql数据库变量
var LDAP *ldap.Conn

// Init 初始化连接
func InitLDAP() {
	// Dail有两个参数 network,  address, 返回 (*Conn,  error)
	ldap, err := ldap.DialURL(config.Conf.Ldap.LdapUrl, ldap.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}))
	if err != nil {
		Log.Panicf("初始化ldap连接异常: %v", err)
		panic(fmt.Errorf("初始化ldap连接异常: %v", err))
	}
	err = ldap.Bind(config.Conf.Ldap.LdapAdminDN, config.Conf.Ldap.LdapAdminPass)
	if err != nil {
		Log.Panicf("绑定admin账号异常: %v", err)
		panic(fmt.Errorf("绑定admin账号异常: %v", err))
	}

	// 全局LDAP赋值
	LDAP = ldap

	// 隐藏密码
	showDsn := fmt.Sprintf(
		"%s:******@tcp(%s)",
		config.Conf.Ldap.LdapAdminDN,
		config.Conf.Ldap.LdapUrl,
	)

	Log.Info("初始化ldap完成! dsn: ", showDsn)
}
