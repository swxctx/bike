// generate by ginpt
package usercontroller

import (
	"github.com/domego/ginkits/middleware"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/middleware"
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
		// 获取资料
		UserRouteGroup.GET("/v1/get_profile", bikemiddleware.AccessTokenTokenHandler(), DoGetProfile)
		UserRouteGroup.POST("/v1/get_profile", bikemiddleware.AccessTokenTokenHandler(), DoGetProfile)
		// 实名认证
		UserRouteGroup.GET("/v1/auth_card", bikemiddleware.AccessTokenTokenHandler(), DoAuthCard)
		UserRouteGroup.POST("/v1/auth_card", bikemiddleware.AccessTokenTokenHandler(), DoAuthCard)
		// 更新资料
		UserRouteGroup.GET("/v1/update_profile", bikemiddleware.AccessTokenTokenHandler(), DoUpdateProfile)
		UserRouteGroup.POST("/v1/update_profile", bikemiddleware.AccessTokenTokenHandler(), DoUpdateProfile)
		// 修改密码
		UserRouteGroup.GET("/v1/update_password", bikemiddleware.AccessTokenTokenHandler(), DoUpdatePassword)
		UserRouteGroup.POST("/v1/update_password", bikemiddleware.AccessTokenTokenHandler(), DoUpdatePassword)
	}
}
