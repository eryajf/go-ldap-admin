package logic

import (
	"fmt"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/eryajf/go-ldap-admin/service/ildap"
	"github.com/eryajf/go-ldap-admin/service/isql"
)

func CommonAddGroup(group *model.Group) error {
	// 先在ldap中创建组
	err := ildap.Group.Add(group)
	if err != nil {
		return err
	}

	// 然后在数据库中创建组
	err = isql.Group.Add(group)
	if err != nil {
		return err
	}

	// 默认创建分组之后，需要将admin添加到分组中
	adminInfo := new(model.User)
	err = isql.User.Find(tools.H{"id": 1}, adminInfo)
	if err != nil {
		return err
	}

	err = isql.Group.AddUserToGroup(group, []model.User{*adminInfo})
	if err != nil {
		return err
	}

	return nil
}

func CommonUpdateGroup(oldGroup, newGroup *model.Group) error {
	//若配置了不允许修改分组名称，则不更新分组名称
	if !config.Conf.Ldap.GroupNameModify {
		newGroup.GroupName = oldGroup.GroupName
	}

	err := ildap.Group.Update(oldGroup, newGroup)
	if err != nil {
		return err
	}
	err = isql.Group.Update(newGroup)
	if err != nil {
		return err
	}
	return nil
}

func CommonAddUser(user *model.User, groupId []uint) error {
	if user.Departments == "" {
		user.Departments = "默认:研发中心"
	}
	if user.GivenName == "" {
		user.GivenName = user.Nickname
	}
	if user.PostalAddress == "" {
		user.PostalAddress = "默认:地球"
	}
	if user.Position == "" {
		user.Position = "默认:技术"
	}
	if user.Introduction == "" {
		user.Introduction = user.Nickname
	}
	// 先将用户添加到MySQL
	err := isql.User.Add(user)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("向MySQL创建用户失败：" + err.Error()))
	}
	// 再将用户添加到ldap
	err = ildap.User.Add(user)
	if err != nil {
		return tools.NewLdapError(fmt.Errorf("AddUser向LDAP创建用户失败：" + err.Error()))
	}
	// 获取用户将要添加的分组
	groups, err := isql.Group.GetGroupByIds(groupId)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("根据部门ID获取部门信息失败" + err.Error()))
	}
	for _, group := range groups {
		if group.GroupDN[:3] == "ou=" {
			continue
		}
		// 先将用户和部门信息维护到MySQL
		err := isql.Group.AddUserToGroup(group, []model.User{*user})
		if err != nil {
			return tools.NewMySqlError(fmt.Errorf("向MySQL添加用户到分组关系失败：" + err.Error()))
		}
		//根据选择的部门，添加到部门内
		err = ildap.Group.AddUserToGroup(group.GroupDN, user.UserDN)
		if err != nil {
			return tools.NewMySqlError(fmt.Errorf("向Ldap添加用户到分组关系失败：" + err.Error()))
		}
	}
	return nil
}

func CommonUpdateUser(oldUser, newUser *model.User, groupId []uint) error {
	// 更新用户
	if !config.Conf.Ldap.UserNameModify {
		newUser.Username = oldUser.Username
	}

	err := ildap.User.Update(oldUser.Username, newUser)
	if err != nil {
		return tools.NewLdapError(fmt.Errorf("在LDAP更新用户失败：" + err.Error()))
	}

	err = isql.User.Update(newUser)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("在MySQL更新用户失败：" + err.Error()))
	}

	//判断部门信息是否有变化有变化则更新相应的数据库
	oldDeptIds := tools.StringToSlice(oldUser.DepartmentId, ",")
	addDeptIds, removeDeptIds := tools.ArrUintCmp(oldDeptIds, groupId)

	// 先处理添加的部门
	addgroups, err := isql.Group.GetGroupByIds(addDeptIds)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("根据部门ID获取部门信息失败" + err.Error()))
	}
	for _, group := range addgroups {
		if group.GroupDN[:3] == "ou=" {
			continue
		}
		// 先将用户和部门信息维护到MySQL
		err := isql.Group.AddUserToGroup(group, []model.User{*newUser})
		if err != nil {
			return tools.NewMySqlError(fmt.Errorf("向MySQL添加用户到分组关系失败：" + err.Error()))
		}
		//根据选择的部门，添加到部门内
		err = ildap.Group.AddUserToGroup(group.GroupDN, newUser.UserDN)
		if err != nil {
			return tools.NewLdapError(fmt.Errorf("向Ldap添加用户到分组关系失败：" + err.Error()))
		}
	}

	// 再处理删除的部门
	removegroups, err := isql.Group.GetGroupByIds(removeDeptIds)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("根据部门ID获取部门信息失败" + err.Error()))
	}
	for _, group := range removegroups {
		if group.GroupDN[:3] == "ou=" {
			continue
		}
		err := isql.Group.RemoveUserFromGroup(group, []model.User{*newUser})
		if err != nil {
			return tools.NewMySqlError(fmt.Errorf("在MySQL将用户从分组移除失败：" + err.Error()))
		}
		err = ildap.Group.RemoveUserFromGroup(group.GroupDN, newUser.UserDN)
		if err != nil {
			return tools.NewMySqlError(fmt.Errorf("在ldap将用户从分组移除失败：" + err.Error()))
		}
	}
	return nil
}
