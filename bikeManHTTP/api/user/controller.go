// generate by ginpt
package usercontroller

import (
	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/domego/ginkits/params"

	"github.com/gin-gonic/gin"
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
			Message: "无效的status值",
			Content: "用户名或密码不正确",
		})
		return
	}
	kits.RenderSuccess(c, gin.H{
		"username": user.Username,
		"name":     user.Name,
		"is_admin": user.IsAdmin,
	})
}
