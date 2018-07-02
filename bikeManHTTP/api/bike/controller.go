// generate by ginpt
package bikecontroller

import (
	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/domego/ginkits/params"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/model/bike"
)

var _ = kits.RenderSuccess

type BaseParam paramkits.BaseParam

// doBikeList 单车列表
func doBikeList(c *gin.Context, params *BikeListArgs) {
	bikes := bike.GetBikeWhere("`deleted`=0")
	kits.RenderSuccess(c, gin.H{
		"list": bikes,
	})
}

// doChange 单车状态管理
func doChange(c *gin.Context, params *ChangeArgs) {
	bikeInfo := bike.GetBikeFirst("`deleted`=0 AND `by_id`=?", params.BikeId)
	if bikeInfo == nil {
		kits.RenderError(c, &errorkits.ErrorMessage{
			Code:    100201,
			Content: "网络不给力，请稍后再试",
		})
		return
	}
	bikeInfo.Status = params.Status
	bikeInfo.Update()
	return
}

// doAdd 新增单车
func doAdd(c *gin.Context, params *AddArgs) {
	bikeInfo := bike.NewBike()
	bikeInfo.Ip = params.Ip
	bikeInfo.Port = params.Port
	bikeInfo.Status = 0
	info := bike.GetBikeFirst("`deleted`=0 ORDER BY `created_at` DESC")
	if info != nil {
		if info.ById != 0 {
			bikeInfo.ById = info.ById + 1
		} else {
			bikeInfo.ById = 1001
		}
	}
	bikeInfo.Insert()
	return
}

// doAlarmList 告警信息
func doAlarmList(c *gin.Context, params *AlarmListArgs) {
	list := &AlarmList{}
	bikes := bike.GetBikeAlarmWhere("`deleted`=0")
	for _, b := range bikes {
		list.List = append(list.List, convertAlarmInfo(b))
	}
	kits.RenderSuccess(c, gin.H{
		"list": list,
	})
}
