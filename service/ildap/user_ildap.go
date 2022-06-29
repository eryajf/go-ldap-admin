package ildap

import (
	"fmt"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/public/common"
	"github.com/eryajf/go-ldap-admin/public/tools"

	ldap "github.com/go-ldap/ldap/v3"
)

type UserService struct{}

// 创建资源
func (x UserService) Add(user *model.User) error {
	add := ldap.NewAddRequest(user.UserDN, nil)
	add.Attribute("objectClass", []string{"inetOrgPerson"})
	add.Attribute("cn", []string{user.Username})
	add.Attribute("sn", []string{user.Nickname})
	add.Attribute("businessCategory", []string{user.Departments})
	add.Attribute("departmentNumber", []string{user.Position})
	add.Attribute("description", []string{user.Introduction})
	add.Attribute("displayName", []string{user.Nickname})
	add.Attribute("mail", []string{user.Mail})
	add.Attribute("employeeNumber", []string{user.JobNumber})
	add.Attribute("givenName", []string{user.GivenName})
	add.Attribute("postalAddress", []string{user.PostalAddress})
	add.Attribute("mobile", []string{user.Mobile})
	add.Attribute("uid", []string{user.Username})
	add.Attribute("userPassword", []string{tools.NewParPasswd(user.Password)})
	return common.LDAP.Add(add)
}

// Update 更新资源
func (x UserService) Update(oldusername string, user *model.User) error {
	modify := ldap.NewModifyRequest(user.UserDN, nil)
	modify.Replace("cn", []string{user.Nickname})
	modify.Replace("sn", []string{oldusername})
	modify.Replace("businessCategory", []string{user.Departments})
	modify.Replace("departmentNumber", []string{user.Position})
	modify.Replace("description", []string{user.Introduction})
	modify.Replace("displayName", []string{user.Nickname})
	modify.Replace("mail", []string{user.Mail})
	modify.Replace("employeeNumber", []string{user.JobNumber})
	modify.Replace("givenName", []string{user.GivenName})
	modify.Replace("postalAddress", []string{user.PostalAddress})
	modify.Replace("mobile", []string{user.Mobile})
	err := common.LDAP.Modify(modify)
	if err != nil {
		return err
	}
	if config.Conf.Ldap.UserNameModify && oldusername != user.Username {
		modifyDn := ldap.NewModifyDNRequest(fmt.Sprintf("uid=%s,%s", oldusername, config.Conf.Ldap.UserDN), fmt.Sprintf("uid=%s", user.Username), true, "")
		return common.LDAP.ModifyDN(modifyDn)
	}
	return nil
}

// Delete 删除资源
func (x UserService) Delete(udn string) error {
	del := ldap.NewDelRequest(udn, nil)
	return common.LDAP.Del(del)
}

// ChangePwd 修改用户密码，此处旧密码也可以为空，ldap可以直接通过用户DN加上新密码来进行修改
func (x UserService) ChangePwd(udn, oldpasswd, newpasswd string) error {
	modifyPass := ldap.NewPasswordModifyRequest(udn, oldpasswd, newpasswd)
	_, err := common.LDAP.PasswordModify(modifyPass)
	if err != nil {
		return fmt.Errorf("password modify failed for %s, err: %v", udn, err)
	}
	return nil
}

// NewPwd 新旧密码都是空，通过管理员可以修改成功并返回新的密码
func (x UserService) NewPwd(username string) (string, error) {
	udn := fmt.Sprintf("uid=%s,%s", username, config.Conf.Ldap.UserDN)
	if username == "admin" {
		udn = config.Conf.Ldap.AdminDN
	}
	modifyPass := ldap.NewPasswordModifyRequest(udn, "", "")
	newpass, err := common.LDAP.PasswordModify(modifyPass)
	if err != nil {
		return "", fmt.Errorf("password modify failed for %s, err: %v", username, err)
	}
	return newpass.GeneratedPassword, nil
}
