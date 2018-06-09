// generate by ginpt
package accountcontroller

import (
	"github.com/domego/ginkits/middleware"
	"github.com/gin-gonic/gin"
)

var AccountRouteGroup *gin.RouterGroup

func RegistRoute(route *gin.Engine) {
	AccountRouteGroup = route.Group("/bike_mp/account")
	AccountRouteGroup.Use(middleware.CommonHandler())
	{
		// 登录
		AccountRouteGroup.GET("/v1/login", DoLogin)
		AccountRouteGroup.POST("/v1/login", DoLogin)
		// 修改密码
		AccountRouteGroup.GET("/v1/update_password", DoUpdatePassword)
		AccountRouteGroup.POST("/v1/update_password", DoUpdatePassword)
		// 添加用户
		AccountRouteGroup.GET("/v1/update_user", DoAddUser)
		AccountRouteGroup.POST("/v1/update_user", DoAddUser)
	}
}
