// generate by ginpt
package usecontroller

import (
	"github.com/domego/ginkits/middleware"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/middleware"
)

var UseRouteGroup *gin.RouterGroup

func RegistRoute(route *gin.Engine) {
	UseRouteGroup = route.Group("/bike/use")
	UseRouteGroup.Use(middleware.CommonHandler())
	{
		// 使用记录
		UseRouteGroup.GET("/v1/history_log", bikemiddleware.AccessTokenTokenHandler(), DoHistoryLog)
		UseRouteGroup.POST("/v1/history_log", bikemiddleware.AccessTokenTokenHandler(), DoHistoryLog)
		// 开始使用
		UseRouteGroup.GET("/v1/start_use", bikemiddleware.AccessTokenTokenHandler(), DoStartUse)
		UseRouteGroup.POST("/v1/start_use", bikemiddleware.AccessTokenTokenHandler(), DoStartUse)
	}
}
