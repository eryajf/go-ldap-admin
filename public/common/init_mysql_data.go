package common

import (
	"errors"

	"github.com/eryajf-world/go-ldap-admin/config"
	"github.com/eryajf-world/go-ldap-admin/model"
	"github.com/eryajf-world/go-ldap-admin/public/tools"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

// 初始化mysql数据
func InitData() {
	// 是否初始化数据
	if !config.Conf.System.InitData {
		return
	}

	// 1.写入角色数据
	newRoles := make([]*model.Role, 0)
	roles := []*model.Role{
		{
			Model:   gorm.Model{ID: 1},
			Name:    "管理员",
			Keyword: "admin",
			Remark:  "",
			Sort:    1,
			Status:  1,
			Creator: "系统",
		},
		{
			Model:   gorm.Model{ID: 2},
			Name:    "普通用户",
			Keyword: "user",
			Remark:  "",
			Sort:    3,
			Status:  1,
			Creator: "系统",
		},
		{
			Model:   gorm.Model{ID: 3},
			Name:    "访客",
			Keyword: "guest",
			Remark:  "",
			Sort:    5,
			Status:  1,
			Creator: "系统",
		},
	}

	for _, role := range roles {
		err := DB.First(&role, role.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newRoles = append(newRoles, role)
		}
	}

	if len(newRoles) > 0 {
		err := DB.Create(&newRoles).Error
		if err != nil {
			Log.Errorf("写入系统角色数据失败：%v", err)
		}
	}

	// 2写入菜单
	newMenus := make([]model.Menu, 0)
	var uint0 uint = 0
	var uint1 uint = 1
	var uint4 uint = 4
	var uint8 uint = 8
	componentStr := "component"
	systemRoleStr := "/system/role"
	personnelManageStr := "/personnel/user"
	userStr := "user"
	peopleStr := "people"
	groupStr := "peoples"
	roleStr := "eye-open"
	treeTableStr := "tree-table"
	treeStr := "tree"
	exampleStr := "example"
	logOperationStr := "/log/operation-log"
	documentationStr := "documentation"
	menus := []model.Menu{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "UserManage",
			Title:     "人员管理",
			Icon:      userStr,
			Path:      "/personnel",
			Component: "Layout",
			Redirect:  personnelManageStr,
			Sort:      6,
			ParentId:  uint0,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "User",
			Title:     "用户管理",
			Icon:      peopleStr,
			Path:      "user",
			Component: "/personnel/user/index",
			Sort:      7,
			ParentId:  uint1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 3},
			Name:      "Group",
			Title:     "分组管理",
			Icon:      groupStr,
			Path:      "group",
			Component: "/personnel/group/index",
			Sort:      8,
			ParentId:  uint1,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 4},
			Name:      "System",
			Title:     "系统管理",
			Icon:      componentStr,
			Path:      "/system",
			Component: "Layout",
			Redirect:  systemRoleStr,
			Sort:      9,
			ParentId:  uint0,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 5},
			Name:      "Role",
			Title:     "角色管理",
			Icon:      roleStr,
			Path:      "role",
			Component: "/system/role/index",
			Sort:      10,
			ParentId:  uint4,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 6},
			Name:      "Menu",
			Title:     "菜单管理",
			Icon:      treeTableStr,
			Path:      "menu",
			Component: "/system/menu/index",
			Sort:      13,
			ParentId:  uint4,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 7},
			Name:      "Api",
			Title:     "接口管理",
			Icon:      treeStr,
			Path:      "api",
			Component: "/system/api/index",
			Sort:      14,
			ParentId:  uint4,
			Roles:     roles[:1],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 8},
			Name:      "Log",
			Title:     "日志管理",
			Icon:      exampleStr,
			Path:      "/log",
			Component: "Layout",
			Redirect:  logOperationStr,
			Sort:      20,
			ParentId:  uint0,
			Roles:     roles[:2],
			Creator:   "系统",
		},
		{
			Model:     gorm.Model{ID: 9},
			Name:      "OperationLog",
			Title:     "操作日志",
			Icon:      documentationStr,
			Path:      "operation-log",
			Component: "/log/operation-log/index",
			Sort:      21,
			ParentId:  uint8,
			Roles:     roles[:2],
			Creator:   "系统",
		},
	}
	for _, menu := range menus {
		err := DB.First(&menu, menu.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newMenus = append(newMenus, menu)
		}
	}
	if len(newMenus) > 0 {
		err := DB.Create(&newMenus).Error
		if err != nil {
			Log.Errorf("写入系统菜单数据失败：%v", err)
		}
	}

	// 3.写入用户
	newUsers := make([]model.User, 0)
	users := []model.User{
		{
			Model:         gorm.Model{ID: 1},
			Username:      "admin",
			Password:      tools.NewGenPasswd(config.Conf.Ldap.LdapAdminPass),
			Nickname:      "管理员",
			GivenName:     "最强后台",
			Mail:          "admin@eryajf.net",
			JobNumber:     "0000",
			Mobile:        "18888888888",
			Avatar:        "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			PostalAddress: "中国河南省南阳市",
			Departments:   "运维部",
			Position:      "系统管理员",
			Introduction:  "最强后台的管理员",
			Status:        1,
			Creator:       "系统",
			Roles:         roles[:1],
		},
	}

	for _, user := range users {
		err := DB.First(&user, user.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newUsers = append(newUsers, user)
		}
	}

	if len(newUsers) > 0 {
		err := DB.Create(&newUsers).Error
		if err != nil {
			Log.Errorf("写入用户数据失败：%v", err)
		}
	}

	// 4.写入api
	apis := []model.Api{
		{
			Method:   "POST",
			Path:     "/base/login",
			Category: "base",
			Remark:   "用户登录",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/base/logout",
			Category: "base",
			Remark:   "用户登出",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/base/refreshToken",
			Category: "base",
			Remark:   "刷新JWT令牌",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/base/changePwd",
			Category: "base",
			Remark:   "通过邮箱修改密码",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/user/info",
			Category: "user",
			Remark:   "获取当前登录用户信息",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/user/list",
			Category: "user",
			Remark:   "获取用户列表",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/user/changePwd",
			Category: "user",
			Remark:   "更新用户登录密码",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/user/add",
			Category: "user",
			Remark:   "创建用户",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/user/update",
			Category: "user",
			Remark:   "更新用户",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/user/delete",
			Category: "user",
			Remark:   "批量删除用户",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/user/changeUserStatus",
			Category: "user",
			Remark:   "更改用户在职状态",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/group/list",
			Category: "group",
			Remark:   "获取分组列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/group/tree",
			Category: "group",
			Remark:   "获取分组列表树",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/group/add",
			Category: "group",
			Remark:   "创建分组",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/group/update",
			Category: "group",
			Remark:   "更新分组",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/group/delete",
			Category: "group",
			Remark:   "批量删除分组",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/group/adduser",
			Category: "group",
			Remark:   "添加用户到分组",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/group/removeuser",
			Category: "group",
			Remark:   "将用户从分组移出",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/group/useringroup",
			Category: "group",
			Remark:   "获取在分组内的用户列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/group/usernoingroup",
			Category: "group",
			Remark:   "获取不在分组内的用户列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/role/list",
			Category: "role",
			Remark:   "获取角色列表",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/role/add",
			Category: "role",
			Remark:   "创建角色",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/role/update",
			Category: "role",
			Remark:   "更新角色",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/role/getmenulist",
			Category: "role",
			Remark:   "获取角色的权限菜单",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/role/updatemenus",
			Category: "role",
			Remark:   "更新角色的权限菜单",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/role/getapilist",
			Category: "role",
			Remark:   "获取角色的权限接口",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/role/updateapis",
			Category: "role",
			Remark:   "更新角色的权限接口",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/role/delete",
			Category: "role",
			Remark:   "批量删除角色",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/menu/list",
			Category: "menu",
			Remark:   "获取菜单列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/menu/tree",
			Category: "menu",
			Remark:   "获取菜单树",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/menu/add",
			Category: "menu",
			Remark:   "创建菜单",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/menu/update",
			Category: "menu",
			Remark:   "更新菜单",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/menu/delete",
			Category: "menu",
			Remark:   "批量删除菜单",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/api/list",
			Category: "api",
			Remark:   "获取接口列表",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/api/tree",
			Category: "api",
			Remark:   "获取接口树",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/api/add",
			Category: "api",
			Remark:   "创建接口",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/api/update",
			Category: "api",
			Remark:   "更新接口",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/api/delete",
			Category: "api",
			Remark:   "批量删除接口",
			Creator:  "系统",
		},
		{
			Method:   "GET",
			Path:     "/log/operation/list",
			Category: "log",
			Remark:   "获取操作日志列表",
			Creator:  "系统",
		},
		{
			Method:   "POST",
			Path:     "/log/operation/delete",
			Category: "log",
			Remark:   "批量删除操作日志",
			Creator:  "系统",
		},
	}

	// 5. 将角色绑定给菜单
	newApi := make([]model.Api, 0)
	newRoleCasbin := make([]model.RoleCasbin, 0)
	for i, api := range apis {
		api.ID = uint(i + 1)
		err := DB.First(&api, api.ID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newApi = append(newApi, api)

			// 管理员拥有所有API权限
			newRoleCasbin = append(newRoleCasbin, model.RoleCasbin{
				Keyword: roles[0].Keyword,
				Path:    api.Path,
				Method:  api.Method,
			})

			// 非管理员拥有基础权限
			basePaths := []string{
				"/base/login",
				"/base/logout",
				"/base/refreshToken",
				"/base/changePwd",
				"/base/dashboard",
				"/user/info",
				"/user/list",
				"/user/changePwd",
				"/group/list",
				"/group/tree",
				"/group/useringroup",
				"/group/usernoingroup",
				"/role/list",
				"/role/getmenulist",
				"/role/getapilist",
				"/menu/tree",
				"/menu/list",
				"/api/list",
				"/api/tree",
				"/log/operation/list",
			}

			if funk.ContainsString(basePaths, api.Path) {
				newRoleCasbin = append(newRoleCasbin, model.RoleCasbin{
					Keyword: roles[1].Keyword,
					Path:    api.Path,
					Method:  api.Method,
				})
			}
		}
	}

	if len(newApi) > 0 {
		if err := DB.Create(&newApi).Error; err != nil {
			Log.Errorf("写入api数据失败：%v", err)
		}
	}

	if len(newRoleCasbin) > 0 {
		rules := make([][]string, 0)
		for _, c := range newRoleCasbin {
			rules = append(rules, []string{
				c.Keyword, c.Path, c.Method,
			})
		}
		isAdd, err := CasbinEnforcer.AddPolicies(rules)
		if !isAdd {
			Log.Errorf("写入casbin数据失败：%v", err)
		}
	}
}
