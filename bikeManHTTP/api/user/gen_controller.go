// generate by ginpt
package usercontroller

import (
	"fmt"

	"github.com/domego/ginkits"
	"github.com/domego/ginkits/params"
	"github.com/domego/gokits"
	"github.com/domego/gokits/log"
	"github.com/gin-gonic/gin"
)

var _ = log.Tracef
var _ = kits.RenderSuccess
var _ = utils.IsEmpty
var _ = fmt.Sprintf

// LoginArgs 登录请求参数
type LoginArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// DoLogin 登录
func DoLogin(c *gin.Context) {
	params := LoginArgs{
		CacheAPIKey: "user:login",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
		if utils.IsEmpty(params.Username) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "username is required",
			})
			return
		}

		if utils.IsEmpty(params.Password) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "password is required",
			})
			return
		}
	}

	doLogin(c, &params)

}
