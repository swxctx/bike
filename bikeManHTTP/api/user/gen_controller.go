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

// UserListArgs 用户列表请求参数
type UserListArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string
}

// DoUserList 用户列表
func DoUserList(c *gin.Context) {
	params := UserListArgs{
		CacheAPIKey: "user:user_list",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)

	doUserList(c, &params)

}

// ForbidArgs 封禁用户请求参数
type ForbidArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	Id     int32 `form:"id" json:"id"`
	Status int32 `form:"status" json:"status"`
}

// DoForbid 封禁用户
func DoForbid(c *gin.Context) {
	params := ForbidArgs{
		CacheAPIKey: "user:forbid",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
		if utils.IsEmpty(params.Id) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "id is required",
			})
			return
		}

	}

	doForbid(c, &params)

}
