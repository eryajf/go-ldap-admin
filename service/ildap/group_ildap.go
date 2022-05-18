package ildap

import (
	"fmt"

	"github.com/eryajf-world/go-ldap-admin/config"
	"github.com/eryajf-world/go-ldap-admin/model"
	"github.com/eryajf-world/go-ldap-admin/public/common"

	ldap "github.com/go-ldap/ldap/v3"
)

type GroupService struct{}

// Add 添加资源
func (x GroupService) Add(g *model.Group) error {
	add := ldap.NewAddRequest(fmt.Sprintf("cn=%s,%s", g.GroupName, config.Conf.Ldap.LdapGroupDN), nil)
	add.Attribute("objectClass", []string{"groupOfNames", "top"}) // 如果定义了 groupOfNAmes，那么必须指定member，否则报错如下：object class 'groupOfNames' requires attribute 'member'
	add.Attribute("cn", []string{g.GroupName})
	add.Attribute("description", []string{g.Remark})
	add.Attribute("member", []string{config.Conf.Ldap.LdapAdminDN}) // 所以这里创建组的时候，默认将admin加入其中，以免创建时没有人员而报上边的错误

	return common.LDAP.Add(add)
}

// UpdateGroup 更新一个分组
func (x GroupService) Update(g *model.Group) error {
	modify := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,%s", g.GroupName, config.Conf.Ldap.LdapGroupDN), nil)
	modify.Replace("description", []string{g.Remark})
	return common.LDAP.Modify(modify)
}

// Delete 删除资源
func (x GroupService) Delete(group string) error {
	del := ldap.NewDelRequest(fmt.Sprintf("cn=%s,%s", group, config.Conf.Ldap.LdapGroupDN), nil)
	return common.LDAP.Del(del)
}

// AddUserToGroup 添加用户到分组
func (x GroupService) AddUserToGroup(group, user string) error {
	udn := fmt.Sprintf("uid=%s,%s", user, config.Conf.Ldap.LdapUserDN)
	if user == "admin" {
		udn = config.Conf.Ldap.LdapAdminDN
	}
	gdn := fmt.Sprintf("cn=%s,%s", group, config.Conf.Ldap.LdapGroupDN)
	newmr := ldap.NewModifyRequest(gdn, nil)
	newmr.Add("member", []string{udn})
	return common.LDAP.Modify(newmr)
}

// DelUserFromGroup 将用户从分组删除
func (x GroupService) RemoveUserFromGroup(group, user string) error {
	udn := fmt.Sprintf("uid=%s,%s", user, config.Conf.Ldap.LdapUserDN)
	gdn := fmt.Sprintf("cn=%s,%s", group, config.Conf.Ldap.LdapGroupDN)
	newmr := ldap.NewModifyRequest(gdn, nil)
	newmr.Delete("member", []string{udn})
	return common.LDAP.Modify(newmr)
}
