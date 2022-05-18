package model

// 角色权限规则
type RoleCasbin struct {
	Keyword string `json:"keyword"` // 角色关键字
	Path    string `json:"path"`    // 访问路径
	Method  string `json:"method"`  // 请求方式
}
