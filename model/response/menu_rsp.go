package response

import "github.com/eryajf/go-ldap-admin/model"

type MenuListRsp struct {
	Total int64        `json:"total"`
	Menus []model.Menu `json:"menus"`
}
