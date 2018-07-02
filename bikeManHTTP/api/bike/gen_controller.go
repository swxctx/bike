// generate by ginpt
package bikecontroller

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

// BikeListArgs 单车列表请求参数
type BikeListArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string
}

// DoBikeList 单车列表
func DoBikeList(c *gin.Context) {
	params := BikeListArgs{
		CacheAPIKey: "bike:bike_list",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)

	doBikeList(c, &params)

}

// ChangeArgs 单车状态管理请求参数
type ChangeArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	BikeId int32 `form:"bike_id" json:"bike_id"`
	Status int32 `form:"status" json:"status"`
}

// DoChange 单车状态管理
func DoChange(c *gin.Context) {
	params := ChangeArgs{
		CacheAPIKey: "bike:change",
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

	}

	doChange(c, &params)

}

// AddArgs 新增单车请求参数
type AddArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	Ip   string `form:"ip" json:"ip"`
	Port string `form:"port" json:"port"`
}

// DoAdd 新增单车
func DoAdd(c *gin.Context) {
	params := AddArgs{
		CacheAPIKey: "bike:add",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
		if utils.IsEmpty(params.Ip) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "ip is required",
			})
			return
		}

	}

	doAdd(c, &params)

}

// AlarmListArgs 告警信息请求参数
type AlarmListArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string
}

// DoAlarmList 告警信息
func DoAlarmList(c *gin.Context) {
	params := AlarmListArgs{
		CacheAPIKey: "bike:alarm_list",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)

	doAlarmList(c, &params)

}
