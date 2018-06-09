// generate by ginpt
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/bikeManHTTP/api/account"
	"github.com/swxctx/bike/bikeManHTTP/api/bike"
	"github.com/swxctx/bike/bikeManHTTP/api/order"
	"github.com/swxctx/bike/bikeManHTTP/api/user"
)

func registRoute(route *gin.Engine) {

	// 用户相关
	accountcontroller.RegistRoute(route)
	// 用户相关
	usercontroller.RegistRoute(route)
	// 单车相关
	bikecontroller.RegistRoute(route)
	// 订单相关
	ordercontroller.RegistRoute(route)
}
