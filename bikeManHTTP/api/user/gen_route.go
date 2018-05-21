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
		// 登录
		UserRouteGroup.GET("/v1/login", DoLogin)
		UserRouteGroup.POST("/v1/login", DoLogin)
	}
}
