package ildap

import (
	"errors"
	"fmt"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/public/common"

	ldap "github.com/go-ldap/ldap/v3"
)

type GroupService struct{}

// Add 添加资源
func (x GroupService) Add(g *model.Group, pdn string) error { //organizationalUnit
	parentDn := config.Conf.Ldap.LdapBaseDN
	if pdn != "" {
		parentDn = fmt.Sprintf("%s,%s", pdn, config.Conf.Ldap.LdapBaseDN)
	}
	dn := fmt.Sprintf("%s=%s,%s", g.GroupType, g.GroupName, parentDn)
	add := ldap.NewAddRequest(dn, nil)
	if g.GroupType == "ou" {
		add.Attribute("objectClass", []string{"organizationalUnit", "top"}) // 如果定义了 groupOfNAmes，那么必须指定member，否则报错如下：object class 'groupOfNames' requires attribute 'member'
	}
	if g.GroupType == "cn" {
		add.Attribute("objectClass", []string{"groupOfUniqueNames", "top"})
		add.Attribute("uniqueMember", []string{config.Conf.Ldap.LdapAdminDN}) // 所以这里创建组的时候，默认将admin加入其中，以免创建时没有人员而报上边的错误
	}
	add.Attribute(g.GroupType, []string{g.GroupName})
	add.Attribute("description", []string{g.Remark})

	return common.LDAP.Add(add)
}

// UpdateGroup 更新一个分组
func (x GroupService) Update(g *model.Group, pdn string, oldGroupName, oldRemark string) error {
	parentDn := "," + config.Conf.Ldap.LdapBaseDN
	if pdn != "" {
		parentDn = fmt.Sprintf("%s,%s", pdn, config.Conf.Ldap.LdapBaseDN)
	}
	//默认更新remark字段
	if g.Remark != oldRemark {
		modify := ldap.NewModifyRequest(parentDn, nil)
		modify.Replace("description", []string{g.Remark})
		err := common.LDAP.Modify(modify)
		if err != nil {
			return err
		}
	}
	// 如果配置文件允许修改分组名称，且分组名称发生了变化，那么执行修改分组名称
	if config.Conf.Ldap.LdapGroupNameModify && g.GroupName != oldGroupName {
		rdn := fmt.Sprintf("%s=%s", g.GroupType, g.GroupName)
		modify := ldap.NewModifyDNRequest(parentDn, rdn, true, "")
		err := common.LDAP.ModifyDN(modify)
		if err != nil {
			return err
		}
	}
	return nil
}

// Delete 删除资源
func (x GroupService) Delete(pdn string) error {
	del := ldap.NewDelRequest(pdn, nil)
	return common.LDAP.Del(del)
}

// AddUserToGroup 添加用户到分组
func (x GroupService) AddUserToGroup(dn, udn string) error {
	//判断dn是否以ou开头
	if dn[:3] == "ou=" {
		return errors.New("不能添加用户到OU组织单元")
	}
	newmr := ldap.NewModifyRequest(dn, nil)
	newmr.Add("uniqueMember", []string{udn})
	return common.LDAP.Modify(newmr)
}

// DelUserFromGroup 将用户从分组删除
func (x GroupService) RemoveUserFromGroup(gdn, user string) error {
	udn := fmt.Sprintf("uid=%s,%s", user, config.Conf.Ldap.LdapUserDN)
	newmr := ldap.NewModifyRequest(gdn, nil)
	newmr.Delete("uniqueMember", []string{udn})
	return common.LDAP.Modify(newmr)
}
