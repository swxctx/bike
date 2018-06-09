// generate by ginpt
package accountcontroller

import (
	"time"

	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/domego/ginkits/params"

	"github.com/gin-gonic/gin"
	"github.com/swxctx/bike/middleware"
	"github.com/swxctx/bike/model/bike"
)

var _ = kits.RenderSuccess

type BaseParam paramkits.BaseParam

// doLogin 登录
func doLogin(c *gin.Context, params *LoginArgs) {
	pass := bikemiddleware.GenPassword(params.Password)
	user := bike.GetAdminFirst("`deleted`=0 AND `username`=? AND `password`=?", params.Username, pass)
	if user == nil {
		kits.RenderError(c, &errorkits.ErrorMessage{
			Code:    100101,
			Content: "用户名或密码不正确",
		})
		return
	}
	user.LastLoginTs = time.Now().Unix()
	kits.RenderSuccess(c, gin.H{
		"username": user.Username,
		"name":     user.Name,
		"is_admin": user.IsAdmin,
	})
}

// doUpdatePassword 修改密码
func doUpdatePassword(c *gin.Context, params *UpdatePasswordArgs) {
	admin := bike.GetAdminFirst("`deleted`=0 AND `username`=?", params.Username)
	if admin == nil {
		kits.RenderError(c, &errorkits.ErrorMessage{
			Code:    100102,
			Content: "网络不给力，请稍后再试",
		})
		return
	}
	pass := bikemiddleware.GenPassword(params.OldPassword)
	if admin.Password != pass {
		kits.RenderError(c, &errorkits.ErrorMessage{
			Code:    100103,
			Content: "旧密码输入错误，请确认",
		})
		return
	}
	admin.Password = bikemiddleware.GenPassword(params.Password)
	admin.Update()
}

// doAddUser 添加用户
func doAddUser(c *gin.Context, params *AddUserArgs) {
	admin := bike.GetAdminFirst("`deleted`=0 AND `username`=?", params.Username)
	if admin != nil {
		kits.RenderError(c, &errorkits.ErrorMessage{
			Code:    100104,
			Content: "用户已经存在了",
		})
		return
	}
	admin = bike.NewAdmin()
	admin.Username = params.Username
	admin.Name = params.Name
	admin.IsAdmin = params.IsAdmin
	pass := bikemiddleware.GenPassword(params.Password)
	admin.Password = pass
	admin.Insert()
}
