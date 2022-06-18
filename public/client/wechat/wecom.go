package wechat

import (
	"strconv"

	"github.com/wenerme/go-wecom/wecom"
)

// GetAllDepts 获取所有部门
func GetAllDepts() ([]wecom.ListDepartmentResponseItem, error) {
	depts, err := InitWeComClient().ListDepartment(
		&wecom.ListDepartmentRequest{},
	)
	if err != nil {
		return nil, err
	}
	return depts.Department, nil
}

// GetAllDepts 获取所有部门
func GetAllUsers() ([]wecom.ListUserResponseItem, error) {
	depts, err := GetAllDepts()
	if err != nil {
		return nil, err
	}
	var us []wecom.ListUserResponseItem
	for _, dept := range depts {
		users, err := InitWeComClient().ListUser(
			&wecom.ListUserRequest{
				DepartmentID: strconv.Itoa(dept.ID),
				FetchChild:   "1",
			},
		)
		if err != nil {
			return nil, err
		}
		us = append(us, users.UserList...)
	}
	return us, nil
}
