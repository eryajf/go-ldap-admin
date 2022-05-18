package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"

	"github.com/eryajf-world/go-ldap-admin/config"
	"github.com/eryajf-world/go-ldap-admin/public/tools"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zht "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Api          = &ApiController{}
	Group        = &GroupController{}
	Menu         = &MenuController{}
	Role         = &RoleController{}
	User         = &UserController{}
	OperationLog = &OperationLogController{}
	Base         = &BaseController{}

	validate = validator.New()
	trans    ut.Translator
)

func init() {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	_ = zht.RegisterDefaultTranslations(validate, trans)
	_ = validate.RegisterValidation("checkMobile", checkMobile)
}

func checkMobile(fl validator.FieldLevel) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(fl.Field().String())
}

func Run(c *gin.Context, req interface{}, fn func() (interface{}, interface{})) {
	var err error
	// bind struct
	err = c.Bind(req)
	if err != nil {
		tools.Err(c, tools.NewValidatorError(err), nil)
		return
	}
	// 校验
	err = validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			tools.Err(c, tools.NewValidatorError(fmt.Errorf(err.Translate(trans))), nil)
			return
		}
	}
	data, err1 := fn()
	if err1 != nil {
		tools.Err(c, tools.ReloadErr(err1), data)
		return
	}
	tools.Success(c, data)
}

func Demo(c *gin.Context) {
	CodeDebug()
	c.JSON(http.StatusOK, tools.H{"code": 200, "msg": "ok", "data": "pong"})
}

func CodeDebug() {
	// pass := "SI8HqZxBLGDTU5ZsL8fQyyFLKYSF3bI1KIMZ9yBqo6xWFQDk0HH7AvZUFqbiWNSPNWWNjS9TNsgS1ubTg5Lh7bV+AeSuW3cEuLN9wJI/9tg50eS94O3NETWf3RoZ2jBrd/huwcDRrNk5+cqLffUXI5Da68i1QEiQ3X1w/DW6VH4="
	// fmt.Printf("秘钥为：%s\n", config.Conf.System.RSAPrivateBytes)
	// // 密码通过RSA解密
	// decodeData, err := tools.RSADecrypt([]byte(pass), config.Conf.System.RSAPrivateBytes)
	// if err != nil {
	// 	fmt.Printf("密码解密失败：%s\n", err)
	// }
	// fmt.Printf("密码解密后为：%s\n", string(decodeData))
	// users, err := isql.User.GetUserByIds([]uint{1, 2, 3})
	// if err != nil {
	// 	fmt.Printf("获取用户失败：%s\n", err)
	// }
	// for _, user := range users {
	// 	fmt.Println("===============", user.Username)
	// }

	// user, _ := isql.User.GetUserById(1)
	// fmt.Println(user)
	// user1 := new(model.User)
	// _ = isql.User.Find(tools.H{"id": 1}, user1)
	// fmt.Println(user1)

	// user2, _ := isql.User.GetCurrentUser(c)
	// fmt.Println("========", user2)
	data := []byte("hello world")
	m, _ := tools.RSAEncrypt(data, config.Conf.System.RSAPublicBytes)
	fmt.Println(base64.StdEncoding.EncodeToString(m))
	s, _ := tools.RSADecrypt(m, config.Conf.System.RSAPrivateBytes)
	fmt.Println(string(s))
	new := tools.NewGenPasswd("hello world")
	fmt.Println("==========", new)
	s1 := tools.NewParPasswd(new)
	fmt.Println(string(s1))
}
