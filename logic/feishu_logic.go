package logic

import (
	"fmt"
	"strings"

	"github.com/chyroc/lark"
	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/model/request"
	"github.com/eryajf/go-ldap-admin/public/client/feishu"
	"github.com/mozillazg/go-pinyin"

	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/eryajf/go-ldap-admin/service/ildap"
	"github.com/eryajf/go-ldap-admin/service/isql"
	"github.com/gin-gonic/gin"
)

type FeiShuLogic struct {
}

//通过飞书获取部门信息
func (d *FeiShuLogic) SyncFeiShuDepts(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	// 1.获取所有部门
	depts, err := feishu.GetAllDepts()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("获取飞书部门列表失败：%s", err.Error()))
	}
	// 2.将部门这个数组进行拆分，一组是父ID为根的，一组是父ID不为根的
	var firstDepts []*lark.GetDepartmentListRespItem // 父ID为根的部门
	var otherDepts []*lark.GetDepartmentListRespItem // 父ID不为根的部门
	for _, dept := range depts {
		if dept.ParentDepartmentID == "0" {
			firstDepts = append(firstDepts, dept)
		} else {
			otherDepts = append(otherDepts, dept)
		}
	}
	// 3.先写父ID为根的，再写父ID不为根的
	for _, dept := range firstDepts {
		err := d.AddDepts(&request.WeComGroupAddReq{
			GroupType:          "cn",
			GroupName:          strings.Join(pinyin.LazyConvert(dept.Name, nil), ""),
			Remark:             dept.Name,
			SourceDeptId:       fmt.Sprintf("%s_%s", config.Conf.FeiShu.Flag, dept.OpenDepartmentID),
			Source:             config.Conf.FeiShu.Flag,
			SourceDeptParentId: fmt.Sprintf("%s_%d", config.Conf.FeiShu.Flag, 1),
		})
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncFeiShuDepts添加根部门失败：%s", err.Error()))
		}
	}

	for _, dept := range otherDepts {
		err := d.AddDepts(&request.WeComGroupAddReq{
			GroupType:          "cn",
			GroupName:          strings.Join(pinyin.LazyConvert(dept.Name, nil), ""),
			Remark:             dept.Name,
			SourceDeptId:       fmt.Sprintf("%s_%s", config.Conf.FeiShu.Flag, dept.OpenDepartmentID),
			Source:             config.Conf.FeiShu.Flag,
			SourceDeptParentId: fmt.Sprintf("%s_%s", config.Conf.FeiShu.Flag, dept.ParentDepartmentID),
		})
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncFeiShuDepts添加根部门失败：%s", err.Error()))
		}
	}
	return nil, nil
}

// AddGroup 添加部门数据
func (d FeiShuLogic) AddDepts(r *request.WeComGroupAddReq) error {
	// 判断部门名称是否存在
	parentGroup := new(model.Group)
	err := isql.Group.Find(tools.H{"source_dept_id": r.SourceDeptParentId}, parentGroup)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("查询父级部门失败：%s", err.Error()))
	}
	if !isql.Group.Exist(tools.H{"source_dept_id": r.SourceDeptId}) {
		groupTmp := model.Group{
			GroupName:          r.GroupName,
			Remark:             r.Remark,
			Creator:            "system",
			GroupType:          "cn",
			ParentId:           parentGroup.ID,
			SourceDeptId:       r.SourceDeptId,
			Source:             r.Source,
			SourceDeptParentId: r.SourceDeptParentId,
			GroupDN:            fmt.Sprintf("cn=%s,%s", r.GroupName, parentGroup.GroupDN),
		}
		err = CommonAddGroup(&groupTmp)
		if err != nil {
			return tools.NewOperationError(fmt.Errorf("添加部门失败：%s", err.Error()))
		}
	}
	// todo: 分组存在，但是信息有变更的情况，需要考量，但是这种组织架构的调整，通常是比较复杂的情况，这里并不好与之一一对应同步，暂时不做支持
	return nil
}

//根据现有数据库同步到的部门信息，开启用户同步
func (d FeiShuLogic) SyncFeiShuUsers(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	// 1.获取飞书用户列表
	users, err := feishu.GetAllUsers()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("SyncFeiShuUsers获取飞书用户列表失败：%s", err.Error()))
	}
	// 2.遍历用户，开始写入
	for _, detail := range users {
		// 用户名的几种情况
		var userName string
		if detail.Email != "" {
			userName = strings.Split(detail.Email, "@")[0]
		}
		if userName == "" && detail.Name != "" {
			userName = strings.Join(pinyin.LazyConvert(detail.Name, nil), "")
		}
		if userName == "" && detail.Mobile != "" {
			userName = detail.Mobile
		}
		if userName == "" && detail.Email != "" {
			userName = strings.Split(detail.Email, "@")[0]
		}

		// 如果企业内没有工号，则工号用名字占位
		if detail.EmployeeNo == "" {
			detail.EmployeeNo = detail.Mobile
		}

		//飞书部门ids,转换为内部部门id
		var sourceDeptIds []string
		for _, deptId := range detail.DepartmentIDs {
			sourceDeptIds = append(sourceDeptIds, fmt.Sprintf("%s_%s", config.Conf.FeiShu.Flag, deptId))
		}
		groupIds, err := isql.Group.DeptIdsToGroupIds(sourceDeptIds)
		if err != nil {
			return nil, tools.NewMySqlError(fmt.Errorf("SyncFeiShuUsers获取飞书部门ids转换为内部部门id失败：%s", err.Error()))
		}

		// 写入用户
		user := request.WeComUserAddReq{
			Username:      userName,
			Password:      config.Conf.Ldap.UserInitPassword,
			Nickname:      detail.Name,
			GivenName:     detail.Name,
			Mail:          detail.Email,
			JobNumber:     detail.Name, // 工号暂用名字替代
			Mobile:        detail.Mobile[3:],
			Avatar:        detail.Avatar.AvatarOrigin,
			PostalAddress: detail.City,
			Position:      detail.JobTitle,
			Introduction:  detail.Name,
			Status:        1,
			DepartmentId:  groupIds,
			Source:        config.Conf.FeiShu.Flag,
			SourceUserId:  fmt.Sprintf("%s_%s", config.Conf.FeiShu.Flag, detail.UserID),
			SourceUnionId: fmt.Sprintf("%s_%s", config.Conf.FeiShu.Flag, detail.UnionID),
		}
		// 入库
		err = d.AddUsers(&user)
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncFeiShuUsers写入用户失败：%s", err.Error()))
		}
	}

	// 3.获取飞书已离职用户id列表
	userIds, err := feishu.GetLeaveUserIds()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("SyncFeiShuUsers获取飞书离职用户列表失败：%s", err.Error()))
	}
	// 4.遍历id，开始处理
	for _, uid := range userIds {
		user := new(model.User)
		err = isql.User.Find(tools.H{"source_union_id": fmt.Sprintf("%s_%s", config.Conf.FeiShu.Flag, uid)}, user)
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
func (d FeiShuLogic) AddUsers(r *request.WeComUserAddReq) error {
	// 根据 unionid 查询用户,不存在则创建
	if !isql.User.Exist(tools.H{"source_union_id": r.SourceUnionId}) {
		// 根据角色id获取角色
		r.RoleIds = []uint{2} // 默认添加为普通用户角色
		roles, err := isql.Role.GetRolesByIds(r.RoleIds)
		if err != nil {
			return tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败:%s", err.Error()))
		}

		deptIds := tools.SliceToString(r.DepartmentId, ",")
		user := model.User{
			Username:      r.Username,
			Password:      r.Password,
			Nickname:      r.Nickname,
			GivenName:     r.GivenName,
			Mail:          r.Mail,
			JobNumber:     r.JobNumber,
			Mobile:        r.Mobile,
			Avatar:        r.Avatar,
			PostalAddress: r.PostalAddress,
			Departments:   r.Departments,
			Position:      r.Position,
			Introduction:  r.Introduction,
			Status:        r.Status,
			Creator:       "system",
			DepartmentId:  deptIds,
			Roles:         roles,
			Source:        r.Source,
			SourceUserId:  r.SourceUserId,
			SourceUnionId: r.SourceUnionId,
			UserDN:        fmt.Sprintf("uid=%s,%s", r.Username, config.Conf.Ldap.UserDN),
		}
		err = CommonAddUser(&user, r.DepartmentId)
		if err != nil {
			return err
		}
	}
	// todo: 用户如果存在，则暂时跳过，目前用户名取自邮箱等内容，因为这个不确定性，可能会造成一些逻辑上的问题，因为默认情况下，用户名是无法在ldap中更改的，所以暂时跳过，如果用户有这里的需求，可以根据自己的情况固定用户名的字段，也就可以打开如下的注释了
	// else {
	// 	oldData := new(model.User)
	// 	if err := isql.User.Find(tools.H{"source_union_id": r.SourceUnionId}, oldData); err != nil {
	// 		return err
	// 	}
	// 	if r.Username != oldData.Username || r.Mail != oldData.Mail || r.Mobile != oldData.Mobile {
	// 		user := model.User{
	// 			Model:         oldData.Model,
	// 			Username:      r.Username,
	// 			Nickname:      r.Nickname,
	// 			GivenName:     r.GivenName,
	// 			Mail:          r.Mail,
	// 			JobNumber:     r.JobNumber,
	// 			Mobile:        r.Mobile,
	// 			Avatar:        r.Avatar,
	// 			PostalAddress: r.PostalAddress,
	// 			Departments:   r.Departments,
	// 			Position:      r.Position,
	// 			Introduction:  r.Introduction,
	// 			Creator:       oldData.Creator,
	// 			DepartmentId:  tools.SliceToString(r.DepartmentId, ","),
	// 			Source:        oldData.Source,
	// 			Roles:         oldData.Roles,
	// 			UserDN:        oldData.UserDN,
	// 		}
	// 		if err := CommonUpdateUser(oldData, &user, r.DepartmentId); err != nil {
	// 			return err
	// 		}
	// 	}
	// }
	return nil
}
