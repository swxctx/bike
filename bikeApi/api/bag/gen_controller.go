// generate by ginpt
package bagcontroller

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

// DetailArgs 余额明细请求参数
type DetailArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string
}

// DoDetail 余额明细
func DoDetail(c *gin.Context) {
	params := DetailArgs{
		CacheAPIKey: "bag:detail",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)

	result, hasError := doDetail(c, &params)
	if hasError {
		kits.RenderError(c, result)
		return
	}
	kits.RenderSuccess(c, result)
}

// TopUpArgs 充值请求参数
type TopUpArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	Phone  string `form:"phone" json:"phone"`
	Amount int32  `form:"amount" json:"amount"`
}

// DoTopUp 充值
func DoTopUp(c *gin.Context) {
	params := TopUpArgs{
		CacheAPIKey: "bag:top_up",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
		if utils.IsEmpty(params.Phone) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "phone is required",
			})
			return
		}

		if utils.IsEmpty(params.Amount) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "amount is required",
			})
			return
		}
	}

	result, hasError := doTopUp(c, &params)
	if hasError {
		kits.RenderError(c, result)
		return
	}
	kits.RenderSuccess(c, result)
}
