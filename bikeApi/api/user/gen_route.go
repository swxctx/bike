// generate by ginpt
package usercontroller

import (
	"github.com/domego/ginkits/middleware"
	"github.com/gin-gonic/gin"
)

var UserRouteGroup *gin.RouterGroup

func RegistRoute(route *gin.Engine) {
	UserRouteGroup = route.Group("/bike/user")
	UserRouteGroup.Use(middleware.CommonHandler())
	{
		// 注册
		UserRouteGroup.GET("/v1/register", DoRegister)
		UserRouteGroup.POST("/v1/register", DoRegister)
		// 登录
		UserRouteGroup.GET("/v1/login", DoLogin)
		UserRouteGroup.POST("/v1/login", DoLogin)
	}
}
