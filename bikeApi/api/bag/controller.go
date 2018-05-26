// generate by ginpt
package bagcontroller

import (
	"time"

	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/middleware"
	"github.com/swxctx/bike/model/bike"
)

var _ = kits.RenderSuccess

type BaseParam bikemiddleware.BaseParam

// doDetail 余额明细
func doDetail(c *gin.Context, params *DetailArgs) (result interface{}, hasError bool) {
	balanceDetail := &BalanceDetail{}
	bag := bike.GetUserBagFirst("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	if bag != nil {
		balanceDetail.Balance = bag.Balance
	}
	topLog := bike.GetTopLogWhere("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	for _, top := range topLog {
		if top != nil {
			balanceDetail.List = append(balanceDetail.List, convertTopLog(top))
		}
	}
	return balanceDetail, false
}

// doTopUp 充值
func doTopUp(c *gin.Context, params *TopUpArgs) (result interface{}, hasError bool) {
	if !bikemiddleware.CheckPhone(params.Phone) {
		return &errorkits.ErrorMessage{
			Code:    300101,
			Content: "手机号不合法",
		}, true
	}
	if params.Amount == 0 {
		return &errorkits.ErrorMessage{
			Code:    300102,
			Content: "输入金额不能为空",
		}, true
	}
	if params.BaseParam.Uid == 0 {
		return &errorkits.ErrorMessage{
			Code:    300103,
			Content: "请先登录",
		}, true
	}
	topLog := bike.NewTopLog()
	topLog.Uid = params.BaseParam.Uid
	topLog.Count = params.Amount
	topLog.TopTs = time.Now().Unix()
	topLog.Insert()
	return
}
