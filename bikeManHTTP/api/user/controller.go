// generate by ginpt
package usercontroller

import (
	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/domego/ginkits/params"
	"github.com/domego/gokits/log"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/model/bike"
)

var _ = kits.RenderSuccess

type BaseParam paramkits.BaseParam

// doLogin 登录
func doLogin(c *gin.Context, params *LoginArgs) (result interface{}, hasError bool) {
	user := bike.GetUserFirst("`deleted`=0")
	if user != nil {
		log.Infof("user:%v", user)
	} else {
		return &errorkits.ErrorMessage{
			Code:    100101,
			Content: "用户名已经被注册了",
		}, true
	}
	return
}
