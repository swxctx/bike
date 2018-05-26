// generate by ginpt
package accountcontroller

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
		CacheAPIKey: "account:login",
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

// UpdatePasswordArgs 修改密码请求参数
type UpdatePasswordArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	Password string `form:"password" json:"password"`
}

// DoUpdatePassword 修改密码
func DoUpdatePassword(c *gin.Context) {
	params := UpdatePasswordArgs{
		CacheAPIKey: "account:update_password",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
	}

	doUpdatePassword(c, &params)

}

// AddUserArgs 添加用户请求参数
type AddUserArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	Username string `form:"username" json:"username"`
	Name     string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	IsAdmin  int32  `form:"is_admin" json:"is_admin"`
}

// DoAddUser 添加用户
func DoAddUser(c *gin.Context) {
	params := AddUserArgs{
		CacheAPIKey: "account:add_user",
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

		if utils.IsEmpty(params.Name) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "name is required",
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

		if utils.IsEmpty(params.IsAdmin) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "is_admin is required",
			})
			return
		}
	}

	doAddUser(c, &params)

}
