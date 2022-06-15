package wechat

import (
	"github.com/eryajf/go-ldap-admin/config"
	"github.com/wenerme/go-wecom/wecom"
)

func InitWeComClient() *wecom.Client {
	client := wecom.NewClient(wecom.Conf{
		CorpID:     config.Conf.WeCom.CorpID,
		AgentID:    config.Conf.WeCom.AgentID,
		CorpSecret: config.Conf.WeCom.CorpSecret,
	})
	return client
}
