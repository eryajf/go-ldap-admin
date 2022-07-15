package dingtalk

import (
	"fmt"
	"strings"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

// 官方文档地址： https://open.dingtalk.com/document/orgapp-server/obtain-the-department-list
// GetAllDepts 获取所有部门
func GetAllDepts() (ret []map[string]interface{}, err error) {
	depts, err := InitDingTalkClient().FetchDeptList(1, true, "zh_CN")
	if err != nil {
		return ret, err
	}
	ret = make([]map[string]interface{}, 0)
	for _, dept := range depts.Dept {
		ele := make(map[string]interface{})
		ele["id"] = dept.Id
		ele["name"] = dept.Name
		ele["parentid"] = dept.ParentId
		ele["custom_name_pinyin"] = tools.ConvertToPinYin(dept.Name)
		ret = append(ret, ele)
	}
	return
}

// 官方文档地址： https://open.dingtalk.com/document/orgapp-server/queries-the-complete-information-of-a-department-user
// GetAllUsers 获取所有员工信息
func GetAllUsers() (ret []map[string]interface{}, err error) {
	depts, err := GetAllDepts()
	if err != nil {
		return nil, err
	}
	for _, dept := range depts {
		r := request.DeptDetailUserInfo{
			DeptId:   dept["id"].(int),
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
				ele := make(map[string]interface{})
				ele["userid"] = user.UserId
				ele["unionid"] = user.UnionId
				ele["custom_name_pinyin"] = tools.ConvertToPinYin(user.Name)
				ele["name"] = user.Name
				ele["avatar"] = user.Avatar
				ele["mobile"] = user.Mobile
				ele["job_number"] = user.JobNumber
				ele["title"] = user.Title
				ele["work_place"] = user.WorkPlace
				ele["remark"] = user.Remark
				ele["leader"] = user.Leader
				ele["org_email"] = user.OrgEmail
				if user.OrgEmail != "" {
					ele["custom_nickname_org_email"] = strings.Split(user.OrgEmail, "@")[0]
				}
				ele["email"] = user.Email
				if user.Email != "" {
					ele["custom_nickname_email"] = strings.Split(user.Email, "@")[0]
				}
				// 部门ids
				var sourceDeptIds []string
				for _, deptId := range user.DeptIds {
					sourceDeptIds = append(sourceDeptIds, fmt.Sprintf("%s_%d", config.Conf.DingTalk.Flag, deptId))
				}
				ele["department_ids"] = sourceDeptIds
				ret = append(ret, ele)
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
