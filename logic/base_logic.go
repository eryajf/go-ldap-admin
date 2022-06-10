package logic

import (
	"fmt"

	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/model/request"
	"github.com/eryajf/go-ldap-admin/model/response"
	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/eryajf/go-ldap-admin/service/ildap"
	"github.com/eryajf/go-ldap-admin/service/isql"

	"github.com/gin-gonic/gin"
)

type BaseLogic struct{}

// Add 添加数据
func (l BaseLogic) ChangePwd(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.BaseChangePwdReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	// 判断邮箱是否正确
	if !isql.User.Exist(tools.H{"mail": r.Mail}) {
		return nil, tools.NewValidatorError(fmt.Errorf("邮箱不存在,请检查邮箱是否正确"))
	}

	user := new(model.User)
	err := isql.User.Find(tools.H{"mail": r.Mail}, user)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("通过邮箱查询用户失败" + err.Error()))
	}

	newpass, err := ildap.User.NewPwd(user.Username)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("LDAP生成新密码失败" + err.Error()))
	}

	err = tools.SendMail([]string{user.Mail}, newpass)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("发送邮件失败" + err.Error()))
	}

	// 更新数据库密码
	err = isql.User.ChangePwd(user.Username, tools.NewGenPasswd(newpass))
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL更新密码失败: " + err.Error()))
	}

	return nil, nil
}

// Dashboard 仪表盘
func (l BaseLogic) Dashboard(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	_, ok := req.(*request.BaseDashboardReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	userCount, err := isql.User.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户总数失败"))
	}
	groupCount, err := isql.Group.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取分组总数失败"))
	}
	roleCount, err := isql.Role.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取角色总数失败"))
	}
	menuCount, err := isql.Menu.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取菜单总数失败"))
	}
	apiCount, err := isql.Api.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取接口总数失败"))
	}
	logCount, err := isql.OperationLog.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取日志总数失败"))
	}

	rst := make([]*response.DashboardList, 0)

	rst = append(rst,
		&response.DashboardList{
			DataType:  "user",
			DataName:  "用户",
			DataCount: userCount,
			Icon:      "people",
		},
		&response.DashboardList{
			DataType:  "group",
			DataName:  "分组",
			DataCount: groupCount,
			Icon:      "peoples",
		},
		&response.DashboardList{
			DataType:  "role",
			DataName:  "角色",
			DataCount: roleCount,
			Icon:      "eye-open",
		},
		&response.DashboardList{
			DataType:  "menu",
			DataName:  "菜单",
			DataCount: menuCount,
			Icon:      "tree-table",
		},
		&response.DashboardList{
			DataType:  "api",
			DataName:  "接口",
			DataCount: apiCount,
			Icon:      "tree",
		},
		&response.DashboardList{
			DataType:  "log",
			DataName:  "日志",
			DataCount: logCount,
			Icon:      "documentation",
		},
	)

	return rst, nil
}
