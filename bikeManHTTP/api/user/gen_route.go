// generate by ginpt
package usercontroller

import (
	"github.com/domego/ginkits/middleware"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/middleware"
)

var UserRouteGroup *gin.RouterGroup

func RegistRoute(route *gin.Engine) {
	UserRouteGroup = route.Group("/bike_mp/user")
	UserRouteGroup.Use(middleware.CommonHandler())
	{
		// 用户列表
		UserRouteGroup.GET("/v1/user_list", DoUserList)
		UserRouteGroup.POST("/v1/user_list", DoUserList)
	}
}
