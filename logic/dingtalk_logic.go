package logic

import (
	"fmt"
	"strings"

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

	// 2.将远程数据转换成树
	deptTree := GroupListToTree(fmt.Sprintf("%s_1", config.Conf.DingTalk.Flag), depts)

	// 3.根据树进行创建
	err = d.addDepts(deptTree.Children)

	return nil, err
}

// 添加部门
func (d DingTalkLogic) addDepts(depts []*model.Group) error {
	for _, dept := range depts {
		err := d.AddDepts(dept)
		if err != nil {
			return tools.NewOperationError(fmt.Errorf("DsyncDingTalkDepts添加部门失败: %s", err.Error()))
		}
		if len(dept.Children) != 0 {
			err = d.addDepts(dept.Children)
			if err != nil {
				return tools.NewOperationError(fmt.Errorf("DsyncDingTalkDepts添加部门失败: %s", err.Error()))
			}
		}
	}
	return nil
}

// AddGroup 添加部门数据
func (d DingTalkLogic) AddDepts(group *model.Group) error {
	parentGroup := new(model.Group)
	err := isql.Group.Find(tools.H{"source_dept_id": group.SourceDeptParentId}, parentGroup) // 查询当前分组父ID在MySQL中的数据信息
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("查询父级部门失败：%s", err.Error()))
	}

	// 此时的 group 已经附带了Build后动态关联好的字段，接下来将一些确定性的其他字段值添加上，就可以创建这个分组了
	group.Creator = "system"
	group.GroupType = "cn"
	group.ParentId = parentGroup.ID
	group.Source = config.Conf.DingTalk.Flag
	group.GroupDN = fmt.Sprintf("cn=%s,%s", group.GroupName, parentGroup.GroupDN)

	if !isql.Group.Exist(tools.H{"group_dn": group.GroupDN}) { // 判断当前部门是否已落库
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
	// 根据角色id获取角色
	roles, err := isql.Role.GetRolesByIds([]uint{2}) // 默认添加为普通用户角色
	if err != nil {
		return tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败:%s", err.Error()))
	}
	user.Roles = roles
	user.Creator = "system"
	user.Source = config.Conf.DingTalk.Flag
	user.Password = config.Conf.Ldap.UserInitPassword
	user.UserDN = fmt.Sprintf("uid=%s,%s", user.Username, config.Conf.Ldap.UserDN)

	// 根据 user_dn 查询用户,不存在则创建
	if !isql.User.Exist(tools.H{"user_dn": user.UserDN}) {
		// 获取用户将要添加的分组
		groups, err := isql.Group.GetGroupByIds(tools.StringToSlice(user.DepartmentId, ","))
		if err != nil {
			return tools.NewMySqlError(fmt.Errorf("根据部门ID获取部门信息失败" + err.Error()))
		}
		var deptTmp string
		for _, group := range groups {
			deptTmp = deptTmp + group.GroupName + ","
		}
		user.Departments = strings.TrimRight(deptTmp, ",")

		// 新增用户
		err = CommonAddUser(user, groups)
		if err != nil {
			return err
		}
	}
	return nil
}
