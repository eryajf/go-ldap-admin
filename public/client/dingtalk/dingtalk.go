package dingtalk

import (
	"strings"

	"github.com/mozillazg/go-pinyin"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

// 官方文档地址： https://open.dingtalk.com/document/orgapp-server/obtain-the-department-list
// GetAllDepts 获取所有部门
func GetAllDepts(deptId int) (result []*DingTalkDept, err error) {
	depts, err := InitDingTalkClient().FetchDeptList(deptId, true, "zh_CN")
	if err != nil {
		return result, err
	}

	for _, dept := range depts.Dept {
		result = append(result, &DingTalkDept{
			Id:       dept.Id,
			Name:     strings.Join(pinyin.LazyConvert(dept.Name, nil), ""),
			Remark:   dept.Name,
			ParentId: dept.ParentId,
		})
	}
	return
}

// 官方文档地址： https://open.dingtalk.com/document/orgapp-server/queries-the-complete-information-of-a-department-user
// GetAllUsers 获取所有员工信息
func GetAllUsers() (result []*DingTalkUser, err error) {
	depts, err := GetAllDepts(1)
	if err != nil {
		return nil, err
	}

	for _, dept := range depts {
		r := request.DeptDetailUserInfo{
			DeptId:   dept.Id,
			Cursor:   0,
			Size:     99,
			Language: "zh_CN",
		}
		for {
			//获取钉钉部门人员信息
			rsp, err := InitDingTalkClient().GetDeptDetailUserInfo(&r)
			if err != nil {
				return nil, err
			}
			for _, user := range rsp.DeptDetailUsers {
				result = append(result, &DingTalkUser{
					UserId:               user.UserId,
					UnionId:              user.UnionId,
					Name:                 user.Name,
					Avatar:               user.Avatar,
					StateCode:            user.StateCode,
					ManagerUserId:        user.ManagerUserId,
					Mobile:               user.Mobile,
					HideMobile:           user.HideMobile,
					Telephone:            user.Telephone,
					JobNumber:            user.JobNumber,
					Title:                user.Title,
					WorkPlace:            user.WorkPlace,
					Remark:               user.Remark,
					LoginId:              user.LoginId,
					DeptIds:              user.DeptIds,
					DeptOrder:            user.DeptOrder,
					Extension:            user.Extension,
					HiredDate:            user.HiredDate,
					Active:               user.Active,
					Admin:                user.Admin,
					Boss:                 user.Boss,
					ExclusiveAccount:     user.ExclusiveAccount,
					Leader:               user.Leader,
					ExclusiveAccountType: user.ExclusiveAccountType,
					OrgEmail:             user.OrgEmail,
					Email:                user.Email,
				})
			}
			if !rsp.HasMore {
				break
			}
			r.Cursor = rsp.NextCursor
		}
	}
	return
}

// 官方文档：https://open.dingtalk.com/document/orgapp-server/intelligent-personnel-query-company-turnover-list
// GetLeaveUserIds 获取离职人员ID列表
func GetLeaveUserIds() ([]string, error) {
	var ids []string
	ReqParm := struct {
		Cursor int `json:"cursor"`
		Size   int `json:"size"`
	}{
		Cursor: 0,
		Size:   50,
	}

	for {
		rsp, err := InitDingTalkClient().GetHrmResignEmployeeIds(ReqParm.Cursor, ReqParm.Size)
		if err != nil {
			return nil, err
		}
		ids = append(ids, rsp.UserIds...)
		if rsp.NextCursor == 0 {
			break
		}
		ReqParm.Cursor = rsp.NextCursor
	}
	return ids, nil
}
