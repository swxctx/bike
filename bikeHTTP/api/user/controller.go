// generate by ginpt
package usercontroller

import (
	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/domego/ginkits/params"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/middleware"
	"github.com/swxctx/bike/model/bike"
)

var _ = kits.RenderSuccess

type BaseParam paramkits.BaseParam

// doRegister 注册
func doRegister(c *gin.Context, params *RegisterArgs) (result interface{}, hasError bool) {
	// 1.查询该用户是否注册过
	user := bike.GetUserFirst("`deleted`=0 AND `username`=?", params.Username)
	if user != nil {
		return &errorkits.ErrorMessage{
			Code:    100101,
			Content: "用户名已经被注册了",
		}, true
	}
	// 2.校验手机号
	if !middleware.CheckPhone(params.Phone) {
		return &errorkits.ErrorMessage{
			Code:    100102,
			Content: "手机号不合法",
		}, true
	}
	// 3.校验用户名
	if len(params.Username) > 12 {
		return &errorkits.ErrorMessage{
			Code:    100103,
			Content: "用户名长度超出限制",
		}, true
	}
	// 4.校验密码
	if len(params.Password) > 16 || len(params.Password) < 6 {
		return &errorkits.ErrorMessage{
			Code:    100104,
			Content: "密码长度为6-16",
		}, true
	}
	user.Username = params.Username
	user.Phone = params.Phone
	user.Password = middleware.GenPassword(params.Password)
	meta := bike.NewUserMeta()
	meta.Uid = user.Id
	meta.Bio = params.Bio
	meta.Sex = params.Sex
	user.Insert()
	meta.Insert()

	loginResult := convertLoginResult(user)
	return loginResult, false
}

// doLogin 登录
func doLogin(c *gin.Context, params *LoginArgs) (result interface{}, hasError bool) {
	return result, false
}
