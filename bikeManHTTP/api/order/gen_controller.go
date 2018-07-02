// generate by ginpt
package ordercontroller

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

// OrderListArgs 待处理订单列表请求参数
type OrderListArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string
}

// DoOrderList 待处理订单列表
func DoOrderList(c *gin.Context) {
	params := OrderListArgs{
		CacheAPIKey: "order:order_list",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)

	doOrderList(c, &params)

}

// ChangeArgs 订单状态管理请求参数
type ChangeArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	Id int32 `form:"id" json:"id"`
}

// DoChange 订单状态管理
func DoChange(c *gin.Context) {
	params := ChangeArgs{
		CacheAPIKey: "order:change",
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

	doChange(c, &params)

}
