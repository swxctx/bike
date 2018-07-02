// generate by ginpt
package usercontroller

import (
	"github.com/domego/ginkits/middleware"
	"github.com/gin-gonic/gin"
)

var UserRouteGroup *gin.RouterGroup

func RegistRoute(route *gin.Engine) {
	UserRouteGroup = route.Group("/bike_mp/user")
	UserRouteGroup.Use(middleware.CommonHandler())
	{
		// 用户列表
		UserRouteGroup.GET("/v1/get_list", DoUserList)
		UserRouteGroup.POST("/v1/get_list", DoUserList)
		// 封禁用户
		UserRouteGroup.GET("/v1/forbid", DoForbid)
		UserRouteGroup.POST("/v1/forbid", DoForbid)
	}
}
