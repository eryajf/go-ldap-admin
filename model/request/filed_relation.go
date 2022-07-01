package request

// FieldRelationListReq 获取资源列表结构体
type FieldRelationListReq struct {
}

// FieldRelationAddReq 添加资源结构体
type FieldRelationAddReq struct {
	Flag       string            `json:"flag" validate:"required,min=1,max=20"`
	Attributes map[string]string `json:"attributes" validate:"required,gt=0"`
}

// FieldRelationUpdateReq 更新资源结构体
type FieldRelationUpdateReq struct {
	ID         uint              `json:"id" validate:"required"`
	Flag       string            `json:"flag" validate:"required,min=1,max=20"`
	Attributes map[string]string `json:"attributes" validate:"required,gt=0"`
}

// FieldRelationDeleteReq 删除资源结构体
type FieldRelationDeleteReq struct {
	FieldRelationIds []uint `json:"fieldRelationIds" validate:"required"`
}
