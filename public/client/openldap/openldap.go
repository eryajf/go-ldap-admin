package openldap

import (
	"fmt"
	"strings"

	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/public/common"
	ldap "github.com/go-ldap/ldap/v3"
)

type Dept struct {
	DN       string `json:"dn"`
	Id       string `json:"id"`       // 部门ID
	Name     string `json:"name"`     // 部门名称拼音
	Remark   string `json:"remark"`   // 部门中文名
	ParentId string `json:"parentid"` // 父部门ID
}

type User struct {
	Name             string   `json:"name"`
	DN               string   `json:"dn"`
	CN               string   `json:"cn"`
	SN               string   `json:"sn"`
	Mobile           string   `json:"mobile"`
	BusinessCategory string   `json:"businessCategory"` // 业务类别，部门名字
	DepartmentNumber string   `json:"departmentNumber"` // 部门编号，此处可以存放员工的职位
	Description      string   `json:"description"`      // 描述
	DisplayName      string   `json:"displayName"`      // 展示名字，可以是中文名字
	Mail             string   `json:"mail"`             // 邮箱
	EmployeeNumber   string   `json:"employeeNumber"`   // 员工工号
	GivenName        string   `json:"givenName"`        // 给定名字，如果公司有花名，可以用这个字段
	PostalAddress    string   `json:"postalAddress"`    // 家庭住址
	DepartmentIds    []string `json:"department_ids"`
}

// GetAllDepts 获取所有部门
func GetAllDepts() (ret []*Dept, err error) {
	// Construct query request
	searchRequest := ldap.NewSearchRequest(
		config.Conf.Ldap.BaseDN,                                     // This is basedn, we will start searching from this node.
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, // Here several parameters are respectively scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		"(&(objectClass=*))", // This is Filter for LDAP query
		[]string{},           // Here are the attributes returned by the query, provided as an array. If empty, all attributes are returned
		nil,
	)

	// 获取 LDAP 连接
	conn, err := common.GetLDAPConn()
	defer common.PutLADPConn(conn)
	if err != nil {
		return nil, err
	}

	// Search through ldap built-in search
	sr, err := conn.Search(searchRequest)
	if err != nil {
		return ret, err
	}
	// Refers to the entry that returns data. If it is greater than 0, the interface returns normally.
	if len(sr.Entries) > 0 {
		for _, v := range sr.Entries {
			if v.DN == config.Conf.Ldap.BaseDN || v.DN == config.Conf.Ldap.AdminDN || strings.Contains(v.DN, config.Conf.Ldap.UserDN) {
				continue
			}
			var ele Dept
			ele.DN = v.DN
			ele.Name = strings.Split(strings.Split(v.DN, ",")[0], "=")[1]
			ele.Id = strings.Split(strings.Split(v.DN, ",")[0], "=")[1]
			ele.Remark = v.GetAttributeValue("description")
			if len(strings.Split(v.DN, ","))-len(strings.Split(config.Conf.Ldap.BaseDN, ",")) == 1 {
				ele.ParentId = "0"
			} else {
				ele.ParentId = strings.Split(strings.Split(v.DN, ",")[1], "=")[1]
			}
			ret = append(ret, &ele)
		}
	}
	return
}

// GetAllUsers 获取所有员工信息
func GetAllUsers() (ret []*User, err error) {
	// Construct query request
	searchRequest := ldap.NewSearchRequest(
		config.Conf.Ldap.BaseDN,                                     // This is basedn, we will start searching from this node.
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, // Here several parameters are respectively scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		"(&(objectClass=*))", // This is Filter for LDAP query
		[]string{},           // Here are the attributes returned by the query, provided as an array. If empty, all attributes are returned
		nil,
	)

	// 获取 LDAP 连接
	conn, err := common.GetLDAPConn()
	defer common.PutLADPConn(conn)
	if err != nil {
		return nil, err
	}

	// Search through ldap built-in search
	sr, err := conn.Search(searchRequest)
	if err != nil {
		return ret, err
	}
	// Refers to the entry that returns data. If it is greater than 0, the interface returns normally.
	if len(sr.Entries) > 0 {
		for _, v := range sr.Entries {
			if v.DN == config.Conf.Ldap.UserDN || !strings.Contains(v.DN, config.Conf.Ldap.UserDN) {
				continue
			}
			name := strings.Split(strings.Split(v.DN, ",")[0], "=")[1]
			deptIds, err := GetUserDeptIds(v.DN)
			if err != nil {
				return ret, err
			}
			ret = append(ret, &User{
				Name:             name,
				DN:               v.DN,
				CN:               v.GetAttributeValue("cn"),
				SN:               v.GetAttributeValue("sn"),
				Mobile:           v.GetAttributeValue("mobile"),
				BusinessCategory: v.GetAttributeValue("businessCategory"),
				DepartmentNumber: v.GetAttributeValue("departmentNumber"),
				Description:      v.GetAttributeValue("description"),
				DisplayName:      v.GetAttributeValue("displayName"),
				Mail:             v.GetAttributeValue("mail"),
				EmployeeNumber:   v.GetAttributeValue("employeeNumber"),
				GivenName:        v.GetAttributeValue("givenName"),
				PostalAddress:    v.GetAttributeValue("postalAddress"),
				DepartmentIds:    deptIds,
			})
		}
	}
	return
}

// GetUserDeptIds 获取用户所在的部门
func GetUserDeptIds(udn string) (ret []string, err error) {
	// Construct query request
	searchRequest := ldap.NewSearchRequest(
		config.Conf.Ldap.BaseDN,                                     // This is basedn, we will start searching from this node.
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, // Here several parameters are respectively scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		fmt.Sprintf("(|(Member=%s)(uniqueMember=%s))", udn, udn), // This is Filter for LDAP query
		[]string{}, // Here are the attributes returned by the query, provided as an array. If empty, all attributes are returned
		nil,
	)

	// 获取 LDAP 连接
	conn, err := common.GetLDAPConn()
	defer common.PutLADPConn(conn)
	if err != nil {
		return nil, err
	}

	// Search through ldap built-in search
	sr, err := conn.Search(searchRequest)
	if err != nil {
		return ret, err
	}
	// Refers to the entry that returns data. If it is greater than 0, the interface returns normally.
	if len(sr.Entries) > 0 {
		for _, v := range sr.Entries {
			ret = append(ret, strings.Split(strings.Split(v.DN, ",")[0], "=")[1])
		}
	}
	return ret, nil
}
