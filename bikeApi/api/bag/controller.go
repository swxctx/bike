// generate by ginpt
package bagcontroller

import (
	"github.com/domego/ginkits"
	"github.com/gin-gonic/gin"
	"github.com/swxctx/bike/middleware"
)

var _ = kits.RenderSuccess

type BaseParam bikemiddleware.BaseParam

// doDetail 余额明细
func doDetail(c *gin.Context, params *DetailArgs) (result interface{}, hasError bool) {
	return
}

// doTopUp 充值
func doTopUp(c *gin.Context, params *TopUpArgs) (result interface{}, hasError bool) {
	return
}
