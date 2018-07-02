// generate by ginpt
package ordercontroller

import (
	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/domego/ginkits/params"
	"github.com/gin-gonic/gin"
	"github.com/swxctx/bike/model/bike"
)

var _ = kits.RenderSuccess

type BaseParam paramkits.BaseParam

// doOrderList 待处理订单列表
func doOrderList(c *gin.Context, params *OrderListArgs) {
	orders := bike.GetTopLogWhere("`deleted`=0 AND `status`=0")
	kits.RenderSuccess(c, gin.H{
		"list": orders,
	})
}

// doChange 订单状态管理
func doChange(c *gin.Context, params *ChangeArgs) {
	order := bike.GetTopLogFirst("`deleted`=0 AND `status`=0")
	if order == nil {
		kits.RenderError(c, &errorkits.ErrorMessage{
			Code:    100301,
			Content: "该订单已经处理过了",
		})
	}
	order.Status = 1
	order.Update()
	userBag := bike.GetUserBagFirst("`deleted`=0 AND `uid`=?", order.Uid)
	userBag.Balance += order.Count
	userBag.Update()
}
