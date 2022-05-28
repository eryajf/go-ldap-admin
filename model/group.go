package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupName string   `gorm:"type:varchar(20);comment:'分组名称'" json:"groupName"`
	Remark    string   `gorm:"type:varchar(100);comment:'分组中文说明'" json:"remark"`
	Creator   string   `gorm:"type:varchar(20);comment:'创建人'" json:"creator"`
	GroupType string   `gorm:"type:varchar(20);comment:'分组类型：cn、ou'" json:"groupType"`
	Users     []*User  `gorm:"many2many:group_users" json:"users"`
	ParentId  uint     `gorm:"default:0;comment:'父组编号(编号为0时表示根组)'" json:"parentId"`
	Children  []*Group `gorm:"-" json:"children"`
}
