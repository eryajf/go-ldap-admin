package dingtalk

import (
	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/public/common"
	"github.com/zhaoyunxing92/dingtalk/v2"
)

func InitDingTalkClient() *dingtalk.DingTalk {
	dingTalk, err := dingtalk.NewClient(config.Conf.DingTalk.DingTalkAppKey, config.Conf.DingTalk.DingTalkAppSecret)
	if err != nil {
		common.Log.Error("init dingding client failed, err:%v\n", err)
	}
	return dingTalk
}

// 部门结构体
type DingTalkDept struct {
	Id       int    `json:"dept_id"`   // 部门ID
	Name     string `json:"name"`      // 部门名称拼音
	Remark   string `json:"remark"`    // 部门中文名
	ParentId int    `json:"parent_id"` // 父部门ID
}

// DingTalkUser 部门用户详情
type DingTalkUser struct {
	UserId               string `json:"userid"`
	UnionId              string `json:"unionid"`                // 员工在当前开发者企业账号范围内的唯一标识
	Name                 string `json:"name"`                   // 员工名称
	Avatar               string `json:"avatar"`                 // 头像
	StateCode            string `json:"state_code"`             // 国际电话区号
	ManagerUserId        string `json:"manager_userid"`         // 员工的直属主管
	Mobile               string `json:"mobile"`                 // 手机号码
	HideMobile           bool   `json:"hide_mobile"`            // 是否号码隐藏
	Telephone            string `json:"telephone"`              // 分机号
	JobNumber            string `json:"job_number"`             // 员工工号
	Title                string `json:"title"`                  // 职位
	WorkPlace            string `json:"work_place"`             // 办公地点
	Remark               string `json:"remark"`                 // 备注
	LoginId              string `json:"loginId"`                // 专属帐号登录名
	DeptIds              []int  `json:"dept_id_list"`           // 所属部门ID列表
	DeptOrder            int    `json:"dept_order"`             // 员工在部门中的排序
	Extension            string `json:"extension"`              // 员工在对应的部门中的排序
	HiredDate            int    `json:"hired_date"`             // 入职时间
	Active               bool   `json:"active"`                 // 是否激活了钉钉
	Admin                bool   `json:"admin"`                  //是否为企业的管理员：
	Boss                 bool   `json:"boss"`                   // 是否为企业的老板
	ExclusiveAccount     bool   `json:"exclusive_account"`      // 是否专属帐号
	Leader               bool   `json:"leader"`                 // 是否是部门的主管
	ExclusiveAccountType string `json:"exclusive_account_type"` //专属帐号类型：sso：企业自建专属帐号 dingtalk：钉钉自建专属帐号
	OrgEmail             string `json:"org_email"`              //员工的企业邮箱,如果员工的企业邮箱没有开通，返回信息中不包含该数据
	Email                string `json:"email"`                  //员工邮箱,企业内部应用如果没有返回该字段，需要检查当前应用通讯录权限中邮箱等个人信息权限是否开启,员工信息面板中有邮箱字段值才返回该字段,第三方企业应用不返回该参数
}
