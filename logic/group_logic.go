package logic

import (
	"fmt"
	"strings"

	"github.com/eryajf/go-ldap-admin/config"

	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/eryajf/go-ldap-admin/service/ildap"
	"github.com/eryajf/go-ldap-admin/service/isql"
	"github.com/eryajf/go-ldap-admin/svc/request"
	"github.com/eryajf/go-ldap-admin/svc/response"

	"github.com/gin-gonic/gin"
)

type GroupLogic struct{}

// Add 添加数据
func (l GroupLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.GroupAddReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	if isql.Group.Exist(tools.H{"group_name": r.GroupName}) {
		return nil, tools.NewValidatorError(fmt.Errorf("组名已存在"))
	}

	// 获取当前用户
	ctxUser, err := isql.User.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}

	group := model.Group{
		GroupType:          r.GroupType,
		ParentId:           r.ParentId,
		GroupName:          r.GroupName,
		Remark:             r.Remark,
		Creator:            ctxUser.Username,
		Source:             "platform", //默认是平台添加
		SourceDeptId:       "platform_0",
		SourceDeptParentId: "platform_0",
	}
	pdn := ""
	if group.ParentId > 0 {
		pdn, err = isql.Group.GetGroupDn(r.ParentId, "")
		if err != nil {
			return nil, err.Error()
		}
	}
	err = ildap.Group.Add(&group, pdn)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("向LDAP创建分组失败" + err.Error()))
	}

	// 创建
	err = isql.Group.Add(&group)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("向MySQL创建分组失败"))
	}

	// 默认创建分组之后，需要将admin添加到分组中
	adminInfo := new(model.User)
	err = isql.User.Find(tools.H{"id": 1}, adminInfo)
	if err != nil {
		return nil, tools.NewMySqlError(err)
	}

	err = isql.Group.AddUserToGroup(&group, []model.User{*adminInfo})
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("添加用户到分组失败: %s", err.Error()))
	}

	return nil, nil
}

// List 数据列表
func (l GroupLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.GroupListReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	// 获取数据列表
	groups, err := isql.Group.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组列表失败: %s", err.Error()))
	}

	rets := make([]model.Group, 0)
	for _, group := range groups {
		rets = append(rets, *group)
	}
	count, err := isql.Group.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组总数失败"))
	}

	return response.GroupListRsp{
		Total:  count,
		Groups: rets,
	}, nil
}

// GetTree 数据树
func (l GroupLogic) GetTree(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.GroupListReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	var groups []*model.Group
	groups, err := isql.Group.ListTree(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取资源列表失败: " + err.Error()))
	}

	tree := isql.GenGroupTree(0, groups)

	return tree, nil
}

// Update 更新数据
func (l GroupLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.GroupUpdateReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": int(r.ID)}
	if !isql.Group.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	// 获取当前登陆用户
	ctxUser, err := isql.User.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户失败"))
	}

	oldData := new(model.Group)
	err = isql.Group.Find(filter, oldData)
	if err != nil {
		return nil, tools.NewMySqlError(err)
	}

	group := model.Group{
		Model:     oldData.Model,
		GroupName: r.GroupName,
		Remark:    r.Remark,
		Creator:   ctxUser.Username,
		GroupType: oldData.GroupType,
	}

	oldGroupName := oldData.GroupName
	oldRemark := oldData.Remark
	dn, err := isql.Group.GetGroupDn(r.ID, "")
	if err != nil {
		return nil, err.Error()
	}
	err = ildap.Group.Update(&group, dn, oldGroupName, oldRemark)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("向LDAP更新分组失败：" + err.Error()))
	}
	//若配置了不允许修改分组名称，则不更新分组名称
	if !config.Conf.Ldap.LdapGroupNameModify {
		group.GroupName = oldGroupName
	}
	err = isql.Group.Update(&group)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("向MySQL更新分组失败"))
	}
	return nil, nil
}

// Delete 删除数据
func (l GroupLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.GroupDeleteReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	for _, id := range r.GroupIds {
		filter := tools.H{"id": int(id)}
		if !isql.Group.Exist(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("分组不存在"))
		}
	}

	groups, err := isql.Group.GetGroupByIds(r.GroupIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组列表失败: %s", err.Error()))
	}
	for _, group := range groups {
		// 判断存在子分组，不允许删除
		filter := tools.H{"parent_id": int(group.ID)}
		if isql.Group.Exist(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("存在子分组，不允许删除！"))
		}
		dn, err := isql.Group.GetGroupDn(group.ID, "")
		if err != nil {
			return nil, err.Error()
		}
		gdn := fmt.Sprintf("%s,%s", dn, config.Conf.Ldap.LdapBaseDN)
		err = ildap.Group.Delete(gdn)
		if err != nil {
			return nil, tools.NewLdapError(fmt.Errorf("向LDAP删除分组失败：" + err.Error()))
		}
	}
	// 删除接口
	err = isql.Group.Delete(r.GroupIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除接口失败: %s", err.Error()))
	}
	// TODO: 删除用户组关系

	return nil, nil
}

// AddUser 添加用户到分组
func (l GroupLogic) AddUser(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.GroupAddUserReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": r.GroupID}

	if !isql.Group.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	users, err := isql.User.GetUserByIds(r.UserIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户列表失败: %s", err.Error()))
	}

	group := new(model.Group)
	err = isql.Group.Find(filter, group)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}

	for _, user := range users {
		gdn, err := isql.Group.GetGroupDn(group.ID, "")
		if err != nil {
			return nil, err.Error()
		}
		gdn = fmt.Sprintf("%s,%s", gdn, config.Conf.Ldap.LdapBaseDN)
		udn := config.Conf.Ldap.LdapAdminDN
		if user.Username != "admin" {
			udn = fmt.Sprintf("uid=%s,%s", user.Username, config.Conf.Ldap.LdapUserDN)
		}
		err = ildap.Group.AddUserToGroup(gdn, udn)

		if err != nil {
			return nil, tools.NewLdapError(fmt.Errorf("向LDAP添加用户到分组失败" + err.Error()))
		}
	}

	err = isql.Group.AddUserToGroup(group, users)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("添加用户到分组失败: %s", err.Error()))
	}

	return nil, nil
}

// RemoveUser 移除用户
func (l GroupLogic) RemoveUser(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.GroupRemoveUserReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": r.GroupID}

	if !isql.Group.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	users, err := isql.User.GetUserByIds(r.UserIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户列表失败: %s", err.Error()))
	}

	group := new(model.Group)
	err = isql.Group.Find(filter, group)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}
	gdn, err := isql.Group.GetGroupDn(r.GroupID, "")
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}
	gdn = fmt.Sprintf("%s,%s", gdn, config.Conf.Ldap.LdapBaseDN)
	for _, user := range users {
		err := ildap.Group.RemoveUserFromGroup(gdn, user.Username)
		if err != nil {
			return nil, tools.NewLdapError(fmt.Errorf("将用户从ldap移除失败" + err.Error()))
		}
	}
	err = isql.Group.RemoveUserFromGroup(group, users)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("将用户从MySQL移除失败: %s", err.Error()))
	}

	return nil, nil
}

// UserInGroup 在分组内的用户
func (l GroupLogic) UserInGroup(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.UserInGroupReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": r.GroupID}

	if !isql.Group.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	group := new(model.Group)
	err := isql.Group.Find(filter, group)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}

	rets := make([]response.Guser, 0)

	for _, user := range group.Users {
		if r.Nickname != "" && !strings.Contains(user.Nickname, r.Nickname) {
			continue
		}
		rets = append(rets, response.Guser{
			UserId:       int64(user.ID),
			UserName:     user.Username,
			NickName:     user.Nickname,
			Mail:         user.Mail,
			JobNumber:    user.JobNumber,
			Mobile:       user.Mobile,
			Introduction: user.Introduction,
		})
	}

	return response.GroupUsers{
		GroupId:     int64(group.ID),
		GroupName:   group.GroupName,
		GroupRemark: group.Remark,
		UserList:    rets,
	}, nil
}

// UserNoInGroup 不在分组内的用户
func (l GroupLogic) UserNoInGroup(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.UserNoInGroupReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": r.GroupID}

	if !isql.Group.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	group := new(model.Group)
	err := isql.Group.Find(filter, group)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}

	var userList []*model.User
	userList, err = isql.User.ListAll()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取资源列表失败: " + err.Error()))
	}

	rets := make([]response.Guser, 0)
	for _, user := range userList {
		in := true
		for _, groupUser := range group.Users {
			if user.Username == groupUser.Username {
				in = false
				break
			}
		}
		if in {
			if r.Nickname != "" && !strings.Contains(user.Nickname, r.Nickname) {
				continue
			}
			rets = append(rets, response.Guser{
				UserId:       int64(user.ID),
				UserName:     user.Username,
				NickName:     user.Nickname,
				Mail:         user.Mail,
				JobNumber:    user.JobNumber,
				Mobile:       user.Mobile,
				Introduction: user.Introduction,
			})
		}
	}

	return response.GroupUsers{
		GroupId:     int64(group.ID),
		GroupName:   group.GroupName,
		GroupRemark: group.Remark,
		UserList:    rets,
	}, nil
}
