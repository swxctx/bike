// generate by ginpt
package ordercontroller

import (
	"github.com/domego/ginkits/middleware"
	"github.com/gin-gonic/gin"
)

var OrderRouteGroup *gin.RouterGroup

func RegistRoute(route *gin.Engine) {
	OrderRouteGroup = route.Group("/bike_mp/order")
	OrderRouteGroup.Use(middleware.CommonHandler())
	{
		// 待处理订单列表
		OrderRouteGroup.GET("/v1/get_list", DoOrderList)
		OrderRouteGroup.POST("/v1/get_list", DoOrderList)
		// 订单状态管理
		OrderRouteGroup.GET("/v1/change", DoChange)
		OrderRouteGroup.POST("/v1/change", DoChange)
	}
}
