package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupName string  `gorm:"type:varchar(20);comment:'分组名称'" json:"groupName"`
	Remark    string  `gorm:"type:varchar(100);comment:'分组中文说明'" json:"remark"`
	Creator   string  `gorm:"type:varchar(20);comment:'创建人'" json:"creator"`
	Users     []*User `gorm:"many2many:group_users" json:"users"`
}
