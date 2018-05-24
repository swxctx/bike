// generate by ginpt
package bagcontroller

import (
	"github.com/domego/ginkits/middleware"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/middleware"
)

var BagRouteGroup *gin.RouterGroup

func RegistRoute(route *gin.Engine) {
	BagRouteGroup = route.Group("/bike/bag")
	BagRouteGroup.Use(middleware.CommonHandler())
	{
		// 余额明细
		BagRouteGroup.GET("/v1/detail", bikemiddleware.AccessTokenTokenHandler(), DoDetail)
		BagRouteGroup.POST("/v1/detail", bikemiddleware.AccessTokenTokenHandler(), DoDetail)
		// 充值
		BagRouteGroup.GET("/v1/top_up", bikemiddleware.AccessTokenTokenHandler(), DoTopUp)
		BagRouteGroup.POST("/v1/top_up", bikemiddleware.AccessTokenTokenHandler(), DoTopUp)
	}
}
