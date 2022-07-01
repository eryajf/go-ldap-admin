package dingtalk

import (
	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/public/common"
	"github.com/zhaoyunxing92/dingtalk/v2"
)

func InitDingTalkClient() *dingtalk.DingTalk {
	dingTalk, err := dingtalk.NewClient(config.Conf.DingTalk.AppKey, config.Conf.DingTalk.AppSecret)
	if err != nil {
		common.Log.Error("init dingding client failed, err:%v\n", err)
	}
	return dingTalk
}
