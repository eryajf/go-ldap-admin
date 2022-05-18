package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string  `gorm:"type:varchar(10);not null;unique" json:"username"`           // 用户名
	Password      string  `gorm:"size:255;not null" json:"password"`                          // 用户密码
	Nickname      string  `gorm:"type:varchar(10)" json:"nickname"`                           // 昵称
	GivenName     string  `gorm:"type:varchar(10)" json:"givenName"`                          // 花名，如果有的话，没有的话用昵称占位
	Mail          string  `gorm:"type:varchar(20)" json:"mail"`                               // 邮箱
	JobNumber     string  `gorm:"type:varchar(5)" json:"jobNumber"`                           // 工号
	Mobile        string  `gorm:"type:varchar(11);not null;unique" json:"mobile"`             // 手机号
	Avatar        string  `gorm:"type:varchar(255)" json:"avatar"`                            // 头像
	PostalAddress string  `gorm:"type:varchar(255)" json:"postalAddress"`                     // 地址
	Departments   string  `gorm:"type:varchar(128)" json:"departments"`                       // 部门
	Position      string  `gorm:"type:varchar(128)" json:"position"`                          //  职位
	Introduction  string  `gorm:"type:varchar(255)" json:"introduction"`                      // 个人简介
	Status        uint    `gorm:"type:tinyint(1);default:1;comment:'1在职, 2离职'" json:"status"` // 状态
	Creator       string  `gorm:"type:varchar(20);" json:"creator"`                           // 创建者
	Roles         []*Role `gorm:"many2many:user_roles" json:"roles"`                          // 角色
}
