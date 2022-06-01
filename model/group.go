package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupName          string   `gorm:"type:varchar(20);comment:'分组名称'" json:"groupName"`
	Remark             string   `gorm:"type:varchar(100);comment:'分组中文说明'" json:"remark"`
	Creator            string   `gorm:"type:varchar(20);comment:'创建人'" json:"creator"`
	GroupType          string   `gorm:"type:varchar(20);comment:'分组类型：cn、ou'" json:"groupType"`
	Users              []*User  `gorm:"many2many:group_users" json:"users"`
	ParentId           uint     `gorm:"default:0;comment:'父组编号(编号为0时表示根组)'" json:"parentId"`
	SourceDeptId       string   `gorm:"type:varchar(100);comment:'部门编号'" json:"sourceDeptId"`
	Source             string   `gorm:"type:varchar(20);comment:'来源：dingTalk、weCom、ldap、platform'" json:"source"`
	SourceDeptParentId string   `gorm:"type:varchar(100);comment:'父部门编号'" json:"sourceDeptParentId"`
	SourceUserNum      int      `gorm:"default:0;comment:'部门下的用户数量，从第三方获取的数据'" json:"source_user_num"`
	Children           []*Group `gorm:"-" json:"children"`
}
