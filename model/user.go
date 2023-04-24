package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string  `gorm:"type:varchar(50);not null;unique;comment:'用户名'" json:"username"`                    // 用户名
	Password      string  `gorm:"size:255;not null;comment:'用户密码'" json:"password"`                                  // 用户密码
	Nickname      string  `gorm:"type:varchar(50);comment:'中文名'" json:"nickname"`                                    // 昵称
	GivenName     string  `gorm:"type:varchar(50);comment:'花名'" json:"givenName"`                                    // 花名，如果有的话，没有的话用昵称占位
	Mail          string  `gorm:"type:varchar(100);comment:'邮箱'" json:"mail"`                                        // 邮箱
	JobNumber     string  `gorm:"type:varchar(20);comment:'工号'" json:"jobNumber"`                                    // 工号
	Mobile        string  `gorm:"type:varchar(15);not null;unique;comment:'手机号'" json:"mobile"`                      // 手机号
	Avatar        string  `gorm:"type:varchar(255);comment:'头像'" json:"avatar"`                                      // 头像
	PostalAddress string  `gorm:"type:varchar(255);comment:'地址'" json:"postalAddress"`                               // 地址
	Departments   string  `gorm:"type:varchar(512);comment:'部门'" json:"departments"`                                 // 部门
	Position      string  `gorm:"type:varchar(128);comment:'职位'" json:"position"`                                    //  职位
	Introduction  string  `gorm:"type:varchar(255);comment:'个人简介'" json:"introduction"`                              // 个人简介
	Status        uint    `gorm:"type:tinyint(1);default:1;comment:'状态:1在职, 2离职'" json:"status"`                     // 状态
	Creator       string  `gorm:"type:varchar(20);;comment:'创建者'" json:"creator"`                                    // 创建者
	Source        string  `gorm:"type:varchar(50);comment:'用户来源：dingTalk、wecom、feishu、ldap、platform'" json:"source"` // 来源
	DepartmentId  string  `gorm:"type:varchar(100);not null;comment:'部门id'" json:"departmentId"`                     // 部门id
	Roles         []*Role `gorm:"many2many:user_roles" json:"roles"`                                                 // 角色
	SourceUserId  string  `gorm:"type:varchar(100);not null;comment:'第三方用户id'" json:"sourceUserId"`                  // 第三方用户id
	SourceUnionId string  `gorm:"type:varchar(100);not null;comment:'第三方唯一unionId'" json:"sourceUnionId"`            // 第三方唯一unionId
	UserDN        string  `gorm:"type:varchar(255);not null;comment:'用户dn'" json:"userDn"`                           // 用户在ldap的dn
	SyncState     uint    `gorm:"type:tinyint(1);default:1;comment:'同步状态:1已同步, 2未同步'" json:"syncState"`              // 数据到ldap的同步状态
}

func (u *User) SetUserName(userName string) {
	u.Username = userName
}

func (u *User) SetNickName(nickName string) {
	u.Nickname = nickName
}

func (u *User) SetGivenName(givenName string) {
	u.GivenName = givenName
}

func (u *User) SetMail(mail string) {
	u.Mail = mail
}

func (u *User) SetJobNumber(jobNum string) {
	u.JobNumber = jobNum
}

func (u *User) SetMobile(mobile string) {
	u.Mobile = mobile
}

func (u *User) SetAvatar(avatar string) {
	u.Avatar = avatar
}

func (u *User) SetPostalAddress(address string) {
	u.PostalAddress = address
}

func (u *User) SetPosition(position string) {
	u.Position = position
}

func (u *User) SetIntroduction(desc string) {
	u.Introduction = desc
}

func (u *User) SetSourceUserId(sourceUserId string) {
	u.SourceUserId = sourceUserId
}

func (u *User) SetSourceUnionId(sourceUnionId string) {
	u.SourceUnionId = sourceUnionId
}
