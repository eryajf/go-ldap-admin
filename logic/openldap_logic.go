package logic

import (
	"fmt"

	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/public/client/openldap"

	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/eryajf/go-ldap-admin/service/isql"
	"github.com/gin-gonic/gin"
)

type OpenLdapLogic struct {
}

//通过ldap获取部门信息
func (d *OpenLdapLogic) SyncOpenLdapDepts(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	// 1.获取所有部门
	depts, err := openldap.GetAllDepts()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("获取ldap部门列表失败：%s", err.Error()))
	}
	// 2.将部门这个数组进行拆分，一组是父ID为根的，一组是父ID不为根的
	var firstDepts []*openldap.Dept // 父ID为根的部门
	var otherDepts []*openldap.Dept // 父ID不为根的部门
	for _, dept := range depts {
		if dept.ParentId == "1" {
			firstDepts = append(firstDepts, dept)
		} else {
			otherDepts = append(otherDepts, dept)
		}
	}
	// 3.先写父ID为根的，再写父ID不为根的
	for _, dept := range firstDepts {
		err := d.AddDepts(&model.Group{
			GroupName:          dept.Name,
			Remark:             dept.Remark,
			Creator:            "system",
			GroupType:          "cn",
			SourceDeptId:       dept.Id,
			Source:             "openldap",
			SourceDeptParentId: dept.ParentId,
			GroupDN:            dept.DN,
		})
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncOpenLdapDepts添加根部门失败：%s", err.Error()))
		}
	}

	for _, dept := range otherDepts {
		err := d.AddDepts(&model.Group{
			GroupName:          dept.Name,
			Remark:             dept.Remark,
			Creator:            "system",
			GroupType:          "cn",
			SourceDeptId:       dept.Id,
			Source:             "openldap",
			SourceDeptParentId: dept.ParentId,
			GroupDN:            dept.DN,
		})
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncOpenLdapDepts添加其他部门失败：%s", err.Error()))
		}
	}
	return nil, nil
}

// AddGroup 添加部门数据
func (d OpenLdapLogic) AddDepts(group *model.Group) error {
	// 判断部门名称是否存在
	parentGroup := new(model.Group)
	err := isql.Group.Find(tools.H{"source_dept_id": group.SourceDeptParentId}, parentGroup)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("查询父级部门失败：%s", err.Error()))
	}
	if !isql.Group.Exist(tools.H{"source_dept_id": group.SourceDeptId}) {
		group.ParentId = parentGroup.ID
		// 在数据库中创建组
		err = isql.Group.Add(group)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据现有数据库同步到的部门信息，开启用户同步
func (d OpenLdapLogic) SyncOpenLdapUsers(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	// 1.获取ldap用户列表
	staffs, err := openldap.GetAllUsers()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("获取ldap用户列表失败：%s", err.Error()))
	}
	// 2.遍历用户，开始写入
	for _, staff := range staffs {
		groupIds, err := isql.Group.DeptIdsToGroupIds(staff.DepartmentIds)
		if err != nil {
			return nil, tools.NewMySqlError(fmt.Errorf("将部门ids转换为内部部门id失败：%s", err.Error()))
		}
		// 根据角色id获取角色
		roles, err := isql.Role.GetRolesByIds([]uint{2})
		if err != nil {
			return nil, tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败:%s", err.Error()))
		}
		// 入库
		err = d.AddUsers(&model.User{
			Username:      staff.Name,
			Nickname:      staff.DisplayName,
			GivenName:     staff.GivenName,
			Mail:          staff.Mail,
			JobNumber:     staff.EmployeeNumber,
			Mobile:        staff.Mobile,
			PostalAddress: staff.PostalAddress,
			Departments:   staff.BusinessCategory,
			Position:      staff.DepartmentNumber,
			Introduction:  staff.CN,
			Creator:       "system",
			Source:        "openldap",
			DepartmentId:  tools.SliceToString(groupIds, ","),
			SourceUserId:  staff.Name,
			SourceUnionId: staff.Name,
			Roles:         roles,
			UserDN:        staff.DN,
		})
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncOpenLdapUsers写入用户失败：%s", err.Error()))
		}
	}
	return nil, nil
}

// AddUser 添加用户数据
func (d OpenLdapLogic) AddUsers(user *model.User) error {
	// 根据 unionid 查询用户,不存在则创建
	if !isql.User.Exist(tools.H{"source_union_id": user.SourceUnionId}) {
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
		if user.JobNumber == "" {
			user.JobNumber = "未启用"
		}
		// 先将用户添加到MySQL
		err := isql.User.Add(user)
		if err != nil {
			return tools.NewMySqlError(fmt.Errorf("向MySQL创建用户失败：" + err.Error()))
		}

		// 获取用户将要添加的分组
		groups, err := isql.Group.GetGroupByIds(tools.StringToSlice(user.DepartmentId, ","))
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
		}
		return nil
	}
	return nil
}
