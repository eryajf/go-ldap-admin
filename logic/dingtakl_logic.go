package logic

import (
	"errors"
	"fmt"
	"github.com/eryajf/go-ldap-admin/config"
	"github.com/eryajf/go-ldap-admin/model"
	"github.com/eryajf/go-ldap-admin/public/tools"
	"github.com/eryajf/go-ldap-admin/service/ildap"
	"github.com/eryajf/go-ldap-admin/service/isql"
	"github.com/eryajf/go-ldap-admin/svc/request"
	"github.com/gin-gonic/gin"
	"github.com/mozillazg/go-pinyin"
	"github.com/zhaoyunxing92/dingtalk/v2"
	dingreq "github.com/zhaoyunxing92/dingtalk/v2/request"
	"gorm.io/gorm"
	"regexp"
	"strconv"
	"strings"
)

type DingTalkLogic struct {
}

//通过钉钉获取部门信息
func (d *DingTalkLogic) DsyncDingTalkDepts(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	client, err := dingtalk.NewClient(config.Conf.DingTalk.DingTalkAppKey, config.Conf.DingTalk.DingTalkAppSecret)
	// 先存根部门信息到数据库和ldap，钉钉根部门id为1，ldap根部门名称为：config.Conf.DingTalk.DingTalkRootOu
	r := request.DingGroupAddReq{}
	r.GroupName = config.Conf.DingTalk.DingTalkRootOuName
	r.GroupType = "ou"
	r.ParentId = 0
	r.Remark = "钉钉根部门"
	r.Source = config.Conf.DingTalk.DingTalkIdSource
	r.SourceDeptId = fmt.Sprintf("%s_%s", config.Conf.DingTalk.DingTalkIdSource, "1")
	r.SourceDeptParentId = fmt.Sprintf("%s_%s", config.Conf.DingTalk.DingTalkIdSource, "0")
	group, err := d.AddDept(&r)
	if err != nil {
		return nil, fmt.Sprintf("新增部门失败：部门名称为：%s,钉钉部门id为：%d,错误信息：%s", r.GroupName, r.SourceDeptId, err.Error())
	}
	// 获取根部门下的部门信息，进行处理
	reqDept := &dingreq.DeptList{}
	reqDept.DeptId = 1
	reqDept.Language = "zh_CN"
	err = d.GetSubDepts(client, reqDept, group.ID, r.Source)
	if err != nil {
		return nil, fmt.Sprintf("DsyncDingTalkDepts同步部门出错：%s", err.Error())
	}
	return nil, nil
}

// 通过钉钉获取部门信息，并存入数据库
func (d *DingTalkLogic) GetSubDepts(client *dingtalk.DingTalk, req *dingreq.DeptList, pgId uint, source string) error {
	// 获取子部门列表
	depts, err := client.GetDeptList(req)
	if err != nil {
		return errors.New(fmt.Sprintf("GetSubDepts获取部门列表失败：%s", err.Error()))
	}
	fmt.Println("GetSubDepts获取到的钉钉部门列表：", depts)
	// 遍历并处理当前部门信息
	for _, dept := range depts.Depts {
		//先判断分组类型,默认为cn，方便应对钉钉动态调整原本没有成员的部门加入成员后，导致我们无法增加
		localDept := request.DingGroupAddReq{
			GroupType:          "cn",
			ParentId:           pgId,
			GroupName:          dept.Name,
			Remark:             dept.Name,
			Source:             config.Conf.DingTalk.DingTalkIdSource,
			SourceDeptParentId: fmt.Sprintf("%s_%d", source, dept.ParentId),
			SourceDeptId:       fmt.Sprintf("%s_%d", source, dept.Id),
			SourceUserNum:      0,
		}
		//获取钉钉方，若部门存在人员信息，则设置为cn类型
		//reqTemp := &dingreq.DeptUserId{}
		//reqTemp.DeptId = dept.Id
		//repTemp, err := client.GetDeptUserIds(reqTemp)
		//if err != nil {
		//	return errors.New(fmt.Sprintf("GetSubDepts获取部门用户Id列表失败：%s", err.Error()))
		//}
		//fmt.Println("钉钉部门人员列表：", repTemp)
		//if len(repTemp.UserIds) > 0 {
		//	localDept.GroupType = "cn"
		//	localDept.SourceUserNum = len(repTemp.UserIds)
		//}
		// 处理部门入库
		deptTemp, err := d.AddDept(&localDept)
		if err != nil {
			return errors.New(fmt.Sprintf("GetSubDepts添加部门入库失败：%s", err.Error()))
		}
		// 递归调用
		sub := &dingreq.DeptList{}
		sub.DeptId = dept.Id
		sub.Language = "zh_CN"
		d.GetSubDepts(client, sub, deptTemp.ID, deptTemp.Source)
	}
	return nil
}

//根据现有数据库同步到的部门信息，开启用户同步
func (d DingTalkLogic) SyncDingTalkUsers(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	client, err := dingtalk.NewClient(config.Conf.DingTalk.DingTalkAppKey, config.Conf.DingTalk.DingTalkAppSecret)
	//获取数据库里面的钉钉同步过来的部门信息
	r := request.GroupListAllReq{}
	r.GroupType = "cn"
	r.Source = config.Conf.DingTalk.DingTalkIdSource
	depts, err := isql.Group.ListAll(&r)
	if err != nil {
		return nil, fmt.Sprintf("SyncDingTalkUsers查询本地部门列表失败：", err.Error())
	}
	//遍历处理部门，获取钉钉对应的用户信息
	for index, dept := range depts {
		fmt.Println(fmt.Sprintf("当前进行的步调为：%d,部门名称为：%s", index, dept.GroupName))
		err = d.AddDeptUser(client, dept, 0)
		if err != nil {
			return nil, fmt.Sprintf("SyncDingTalkUsers添加部门下用户失败：", err.Error())
		}
	}
	return nil, nil
}

//获取并处理钉钉部门下的用户信息入库
func (d DingTalkLogic) AddDeptUser(client *dingtalk.DingTalk, dept *model.Group, cursor int) error {
	// 处理部门下的人员信息
	deptId := strings.Split(dept.SourceDeptId, "_")
	tempId, err := strconv.Atoi(deptId[1])
	if err != nil {
		return err
	}
	//方式一：获取部门下用户信息，一次100个，遍历后插入数据库，经过验证，第三方依赖包有问题
	r := dingreq.DeptDetailUserInfo{}
	r.DeptId = tempId
	r.Language = "zh_CN"
	r.Cursor = cursor
	r.Size = 100
	//获取钉钉部门人员信息
	rep, err := client.GetDeptDetailUserInfo(&r)
	fmt.Println(fmt.Sprintf("当前获取的部门名称为：%s,总用户量为：%d", dept.GroupName, len(rep.DeptDetailUsers)))
	if err != nil {
		return errors.New(fmt.Sprintf("AddDeptUser获取钉钉部门人员信息失败：%s", err.Error()))
	}
	//方式二：临时处理方案：获取部门用户id列表，遍历，挨个从钉钉获取用户信息
	//dingr := dingreq.DeptUserId{}
	//dingr.DeptId = tempId
	//deptUserIds, err := client.GetDeptUserIds(&dingr)
	//if err != nil {
	//	return errors.New(fmt.Sprintf("AddDeptUser通过用户部门id从钉钉获取用户id列表失败:%s", err.Error()))
	//}
	// 遍历并处理当前部门下的人员信息
	for _, detail := range rep.DeptDetailUsers {
		//for index, userId := range deptUserIds.UserIds {
		//	fmt.Println(fmt.Sprintf("获取到的部门用户数为：%d,正在处理的用户序号为：%d,总Ids为：", len(deptUserIds.UserIds), index))
		//	fmt.Println(deptUserIds.UserIds)
		//	userReq := dingreq.UserDetail{}
		//	userReq.UserId = userId
		//	userReq.Language = "zh_CN"
		//	detail, err := client.GetUserDetail(&userReq)
		//	if err != nil {
		//		return errors.New(fmt.Sprintf("AddDeptUser通过用户id从钉钉获取用户详情失败:%s", err.Error()))
		//	}
		// 获取人员信息
		fmt.Println("钉钉人员详情：", detail)
		userName := ""
		if detail.OrgEmail != "" {
			emailstr := strings.Split(detail.OrgEmail, "@")
			userName = emailstr[0]
		}
		if userName == "" && detail.Name != "" {
			name := pinyin.LazyConvert(detail.Name, nil)
			userName = strings.Join(name, "")
		}
		if userName == "" && detail.Mobile != "" {
			userName = detail.Mobile
		}

		if detail.JobNumber == "" {
			detail.JobNumber = userName
		}
		//钉钉部门ids,转换为内部部门id
		sourceDeptIds := []string{}
		for _, deptId := range detail.DeptIds {
			sourceDeptIds = append(sourceDeptIds, fmt.Sprintf("%s_%d", config.Conf.DingTalk.DingTalkIdSource, deptId))
		}
		groupIds, err := isql.Group.DingTalkDeptIdsToGroupIds(sourceDeptIds)
		if err != nil {
			return errors.New(fmt.Sprintf("AddDeptUser转换钉钉部门id到本地分组id出错：%s", err.Error()))
		}
		user := request.DingUserAddReq{
			Username:      userName,
			Password:      config.Conf.DingTalk.DingTalkUserInitPassword,
			Nickname:      detail.Name,
			GivenName:     detail.Name,
			Mail:          detail.OrgEmail,
			JobNumber:     detail.JobNumber,
			Mobile:        detail.Mobile,
			Avatar:        detail.Avatar,
			PostalAddress: detail.WorkPlace,
			Departments:   dept.GroupName,
			Position:      detail.Title,
			Introduction:  detail.Remark,
			Status:        1,
			DepartmentId:  groupIds,
			Source:        config.Conf.DingTalk.DingTalkIdSource,
			SourceUserId:  fmt.Sprintf("%s_%s", config.Conf.DingTalk.DingTalkIdSource, detail.UserId),
			SourceUnionId: fmt.Sprintf("%s_%s", config.Conf.DingTalk.DingTalkIdSource, detail.UnionId),
		}
		// 入库
		repUser, err := d.AddUser(&user)
		if err != nil {
			return errors.New(fmt.Sprintf("AddDeptUser添加用户失败：%s", err.Error()))
		}
		fmt.Println("入库成功，用户信息为：")
		fmt.Println(repUser)
	}
	if rep.HasMore {
		err = d.AddDeptUser(client, dept, rep.NextCursor)
		if err != nil {
			return errors.New(fmt.Sprintf("AddDeptUser添加用户失败：%s", err.Error()))
		}
	}
	return nil
}

// AddGroup 添加部门数据
func (d DingTalkLogic) AddDept(r *request.DingGroupAddReq) (data *model.Group, rspError error) {
	// 判断部门名称是否存在
	filter := tools.H{"source_dept_id": r.SourceDeptId}
	dept := new(model.Group)
	err := isql.Group.Find(filter, dept)
	flag := errors.Is(err, gorm.ErrRecordNotFound)
	fmt.Println("部门是否存在：", filter, flag)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New(fmt.Sprintf("AddDept添加部门失败：%s", err.Error()))
	}
	//分组不存在，直接创建（此处通过部门名称和第三方id来共同判定唯一，理论上不会出现重复）
	if errors.Is(err, gorm.ErrRecordNotFound) {
		group := model.Group{
			GroupType:          r.GroupType,
			ParentId:           r.ParentId,
			GroupName:          r.GroupName,
			Remark:             r.Remark,
			Creator:            "system",
			Source:             r.Source,
			SourceDeptParentId: r.SourceDeptParentId,
			SourceDeptId:       r.SourceDeptId,
			SourceUserNum:      r.SourceUserNum,
		}
		pdn := ""
		if group.ParentId > 0 {
			pdn, err = isql.Group.GetGroupDn(r.ParentId, "")
			if err != nil {
				return nil, errors.New(fmt.Sprintf("AddDept获取父级部门dn失败：%s", err.Error()))
			}
		}
		err = ildap.Group.Add(&group, pdn)
		if err != nil {
			return nil, tools.NewLdapError(fmt.Errorf("AddDept向LDAP创建分组失败" + err.Error()))
		}
		// 创建
		err = isql.Group.Add(&group)
		if err != nil {
			return nil, tools.NewLdapError(fmt.Errorf("AddDept向MySQL创建分组失败:" + err.Error()))
		}
		// 默认创建分组之后，需要将admin添加到分组中
		adminInfo := new(model.User)
		err = isql.User.Find(tools.H{"id": 1}, adminInfo)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("AddDept获取admin用户失败：%s", tools.NewMySqlError(err).Error()))
		}

		err = isql.Group.AddUserToGroup(&group, []model.User{*adminInfo})
		if err != nil {
			return nil, tools.NewMySqlError(fmt.Errorf("AddDept添加用户到分组失败: %s", err.Error()))
		}
		return &group, nil
	} else { //分组存在
		//判断是否名字/备注/钉钉部门ID有修改
		if r.GroupName != dept.GroupName || r.Remark != dept.Remark || r.SourceDeptParentId != dept.SourceDeptParentId || r.SourceUserNum != dept.SourceUserNum {
			err = d.UpdateDept(r)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("AddDept更新部门失败：%s", err.Error()))
			}
		}
		//处理父级部门变化
		if r.SourceDeptParentId != dept.SourceDeptParentId {
			// TODO 待处理父级部门变化情况
		}
		return dept, nil
	}
}

// UpdateDept 更新部门数据
func (d DingTalkLogic) UpdateDept(r *request.DingGroupAddReq) error {
	oldData := new(model.Group)
	filter := tools.H{"source_dept_id": r.SourceDeptId}
	err := isql.Group.Find(filter, oldData)
	if err != nil {
		return errors.New(fmt.Sprintf("UpdateDept获取旧的部门信息失败:%s", tools.NewMySqlError(err).Error()))
	}
	dept := model.Group{
		Model:              oldData.Model,
		GroupName:          r.GroupName,
		Remark:             r.Remark,
		Creator:            "system",
		GroupType:          oldData.GroupType,
		SourceDeptId:       r.SourceDeptId,
		SourceDeptParentId: r.SourceDeptParentId,
		SourceUserNum:      r.SourceUserNum,
	}

	oldGroupName := oldData.GroupName
	oldRemark := oldData.Remark
	dn, err := isql.Group.GetGroupDn(oldData.ID, "")
	if err != nil {
		return errors.New(fmt.Sprintf("UpdateDept不去部门dn失败:%s", tools.NewMySqlError(err).Error()))
	}
	err = ildap.Group.Update(&dept, dn, oldGroupName, oldRemark)
	if err != nil {
		return tools.NewLdapError(fmt.Errorf("UpdateDept向LDAP更新分组失败：" + err.Error()))
	}
	//若配置了不允许修改分组名称，则不更新分组名称
	if !config.Conf.Ldap.LdapGroupNameModify {
		dept.GroupName = oldGroupName
	}
	err = isql.Group.Update(&dept)
	if err != nil {
		return tools.NewLdapError(fmt.Errorf("UpdateDept向MySQL更新分组失败:" + err.Error()))
	}
	return nil
}

// AddUser 添加用户数据
func (d DingTalkLogic) AddUser(r *request.DingUserAddReq) (data *model.User, rspError error) {
	// 兼容处理钉钉异常人员信息，若username，mail，mobile都没有的直接跳过
	if r.Username == "" && r.Mail == "" && r.Mobile == "" {
		emptyData := new(model.User)
		emptyData.Introduction = fmt.Sprintf("此用户：%s，username，mail，mobile皆为空，跳过入库，请手动置后台添加", r.Nickname)
		emptyData.Nickname = r.Nickname
		emptyData.SourceUserId = r.SourceUserId
		emptyData.Source = r.Source
		emptyData.GivenName = r.GivenName
		return emptyData, nil
	}

	isExist := false
	oldData := new(model.User)
	if isql.User.Exist(tools.H{"source_user_id": r.SourceUserId}) {
		err := isql.User.Find(tools.H{"source_user_id": r.SourceUserId}, oldData)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("AddUser根据钉钉用户id获取用户失败：%s", err.Error()))
		}
		isExist = true
	}
	if !isExist {
		if isql.User.Exist(tools.H{"source_union_id": r.SourceUnionId}) {
			err := isql.User.Find(tools.H{"source_union_id": r.SourceUnionId}, oldData)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("AddUser根据钉钉用户unionid获取用户失败：%s", err.Error()))
			}
			isExist = true
		}
	}
	//if !isExist {
	//	if r.Mail != "" && isql.User.Exist(tools.H{"mail": r.Mail}) {
	//		err := isql.User.Find(tools.H{"mail": r.Mail}, oldData)
	//		if err != nil {
	//			return nil, errors.New(fmt.Sprintf("AddUser根据钉钉用户mail获取用户失败：%s", err.Error()))
	//		}
	//		isExist = true
	//	}
	//}
	if !isExist {
		if isql.User.Exist(tools.H{"job_number": r.JobNumber}) {
			err := isql.User.Find(tools.H{"job_number": r.JobNumber}, oldData)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("AddUser根据钉钉用户job_number获取用户失败：%s", err.Error()))
			}
			isExist = true
		}
	}
	if !isExist {
		if isql.User.Exist(tools.H{"mobile": r.Mobile}) {
			err := isql.User.Find(tools.H{"mobile": r.Mobile}, oldData)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("AddUser根据钉钉用户mobile获取用户失败：%s", err.Error()))
			}
			isExist = true
		}
	}
	if !isExist {
		//组装用户名
		//先根据钉钉唯一id获取，查看数据库中是否存在
		//不存在，则根据用户名 like 用户名获取尾号最大的账号
		//重新设定用户名
		userData := new(model.User)
		err := isql.User.FindTheSameUserName(r.Username, userData)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		} else {
			// 找到重名用户，
			re := regexp.MustCompile("[0-9]+")
			num := re.FindString(userData.Username)
			n := 1
			if num != "" {
				m, err := strconv.Atoi(num)
				if err != nil {
					return
				}
				n = m + 1
			}
			r.Username = fmt.Sprintf("%s%d", r.Username, n)
		}
	}
	if isExist {
		r.Username = oldData.Username
		user, err := d.UpdateUser(r, oldData)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("AddUser用户已存在，更新用户失败：%s", err.Error()))
		}
		return user, nil
	}
	// 根据角色id获取角色
	r.RoleIds = []uint{2} // 默认添加为普通用户角色
	roles, err := isql.Role.GetRolesByIds(r.RoleIds)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("AddUser根据角色ID获取角色信息失败:%s", err.Error()))
	}

	var reqRoleSorts []int
	for _, role := range roles {
		reqRoleSorts = append(reqRoleSorts, int(role.Sort))
	}
	deptIds := tools.SliceToString(r.DepartmentId, ",")
	user := model.User{
		Username:      r.Username,
		Password:      r.Password,
		Nickname:      r.Nickname,
		GivenName:     r.GivenName,
		Mail:          r.Mail,
		JobNumber:     r.JobNumber,
		Mobile:        r.Mobile,
		Avatar:        r.Avatar,
		PostalAddress: r.PostalAddress,
		Departments:   r.Departments,
		Position:      r.Position,
		Introduction:  r.Introduction,
		Status:        r.Status,
		Creator:       "system",
		DepartmentId:  deptIds,
		Roles:         roles,
		Source:        r.Source,
		SourceUserId:  r.SourceUserId,
		SourceUnionId: r.SourceUnionId,
	}
	if user.Introduction == "" {
		user.Introduction = r.Nickname
	}
	if user.JobNumber == "" {
		user.JobNumber = r.Mobile
	}
	//先识别用户选择的部门是否是OU开头
	gdns := make(map[uint]string)
	for _, deptId := range r.DepartmentId {
		dn, err := isql.Group.GetGroupDn(deptId, "")
		if err != nil {
			return nil, errors.New(fmt.Sprintf("AddUser根据用户dn信息失败:%s", err.Error()))
		}
		gdn := fmt.Sprintf("%s,%s", dn, config.Conf.Ldap.LdapBaseDN)
		if gdn[:3] == "ou=" {
			return nil, errors.New(fmt.Sprintf("AddUser不能添加用户到OU组织单元:%s", gdn))
		}
		gdns[deptId] = gdn
	}
	//先创建用户到默认分组
	err = ildap.User.Add(&user)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("AddUser向LDAP创建用户失败：" + err.Error()))
	}
	isExistUser := false
	for deptId, gdn := range gdns {
		//根据选择的部门，添加到部门内
		err = ildap.Group.AddUserToGroup(gdn, fmt.Sprintf("uid=%s,%s", user.Username, config.Conf.Ldap.LdapUserDN))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("AddUser向部门添加用户失败：%s", err.Error()))
		}
		if !isExistUser {
			err = isql.User.Add(&user)
			if err != nil {
				return nil, tools.NewMySqlError(fmt.Errorf("向MySQL创建用户失败：" + err.Error()))
			}
			isExistUser = true
		}
		//根据部门分配，将用户和部门信息维护到部门关系表里面
		users := []model.User{}
		users = append(users, user)
		depart := new(model.Group)
		filter := tools.H{"id": deptId}
		err = isql.Group.Find(filter, depart)
		if err != nil {
			return nil, tools.NewMySqlError(err)
		}
		err = isql.Group.AddUserToGroup(depart, users)
		if err != nil {
			return nil, tools.NewMySqlError(fmt.Errorf("AddUser向MySQL添加用户到分组关系失败：" + err.Error()))
		}
	}
	return &user, nil
}

// Update 更新数据
func (d DingTalkLogic) UpdateUser(r *request.DingUserAddReq, oldData *model.User) (data *model.User, rspError error) {
	deptIds := tools.SliceToString(r.DepartmentId, ",")
	user := model.User{
		Model:         oldData.Model,
		Username:      r.Username,
		Nickname:      r.Nickname,
		GivenName:     r.GivenName,
		Mail:          r.Mail,
		JobNumber:     r.JobNumber,
		Mobile:        r.Mobile,
		Avatar:        r.Avatar,
		PostalAddress: r.PostalAddress,
		Departments:   r.Departments,
		Position:      r.Position,
		Introduction:  r.Introduction,
		Creator:       "system",
		DepartmentId:  deptIds,
		Source:        oldData.Source,
		Roles:         oldData.Roles,
		SourceUserId:  r.SourceUserId,
		SourceUnionId: r.SourceUnionId,
	}
	if user.Introduction == "" {
		user.Introduction = r.Nickname
	}
	if user.PostalAddress == "" {
		user.PostalAddress = "没有填写地址"
	}
	if user.Position == "" {
		user.Position = "技术"
	}
	if user.JobNumber == "" {
		user.JobNumber = r.Mobile
	}
	err := ildap.User.Update(oldData.Username, &user)
	if err != nil {
		return nil, tools.NewLdapError(fmt.Errorf("UpdateUser在LDAP更新用户失败：" + err.Error()))
	}

	// 更新用户
	if !config.Conf.Ldap.LdapUserNameModify {
		user.Username = oldData.Username
	}
	err = isql.User.Update(&user)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("UpdateUser在MySQL更新用户失败：" + err.Error()))
	}
	//判断部门信息是否有变化有变化则更新相应的数据库
	oldDeptIds := tools.StringToSlice(oldData.DepartmentId, ",")
	addDeptIds, removeDeptIds := tools.ArrUintCmp(oldDeptIds, r.DepartmentId)
	for _, deptId := range removeDeptIds {
		//从旧组中删除
		err = User.RemoveUserToGroup(deptId, []uint{oldData.ID})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("UpdateUser将用户从分组移除失败：%s", err.Error()))
		}
	}
	for _, deptId := range addDeptIds {
		//添加到新分组中
		err = User.AddUserToGroup(deptId, []uint{oldData.ID})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("UpdateUser将用户添加至分组失败：%s", err.Error()))
		}
	}
	return &user, nil
}
