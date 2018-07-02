// generate by ginpt
package bikecontroller

import (
	"github.com/domego/ginkits/middleware"
	"github.com/gin-gonic/gin"
)

var BikeRouteGroup *gin.RouterGroup

func RegistRoute(route *gin.Engine) {
	BikeRouteGroup = route.Group("/bike_mp/bike")
	BikeRouteGroup.Use(middleware.CommonHandler())
	{
		// 单车列表
		BikeRouteGroup.GET("/v1/get_list", DoBikeList)
		BikeRouteGroup.POST("/v1/get_list", DoBikeList)
		// 单车状态管理
		BikeRouteGroup.GET("/v1/change", DoChange)
		BikeRouteGroup.POST("/v1/change", DoChange)
		// 新增单车
		BikeRouteGroup.GET("/v1/add", DoAdd)
		BikeRouteGroup.POST("/v1/add", DoAdd)
		// 告警信息
		BikeRouteGroup.GET("/v1/alarm_list", DoAlarmList)
		BikeRouteGroup.POST("/v1/alarm_list", DoAlarmList)
	}
}
