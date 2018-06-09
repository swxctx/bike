// generate by ginpt
package feedbackcontroller

import (
	"github.com/domego/ginkits/middleware"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/middleware"
)

var FeedBackRouteGroup *gin.RouterGroup

func RegistRoute(route *gin.Engine) {
	FeedBackRouteGroup = route.Group("/bike/feedback")
	FeedBackRouteGroup.Use(middleware.CommonHandler())
	{
		// 提交反馈
		FeedBackRouteGroup.GET("/v1/report", bikemiddleware.AccessTokenTokenHandler(), Doreport)
		FeedBackRouteGroup.POST("/v1/report", bikemiddleware.AccessTokenTokenHandler(), Doreport)
	}
}
