package logic

import (
	"fmt"
	"strings"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/model/request"
	"github.com/eryajf/go-ldap-admin/public/client/wechat"
	"github.com/mozillazg/go-pinyin"
	"github.com/wenerme/go-wecom/wecom"

	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/eryajf/go-ldap-admin/service/isql"
	"github.com/gin-gonic/gin"
)

type WeComLogic struct {
}

//通过钉钉获取部门信息
func (d *WeComLogic) SyncWeComDepts(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	// 1.获取所有部门
	depts, err := wechat.GetAllDepts()
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("获取企业微信部门列表失败：%s", err.Error()))
	}
	// 2.将部门这个数组进行拆分，一组是父ID为1的，一组是父ID不为1的
	var firstDepts []wecom.ListDepartmentResponseItem // 父ID为1的部门
	var otherDepts []wecom.ListDepartmentResponseItem // 父ID不为1的部门
	for _, dept := range depts {
		if dept.ID == 1 { // 跳过ID为1的根部门，由系统配置的根部门进行占位
			continue
		}
		if dept.ParentID == 1 {
			firstDepts = append(firstDepts, dept)
		} else {
			otherDepts = append(otherDepts, dept)
		}
	}
	// 3.先写父ID为1的，再写父ID不为1的
	for _, dept := range firstDepts {
		err := d.AddDepts(&request.DingGroupAddReq{
			GroupType:          "cn",
			GroupName:          strings.Join(pinyin.LazyConvert(dept.Name, nil), ""),
			Remark:             dept.Name,
			SourceDeptId:       fmt.Sprintf("%s_%d", config.Conf.WeCom.Flag, dept.ID),
			Source:             config.Conf.WeCom.Flag,
			SourceDeptParentId: fmt.Sprintf("%s_%d", config.Conf.WeCom.Flag, 1),
		})
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncWeComDepts添加根部门失败：%s", err.Error()))
		}
	}

	for _, dept := range otherDepts {
		err := d.AddDepts(&request.DingGroupAddReq{
			GroupType:          "cn",
			GroupName:          strings.Join(pinyin.LazyConvert(dept.Name, nil), ""),
			Remark:             dept.Name,
			SourceDeptId:       fmt.Sprintf("%s_%d", config.Conf.WeCom.Flag, dept.ID),
			Source:             config.Conf.WeCom.Flag,
			SourceDeptParentId: fmt.Sprintf("%s_%d", config.Conf.WeCom.Flag, dept.ParentID),
		})
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("SyncWeComDepts添加根部门失败：%s", err.Error()))
		}
	}
	return nil, nil
}

// AddGroup 添加部门数据
func (d WeComLogic) AddDepts(r *request.DingGroupAddReq) error {
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
func (d WeComLogic) SyncWeComUsers(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {

	return nil, nil
}

// AddUser 添加用户数据
func (d WeComLogic) AddUsers(r *request.DingUserAddReq) error {
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
