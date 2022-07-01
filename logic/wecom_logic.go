package logic

import (
	"fmt"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/model/request"
	"github.com/eryajf/go-ldap-admin/public/client/wechat"

	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/eryajf/go-ldap-admin/service/ildap"
	"github.com/eryajf/go-ldap-admin/service/isql"
	"github.com/gin-gonic/gin"
)

type WeComLogic struct {
}

//通过企业微信获取部门信息
func (d *WeComLogic) SyncWeComDepts(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	// 1.获取所有部门
	deptSource, err := wechat.GetAllDepts()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("获取企业微信部门列表失败：%s", err.Error()))
	}
	depts, err := ConvertDeptData(config.Conf.WeCom.Flag, deptSource)
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("转换企业微信部门数据失败：%s", err.Error()))
	}
	// 2.将部门这个数组进行拆分，一组是父ID为1的，一组是父ID不为1的
	var firstDepts []*model.Group // 父ID为1的部门
	var otherDepts []*model.Group // 父ID不为1的部门
	for _, dept := range depts {
		if dept.SourceDeptId == fmt.Sprintf("%s_1", config.Conf.WeCom.Flag) { // 跳过ID为1的根部门，由系统配置的根部门进行占位
			continue
		}
		if dept.SourceDeptParentId == fmt.Sprintf("%s_1", config.Conf.WeCom.Flag) {
			firstDepts = append(firstDepts, dept)
		} else {
			otherDepts = append(otherDepts, dept)
		}
	}
	// 3.先写父ID为1的，再写父ID不为1的
	for _, dept := range firstDepts {
		err := d.AddDepts(dept)
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncWeComDepts添加根部门失败：%s", err.Error()))
		}
	}

	for _, dept := range otherDepts {
		err := d.AddDepts(dept)
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncWeComDepts添加根部门失败：%s", err.Error()))
		}
	}
	return nil, nil
}

// AddGroup 添加部门数据
func (d WeComLogic) AddDepts(group *model.Group) error {
	// 判断部门名称是否存在
	parentGroup := new(model.Group)
	err := isql.Group.Find(tools.H{"source_dept_id": group.SourceDeptParentId}, parentGroup)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("查询父级部门失败：%s", err.Error()))
	}
	if !isql.Group.Exist(tools.H{"source_dept_id": group.SourceDeptId}) {
		// 此时的 group 已经附带了Build后动态关联好的字段，接下来将一些确定性的其他字段值添加上，就可以创建这个分组了
		group.Creator = "system"
		group.GroupType = "cn"
		group.ParentId = parentGroup.ID
		group.Source = config.Conf.WeCom.Flag
		group.GroupDN = fmt.Sprintf("cn=%s,%s", group.GroupName, parentGroup.GroupDN)
		err = CommonAddGroup(group)
		if err != nil {
			return tools.NewOperationError(fmt.Errorf("添加部门失败：%s", err.Error()))
		}
	}
	return nil
}

//根据现有数据库同步到的部门信息，开启用户同步
func (d WeComLogic) SyncWeComUsers(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	// 1.获取企业微信用户列表
	staffSource, err := wechat.GetAllUsers()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("获取企业微信用户列表失败：%s", err.Error()))
	}
	staffs, err := ConvertUserData(config.Conf.WeCom.Flag, staffSource)
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("转换企业微信用户数据失败：%s", err.Error()))
	}
	// 2.遍历用户，开始写入
	for _, staff := range staffs {
		// 入库
		err = d.AddUsers(staff)
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncWeComUsers写入用户失败：%s", err.Error()))
		}
	}

	// 3.获取企业微信已离职用户id列表
	// 拿到MySQL所有用户数据，远程没有的，则说明被删除了
	var res []*model.User
	users, err := isql.User.List(&request.UserListReq{})
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户列表失败：" + err.Error()))
	}
	for _, user := range users {
		if user.Username == "admin" {
			continue
		}
		in := true
		for _, staff := range staffs {
			if user.Username == staff.Username {
				in = false
				break
			}
		}
		if in {
			res = append(res, user)
		}
	}
	// 4.遍历id，开始处理
	for _, userTmp := range res {
		user := new(model.User)
		err = isql.User.Find(tools.H{"source_user_id": userTmp.SourceUserId}, user)
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
	return nil, nil
}

// AddUser 添加用户数据
func (d WeComLogic) AddUsers(user *model.User) error {
	// 根据 unionid 查询用户,不存在则创建
	if !isql.User.Exist(tools.H{"source_union_id": user.SourceUnionId}) {
		// 根据角色id获取角色
		roles, err := isql.Role.GetRolesByIds([]uint{2})
		if err != nil {
			return tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败:%s", err.Error()))
		}
		user.Creator = "system"
		user.Roles = roles
		user.Password = config.Conf.Ldap.UserInitPassword
		user.Source = config.Conf.WeCom.Flag
		user.UserDN = fmt.Sprintf("uid=%s,%s", user.Username, config.Conf.Ldap.UserDN)
		err = CommonAddUser(user, tools.StringToSlice(user.DepartmentId, ","))
		if err != nil {
			return err
		}
	}
	return nil
}
