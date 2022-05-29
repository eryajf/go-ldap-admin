package response

import "github.com/eryajf/go-ldap-admin/model"

type GroupListRsp struct {
	Total  int64         `json:"total"`
	Groups []model.Group `json:"groups"`
}

type Guser struct {
	UserId       int64  `json:"userId"`
	UserName     string `json:"userName"`
	NickName     string `json:"nickName"`
	Mail         string `json:"mail"`
	JobNumber    string `json:"jobNumber"`
	Mobile       string `json:"mobile"`
	Introduction string `json:"introduction"`
}

type GroupUsers struct {
	GroupId     int64   `json:"groupId"`
	GroupName   string  `json:"groupName"`
	GroupRemark string  `json:"groupRemark"`
	UserList    []Guser `json:"userList"`
}
