package feishu

import (
	"context"

	"github.com/chyroc/lark"
)

// 官方文档： https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/children
// GetAllDepts 获取所有部门
func GetAllDepts() (depts []*lark.GetDepartmentListRespItem, err error) {
	var (
		fetchChild bool  = true
		pageSize   int64 = 50
	)

	req := lark.GetDepartmentListReq{
		FetchChild:   &fetchChild,
		PageSize:     &pageSize,
		DepartmentID: "0"}

	for {
		res, _, err := InitFeiShuClient().Contact.GetDepartmentList(context.TODO(), &req)
		if err != nil {
			return nil, err
		}
		depts = append(depts, res.Items...)
		if !res.HasMore {
			break
		}
		req.PageToken = &res.PageToken
	}
	return
}

// 官方文档： https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/find_by_department
// GetAllUsers 获取所有员工信息
func GetAllUsers() (users []*lark.GetUserListRespItem, err error) {
	var (
		pageSize int64 = 50
	)
	depts, err := GetAllDepts()
	if err != nil {
		return nil, err
	}
	for _, dept := range depts {
		req := lark.GetUserListReq{
			PageSize:     &pageSize,
			PageToken:    new(string),
			DepartmentID: dept.OpenDepartmentID,
		}
		for {
			res, _, err := InitFeiShuClient().Contact.GetUserList(context.Background(), &req)
			if err != nil {
				return nil, err
			}
			users = append(users, res.Items...)
			if !res.HasMore {
				break
			}
			req.PageToken = &res.PageToken
		}
	}
	return
}

// 官方文档： https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/employee/list
// GetLeaveUserIds 获取离职人员ID列表
func GetLeaveUserIds() ([]string, error) {
	var ids []string
	users, _, err := InitFeiShuClient().EHR.GetEHREmployeeList(context.TODO(), &lark.GetEHREmployeeListReq{
		Status:     []int64{5},
		UserIDType: lark.IDTypePtr(lark.IDTypeUnionID), // 只查询unionID
	})
	if err != nil {
		return nil, err
	}
	for _, user := range users.Items {
		ids = append(ids, user.UserID)
	}
	return ids, nil
}
