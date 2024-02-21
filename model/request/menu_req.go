package request

// MenuAddReq 添加资源结构体
type MenuAddReq struct {
	Name       string `json:"name" validate:"required,min=1,max=50"`
	Title      string `json:"title" validate:"required,min=1,max=50"`
	Icon       string `json:"icon" validate:"min=0,max=50"`
	Path       string `json:"path" validate:"required,min=1,max=100"`
	Redirect   string `json:"redirect" validate:"min=0,max=100"`
	Component  string `json:"component" validate:"required,min=1,max=100"`
	Sort       uint   `json:"sort" validate:"gte=1,lte=999"`
	Status     uint   `json:"status" validate:"oneof=1 2"`
	Hidden     uint   `json:"hidden" validate:"oneof=1 2"`
	NoCache    uint   `json:"noCache" validate:"oneof=1 2"`
	AlwaysShow uint   `json:"alwaysShow" validate:"oneof=1 2"`
	Breadcrumb uint   `json:"breadcrumb" validate:"oneof=1 2"`
	ActiveMenu string `json:"activeMenu" validate:"min=0,max=100"`
	ParentId   uint   `json:"parentId"`
}

// MenuListReq 列表结构体
type MenuListReq struct {
}

// MenuUpdateReq 更新资源结构体
type MenuUpdateReq struct {
	ID         uint   `json:"id" validate:"required"`
	Name       string `json:"name" validate:"required,min=1,max=50"`
	Title      string `json:"title" validate:"required,min=1,max=50"`
	Icon       string `json:"icon" validate:"min=0,max=50"`
	Path       string `json:"path" validate:"required,min=1,max=100"`
	Redirect   string `json:"redirect" validate:"min=0,max=100"`
	Component  string `json:"component" validate:"min=0,max=100"`
	Sort       uint   `json:"sort" validate:"gte=1,lte=999"`
	Status     uint   `json:"status" validate:"oneof=1 2"`
	Hidden     uint   `json:"hidden" validate:"oneof=1 2"`
	NoCache    uint   `json:"noCache" validate:"oneof=1 2"`
	AlwaysShow uint   `json:"alwaysShow" validate:"oneof=1 2"`
	Breadcrumb uint   `json:"breadcrumb" validate:"oneof=1 2"`
	ActiveMenu string `json:"activeMenu" validate:"min=0,max=100"`
	ParentId   uint   `json:"parentId" validate:"gte=0"`
}

// MenuDeleteReq 删除资源结构体
type MenuDeleteReq struct {
	MenuIds []uint `json:"menuIds" validate:"required"`
}

// MenuGetTreeReq 获取菜单树结构体
type MenuGetTreeReq struct {
}

// MenuGetAccessTreeReq 获取用户菜单树
type MenuGetAccessTreeReq struct {
	ID uint `json:"id" form:"id"`
}
