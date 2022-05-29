package isql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/public/common"
	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/eryajf/go-ldap-admin/svc/request"

	"gorm.io/gorm"
)

type GroupService struct{}

// List 获取数据列表
func (s GroupService) List(req *request.GroupListReq) ([]*model.Group, error) {
	var list []*model.Group
	db := common.DB.Model(&model.Group{}).Order("created_at DESC")

	groupName := strings.TrimSpace(req.GroupName)
	if groupName != "" {
		db = db.Where("group_name LIKE ?", fmt.Sprintf("%%%s%%", groupName))
	}
	groupRemark := strings.TrimSpace(req.Remark)
	if groupRemark != "" {
		db = db.Where("remark LIKE ?", fmt.Sprintf("%%%s%%", groupRemark))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("Users").Find(&list).Error
	return list, err
}

// 拼装dn信息
func (s GroupService) GetGroupDn(groupId uint, oldDn string) (dn string, e error) {
	depart := new(model.Group)
	filter := tools.H{"id": int(groupId)}
	err := s.Find(filter, depart)
	if err != nil {
		return "", tools.NewMySqlError(err)
	}
	if oldDn == "" {
		dn = fmt.Sprintf("%s=%s", depart.GroupType, depart.GroupName)
	} else {
		dn = fmt.Sprintf("%s,%s=%s", oldDn, depart.GroupType, depart.GroupName)
	}
	if depart.ParentId > 0 {
		tempDn, err := s.GetGroupDn(depart.ParentId, dn)
		if err != nil {
			return dn, err
		}
		dn = tempDn
		fmt.Println(tempDn)
	}
	return dn, nil
}

// GenGroupTree 生成分组树
func GenGroupTree(parentId uint, groups []*model.Group) []*model.Group {
	tree := make([]*model.Group, 0)

	for _, g := range groups {
		if g.ParentId == parentId {
			children := GenGroupTree(g.ID, groups)
			g.Children = children
			tree = append(tree, g)
		}
	}
	return tree
}

// Count 获取数据总数
func (s GroupService) Count() (int64, error) {
	var count int64
	err := common.DB.Model(&model.Group{}).Count(&count).Error
	return count, err
}

// Add 添加资源
func (s GroupService) Add(data *model.Group) error {
	return common.DB.Create(data).Error
}

// Update 更新资源
func (s GroupService) Update(dataObj *model.Group) error {
	return common.DB.Model(dataObj).Where("id = ?", dataObj.ID).Updates(dataObj).Error
}

// Find 获取单个资源
func (s GroupService) Find(filter map[string]interface{}, data *model.Group) error {
	return common.DB.Where(filter).Preload("Users").First(&data).Error
}

// Exist 判断资源是否存在
func (s GroupService) Exist(filter map[string]interface{}) bool {
	var dataObj model.Group
	err := common.DB.Debug().Order("created_at DESC").Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Delete 批量删除
func (s GroupService) Delete(ids []uint) error {
	return common.DB.Where("id IN (?)", ids).Select("Users").Unscoped().Delete(&model.Group{}).Error
}

// GetApisById 根据接口ID获取接口列表
func (s GroupService) GetGroupByIds(ids []uint) (datas []*model.Group, err error) {
	err = common.DB.Where("id IN (?)", ids).Preload("Users").Find(&datas).Error
	return datas, err
}

// AddUserToGroup 添加用户到分组
func (s GroupService) AddUserToGroup(group *model.Group, users []model.User) error {
	return common.DB.Model(&group).Association("Users").Append(users)
}

// RemoveUserFromGroup 将用户从分组移除
func (s GroupService) RemoveUserFromGroup(group *model.Group, users []model.User) error {
	return common.DB.Model(&group).Association("Users").Delete(users)
}
