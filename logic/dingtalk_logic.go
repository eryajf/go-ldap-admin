package logic

import (
	"fmt"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/public/client/dingtalk"
	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/eryajf/go-ldap-admin/service/ildap"
	"github.com/eryajf/go-ldap-admin/service/isql"
	"github.com/gin-gonic/gin"
)

type DingTalkLogic struct {
}

//通过钉钉获取部门信息
func (d *DingTalkLogic) SyncDingTalkDepts(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	// 1.获取所有部门
	deptSource, err := dingtalk.GetAllDepts()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("获取钉钉部门列表失败：%s", err.Error()))
	}
	depts, err := ConvertDeptData(config.Conf.DingTalk.Flag, deptSource)
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("转换钉钉部门数据失败：%s", err.Error()))
	}
	// 2.将部门这个数组进行拆分，一组是父ID为1的，一组是父ID不为1的
	var firstDepts []*model.Group
	var otherDepts []*model.Group
	for _, dept := range depts {
		if dept.SourceDeptParentId == fmt.Sprintf("%s_1", config.Conf.DingTalk.Flag) {
			firstDepts = append(firstDepts, dept)
		} else {
			otherDepts = append(otherDepts, dept)
		}
	}

	// 3.先写父ID为1的，再写父ID不为1的，因为数据库中需要获取父部门的ID，从远程过来的数据并没有这个ID，而父ID为1的则是固定根部门下的一级部门
	for _, dept := range firstDepts {
		err := d.AddDepts(dept)
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("DsyncDingTalkDepts添加根部门失败：%s", err.Error()))
		}
	}
	for _, dept := range otherDepts {
		err := d.AddDepts(dept)
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("DsyncDingTalkDepts添加其他部门失败：%s", err.Error()))
		}
	}

	return nil, nil
}

// AddGroup 添加部门数据
func (d DingTalkLogic) AddDepts(group *model.Group) error {
	// 判断部门名称是否存在
	parentGroup := new(model.Group)
	err := isql.Group.Find(tools.H{"source_dept_id": group.SourceDeptParentId}, parentGroup) // 查询当前分组父ID在MySQL中的数据信息
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("查询父级部门失败：%s", err.Error()))
	}
	if !isql.Group.Exist(tools.H{"source_dept_id": group.SourceDeptId}) { // 判断当前部门是否已落库
		// 此时的 group 已经附带了Build后动态关联好的字段，接下来将一些确定性的其他字段值添加上，就可以创建这个分组了
		group.Creator = "system"
		group.GroupType = "cn"
		group.ParentId = parentGroup.ID
		group.Source = config.Conf.DingTalk.Flag
		group.GroupDN = fmt.Sprintf("cn=%s,%s", group.GroupName, parentGroup.GroupDN)

		err = CommonAddGroup(group)
		if err != nil {
			return tools.NewOperationError(fmt.Errorf("添加部门失败：%s", err.Error()))
		}
	}
	return nil
}

//根据现有数据库同步到的部门信息，开启用户同步
func (d DingTalkLogic) SyncDingTalkUsers(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	// 1.获取钉钉用户列表
	staffSource, err := dingtalk.GetAllUsers()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("SyncDingTalkUsers获取钉钉用户列表失败：%s", err.Error()))
	}
	staffs, err := ConvertUserData(config.Conf.DingTalk.Flag, staffSource)
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("转换钉钉用户数据失败：%s", err.Error()))
	}
	// 2.遍历用户，开始写入
	for _, staff := range staffs {
		// 入库
		err = d.AddUsers(staff)
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncDingTalkUsers写入用户失败：%s", err.Error()))
		}
	}

	// 3.获取钉钉已离职用户id列表
	userIds, err := dingtalk.GetLeaveUserIds()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("SyncDingTalkUsers获取钉钉离职用户列表失败：%s", err.Error()))
	}
	// 4.遍历id，开始处理
	for _, uid := range userIds {
		if isql.User.Exist(tools.H{"source_user_id": fmt.Sprintf("%s_%s", config.Conf.DingTalk.Flag, uid)}) {
			user := new(model.User)
			err = isql.User.Find(tools.H{"source_user_id": fmt.Sprintf("%s_%s", config.Conf.DingTalk.Flag, uid)}, user)
			if err != nil {
				return nil, tools.NewMySqlError(fmt.Errorf("在MySQL查询用户失败: " + err.Error()))
			}
			// 先从ldap删除用户
			err = ildap.User.Delete(user.UserDN)
			if err != nil {
				return nil, tools.NewLdapError(fmt.Errorf("在LDAP删除用户失败" + err.Error()))
			}
			// 然后更新MySQL中用户状态
			err = isql.User.ChangeStatus(int(user.ID), 2)
			if err != nil {
				return nil, tools.NewMySqlError(fmt.Errorf("在MySQL更新用户状态失败: " + err.Error()))
			}
		}
	}

	return nil, nil
}

// AddUser 添加用户数据
func (d DingTalkLogic) AddUsers(user *model.User) error {
	// 根据 unionid 查询用户,不存在则创建
	if !isql.User.Exist(tools.H{"source_union_id": user.SourceUnionId}) {
		// 根据角色id获取角色
		roles, err := isql.Role.GetRolesByIds([]uint{2}) // 默认添加为普通用户角色
		if err != nil {
			return tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败:%s", err.Error()))
		}
		user.Creator = "system"
		user.Roles = roles
		user.Password = config.Conf.Ldap.UserInitPassword
		user.Source = config.Conf.DingTalk.Flag
		user.UserDN = fmt.Sprintf("uid=%s,%s", user.Username, config.Conf.Ldap.UserDN)
		err = CommonAddUser(user, tools.StringToSlice(user.DepartmentId, ","))
		if err != nil {
			return err
		}
	}
	return nil
}
