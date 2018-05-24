// generate by ginpt
package usecontroller

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

// HistoryLogArgs 使用记录请求参数
type HistoryLogArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string
}

// DoHistoryLog 使用记录
func DoHistoryLog(c *gin.Context) {
	params := HistoryLogArgs{
		CacheAPIKey: "use:history_log",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)

	result, hasError := doHistoryLog(c, &params)
	if hasError {
		kits.RenderError(c, result)
		return
	}
	kits.RenderSuccess(c, result)
}

// StartUseArgs 开始使用请求参数
type StartUseArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	BikeId   int32  `form:"bike_id" json:"bike_id"`
	PointLng string `form:"point_lng" json:"point_lng"`
	PointLat string `form:"point_lat" json:"point_lat"`
}

// DoStartUse 开始使用
func DoStartUse(c *gin.Context) {
	params := StartUseArgs{
		CacheAPIKey: "use:start_use",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
		if utils.IsEmpty(params.BikeId) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "bike_id is required",
			})
			return
		}

		if utils.IsEmpty(params.PointLng) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "point_lng is required",
			})
			return
		}

		if utils.IsEmpty(params.PointLat) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "point_lat is required",
			})
			return
		}
	}

	result, hasError := doStartUse(c, &params)
	if hasError {
		kits.RenderError(c, result)
		return
	}
	kits.RenderSuccess(c, result)
}
