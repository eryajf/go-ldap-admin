package logic

import (
	"errors"
	"fmt"

	"github.com/eryajf-world/go-ldap-admin/config"
	"github.com/eryajf-world/go-ldap-admin/model"
	"github.com/eryajf-world/go-ldap-admin/public/tools"
	"github.com/eryajf-world/go-ldap-admin/service/ildap"
	"github.com/eryajf-world/go-ldap-admin/service/isql"
	"github.com/eryajf-world/go-ldap-admin/svc/request"
	"github.com/eryajf-world/go-ldap-admin/svc/response"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type UserLogic struct{}

// Add 添加数据
func (l UserLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.UserAddReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	if isql.User.Exist(tools.H{"username": r.Username}) {
		return nil, tools.NewValidatorError(fmt.Errorf("用户名已存在,请勿重复添加"))
	}
	if isql.User.Exist(tools.H{"mobile": r.Mobile}) {
		return nil, tools.NewValidatorError(fmt.Errorf("手机号已存在,请勿重复添加"))
	}
	if isql.User.Exist(tools.H{"job_number": r.JobNumber}) {
		return nil, tools.NewValidatorError(fmt.Errorf("工号已存在,请勿重复添加"))
	}
	if isql.User.Exist(tools.H{"mail": r.Mail}) {
		return nil, tools.NewValidatorError(fmt.Errorf("邮箱已存在,请勿重复添加"))
	}

	// 密码通过RSA解密
	// 密码不为空就解密
	if r.Password != "" {
		decodeData, err := tools.RSADecrypt([]byte(r.Password), config.Conf.System.RSAPrivateBytes)
		if err != nil {
			return nil, tools.NewValidatorError(fmt.Errorf("密码解密失败"))
		}
		r.Password = string(decodeData)
		if len(r.Password) < 6 {
			return nil, tools.NewValidatorError(fmt.Errorf("密码长度至少为6位"))
		}
	} else {
		r.Password = "123456"
	}

	// 当前登陆用户角色排序最小值（最高等级角色）以及当前登陆的用户
	currentRoleSortMin, ctxUser, err := isql.User.GetCurrentUserMinRoleSort(c)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("获取当前登陆用户角色排序最小值失败"))
	}

	// 根据角色id获取角色
	if r.RoleIds == nil || len(r.RoleIds) == 0 {
		r.RoleIds = []uint{2} // 默认添加为普通用户角色
	}

	roles, err := isql.Role.GetRolesByIds(r.RoleIds)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败"))
	}

	var reqRoleSorts []int
	for _, role := range roles {
		reqRoleSorts = append(reqRoleSorts, int(role.Sort))
	}
	// 前端传来用户角色排序最小值（最高等级角色）
	reqRoleSortMin := uint(funk.MinInt(reqRoleSorts).(int))

	// 当前用户的角色排序最小值 需要小于 前端传来的角色排序最小值（用户不能创建比自己等级高的或者相同等级的用户）
	if currentRoleSortMin >= reqRoleSortMin {
		return nil, tools.NewValidatorError(fmt.Errorf("用户不能创建比自己等级高的或者相同等级的用户"))
	}

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
		Creator:       ctxUser.Username,
		DepartmentId:  r.DepartmentId,
		Source:        r.Source,
		Roles:         roles,
	}
	if user.Source == "" {
		user.Source = "platform"
	}
	//先识别用户选择的部门是否是OU开头
	dn, err := isql.Group.GetGroupDn(r.DepartmentId, "")
	if err != nil {
		return nil, err.Error()
	}
	gdn := fmt.Sprintf("%s,%s", dn, config.Conf.Ldap.LdapBaseDN)
	if gdn[:3] == "ou=" {
		return nil, errors.New("不能添加用户到OU组织单元")
	}
	//先创建用户到默认分组
	err = ildap.User.Add(&user)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("向LDAP创建用户失败：" + err.Error()))
	}
	//根据选择的部门，添加到部门内
	err = ildap.Group.AddUserToGroup(gdn, fmt.Sprintf("uid=%s,%s", user.Username, config.Conf.Ldap.LdapUserDN))
	if err != nil {
		return nil, err.Error()
	}
	err = isql.User.Add(&user)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("向MySQL创建用户失败：" + err.Error()))
	}
	//根据部门分配，将用户和部门信息维护到部门关系表里面
	users := []model.User{}
	users = append(users, user)
	depart := new(model.Group)
	filter := tools.H{"id": int(r.DepartmentId)}
	err = isql.Group.Find(filter, depart)
	if err != nil {
		return "", tools.NewMySqlError(err)
	}
	err = isql.Group.AddUserToGroup(depart, users)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("向MySQL添加用户到分组关系失败：" + err.Error()))
	}
	return nil, nil
}

// List 数据列表
func (l UserLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.UserListReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	users, err := isql.User.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户列表失败：" + err.Error()))
	}

	rets := make([]model.User, 0)
	for _, user := range users {
		rets = append(rets, *user)
	}
	count, err := isql.User.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户总数失败：" + err.Error()))
	}

	return response.UserListRsp{
		Total: int(count),
		Users: rets,
	}, nil
}

// Update 更新数据
func (l UserLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.UserUpdateReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	if !isql.User.Exist(tools.H{"id": r.ID}) {
		return nil, tools.NewMySqlError(fmt.Errorf("该记录不存在"))
	}

	// 获取当前登陆用户
	ctxUser, err := isql.User.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户失败"))
	}
	// 获取当前登陆用户的所有角色
	currentRoles := ctxUser.Roles
	// 获取当前登陆用户角色的排序，和前端传来的角色排序做比较
	var currentRoleSorts []int
	// 当前登陆用户角色ID集合
	var currentRoleIds []uint
	for _, role := range currentRoles {
		currentRoleSorts = append(currentRoleSorts, int(role.Sort))
		currentRoleIds = append(currentRoleIds, role.ID)
	}
	// 当前登陆用户角色排序最小值（最高等级角色）
	currentRoleSortMin := funk.MinInt(currentRoleSorts).(int)

	// 获取前端传来的用户角色id
	reqRoleIds := r.RoleIds
	// 根据角色id获取角色
	roles, err := isql.Role.GetRolesByIds(reqRoleIds)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败"))
	}
	if len(roles) == 0 {
		return nil, tools.NewValidatorError(fmt.Errorf("未获取到角色信息"))
	}
	var reqRoleSorts []int
	for _, role := range roles {
		reqRoleSorts = append(reqRoleSorts, int(role.Sort))
	}
	// 前端传来用户角色排序最小值（最高等级角色）
	reqRoleSortMin := funk.MinInt(reqRoleSorts).(int)

	oldData := new(model.User)
	err = isql.User.Find(tools.H{"id": r.ID}, oldData)
	if err != nil {
		return nil, tools.NewMySqlError(err)
	}

	user := model.User{
		Model:         oldData.Model,
		Username:      r.Username,
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
		Creator:       ctxUser.Username,
		DepartmentId:  r.DepartmentId,
		Source:        oldData.Source,
		Roles:         roles,
	}

	// 判断是更新自己还是更新别人,如果操作的ID与登陆用户的ID一致，则说明操作的是自己
	if int(r.ID) == int(ctxUser.ID) {
		// 不能更改自己的角色
		reqDiff, currentDiff := funk.Difference(r.RoleIds, currentRoleIds)
		if len(reqDiff.([]uint)) > 0 || len(currentDiff.([]uint)) > 0 {
			return nil, tools.NewValidatorError(fmt.Errorf("不能更改自己的角色"))
		}
	} else {
		// 如果是更新别人，操作者不能更新比自己角色等级高的或者相同等级的用户
		// 根据userIdID获取用户角色排序最小值
		minRoleSorts, err := isql.User.GetUserMinRoleSortsByIds([]uint{uint(r.ID)})
		if err != nil || len(minRoleSorts) == 0 {
			return nil, tools.NewValidatorError(fmt.Errorf("根据用户ID获取用户角色排序最小值失败"))
		}
		if currentRoleSortMin >= minRoleSorts[0] {
			return nil, tools.NewValidatorError(fmt.Errorf("用户不能更新比自己角色等级高的或者相同等级的用户"))
		}
		// 用户不能把别的用户角色等级更新得比自己高或相等
		if currentRoleSortMin >= reqRoleSortMin {
			return nil, tools.NewValidatorError(fmt.Errorf("用户不能把别的用户角色等级更新得比自己高或相等"))
		}
	}
	err = ildap.User.Update(oldData.Username, &user)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("在LDAP更新用户失败：" + err.Error()))
	}

	// 更新用户
	if !config.Conf.Ldap.LdapUserNameModify {
		user.Username = oldData.Username
	}
	err = isql.User.Update(&user)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL更新用户失败：" + err.Error()))
	}
	//判断部门信息是否有变化有变化则更新相应的数据库
	if oldData.DepartmentId != r.DepartmentId {
		//从旧组中删除
		err = l.RemoveUserToGroup(oldData.DepartmentId, []uint{r.ID})
		//添加到新分组中
		err = l.AddUserToGroup(r.DepartmentId, []uint{r.ID})
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// RemoveUser 移除用户
func (l UserLogic) RemoveUserToGroup(groupId uint, userIds []uint) error {
	filter := tools.H{"id": groupId}

	if !isql.Group.Exist(filter) {
		return tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	users, err := isql.User.GetUserByIds(userIds)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("获取用户列表失败: %s", err.Error()))
	}

	group := new(model.Group)
	err = isql.Group.Find(filter, group)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}
	gdn, err := isql.Group.GetGroupDn(groupId, "")
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}
	gdn = fmt.Sprintf("%s,%s", gdn, config.Conf.Ldap.LdapBaseDN)
	for _, user := range users {
		err := ildap.Group.RemoveUserFromGroup(gdn, user.Username)
		if err != nil {
			return tools.NewLdapError(fmt.Errorf("将用户从ldap移除失败" + err.Error()))
		}
	}
	err = isql.Group.RemoveUserFromGroup(group, users)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("将用户从MySQL移除失败: %s", err.Error()))
	}

	return nil
}

//将用户添加到分组中
func (l UserLogic) AddUserToGroup(groupId uint, userIds []uint) error {
	filter := tools.H{"id": groupId}

	if !isql.Group.Exist(filter) {
		return tools.NewMySqlError(fmt.Errorf("分组不存在"))
	}

	users, err := isql.User.GetUserByIds(userIds)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("获取用户列表失败: %s", err.Error()))
	}

	group := new(model.Group)
	err = isql.Group.Find(filter, group)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("获取分组失败: %s", err.Error()))
	}
	for _, user := range users {
		gdn, err := isql.Group.GetGroupDn(group.ID, "")
		if err != nil {
			return err
		}
		gdn = fmt.Sprintf("%s,%s", gdn, config.Conf.Ldap.LdapBaseDN)
		udn := config.Conf.Ldap.LdapAdminDN
		if user.Username != "admin" {
			udn = fmt.Sprintf("uid=%s,%s", user.Username, config.Conf.Ldap.LdapUserDN)
		}
		err = ildap.Group.AddUserToGroup(gdn, udn)

		if err != nil {
			return tools.NewLdapError(fmt.Errorf("向LDAP添加用户到分组失败" + err.Error()))
		}
	}
	err = isql.Group.AddUserToGroup(group, users)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("添加用户到分组失败: %s", err.Error()))
	}
	return nil
}

// Delete 删除数据
func (l UserLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.UserDeleteReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	// 根据用户ID获取用户角色排序最小值
	roleMinSortList, err := isql.User.GetUserMinRoleSortsByIds(r.UserIds) // TODO: 这里应该复用下边的 GetUserByIds 方法
	if err != nil || len(roleMinSortList) == 0 {
		return nil, tools.NewValidatorError(fmt.Errorf("根据用户ID获取用户角色排序最小值失败"))
	}

	// 获取当前登陆用户角色排序最小值（最高等级角色）以及当前用户
	minSort, ctxUser, err := isql.User.GetCurrentUserMinRoleSort(c)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("获取当前登陆用户角色排序最小值失败"))
	}
	currentRoleSortMin := int(minSort)

	// 不能删除自己
	if funk.Contains(r.UserIds, ctxUser.ID) {
		return nil, tools.NewValidatorError(fmt.Errorf("用户不能删除自己"))
	}

	// 不能删除比自己(登陆用户)角色排序低(等级高)的用户
	for _, sort := range roleMinSortList {
		if currentRoleSortMin >= sort {
			return nil, tools.NewValidatorError(fmt.Errorf("用户不能删除比自己角色等级高的用户"))
		}
	}

	users, err := isql.User.GetUserByIds(r.UserIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户信息失败: " + err.Error()))
	}
	for _, user := range users {
		err := ildap.User.Delete(user.Username)
		if err != nil {
			return nil, tools.NewLdapError(fmt.Errorf("在LDAP删除用户失败" + err.Error()))
		}
	}

	err = isql.User.Delete(r.UserIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL删除用户失败: " + err.Error()))
	}

	return nil, nil
}

// ChangePwd 修改密码
func (l UserLogic) ChangePwd(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.UserChangePwdReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c
	// 前端传来的密码是rsa加密的,先解密
	// 密码通过RSA解密
	decodeOldPassword, err := tools.RSADecrypt([]byte(r.OldPassword), config.Conf.System.RSAPrivateBytes)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("原密码解析失败"))
	}
	decodeNewPassword, err := tools.RSADecrypt([]byte(r.NewPassword), config.Conf.System.RSAPrivateBytes)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("新密码解析失败"))
	}
	r.OldPassword = string(decodeOldPassword)
	r.NewPassword = string(decodeNewPassword)
	// 获取当前用户
	user, err := isql.User.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户失败"))
	}
	// 获取用户的真实正确密码
	// correctPasswd := user.Password
	// 判断前端请求的密码是否等于真实密码
	// err = tools.ComparePasswd(correctPasswd, r.OldPassword)
	// if err != nil {
	// 	return nil, tools.NewValidatorError(fmt.Errorf("原密码错误"))
	// }
	if tools.NewParPasswd(user.Password) != r.OldPassword {
		return nil, tools.NewValidatorError(fmt.Errorf("原密码错误"))
	}
	// ldap更新密码时可以直接指定用户DN和新密码即可更改成功
	err = ildap.User.ChangePwd(user.Username, "", r.NewPassword)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("在LDAP更新密码失败" + err.Error()))
	}

	// 更新密码
	err = isql.User.ChangePwd(user.Username, tools.NewGenPasswd(r.NewPassword))
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL更新密码失败: " + err.Error()))
	}

	return nil, nil
}

// ChangeUserStatus 修改用户状态
func (l UserLogic) ChangeUserStatus(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.UserChangeUserStatusReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c
	// 校验工作
	filter := tools.H{"id": r.ID}
	if !isql.User.Exist(filter) {
		return nil, tools.NewValidatorError(fmt.Errorf("该用户不存在"))
	}
	user := new(model.User)
	err := isql.User.Find(filter, user)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL查询用户失败: " + err.Error()))
	}

	if r.Status == 1 && r.Status == user.Status {
		return nil, tools.NewValidatorError(fmt.Errorf("用户已经是在职状态"))
	}
	if r.Status == 2 && r.Status == user.Status {
		return nil, tools.NewValidatorError(fmt.Errorf("用户已经是离职状态"))
	}

	// 获取当前登录用户，只有管理员才能够将用户状态改变
	// 获取当前登陆用户角色排序最小值（最高等级角色）以及当前用户
	minSort, _, err := isql.User.GetCurrentUserMinRoleSort(c)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("获取当前登陆用户角色排序最小值失败"))
	}

	if int(minSort) != 1 {
		return nil, tools.NewValidatorError(fmt.Errorf("只有管理员才能更改用户状态"))
	}

	if r.Status == 2 {
		err = ildap.User.Delete(user.Username)
		if err != nil {
			return nil, tools.NewLdapError(fmt.Errorf("在LDAP删除用户失败" + err.Error()))
		}
	} else {
		err = ildap.User.Add(user)
		if err != nil {
			return nil, tools.NewLdapError(fmt.Errorf("在LDAP添加用户失败" + err.Error()))
		}
	}
	err = isql.User.ChangeStatus(int(r.ID), int(r.Status))
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL更新用户状态失败: " + err.Error()))
	}
	return nil, nil
}

// GetUserInfo 获取用户信息
func (l UserLogic) GetUserInfo(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.UserGetUserInfoReq)
	if !ok {
		return nil, ReqAssertErr
	}

	_ = c
	_ = r

	user, err := isql.User.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前用户信息失败: " + err.Error()))
	}
	return user, nil
}
