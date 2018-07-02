// generate by ginpt
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/bikeApi/api/bag"
	"github.com/swxctx/bike/bikeApi/api/feedback"
	"github.com/swxctx/bike/bikeApi/api/use"
	"github.com/swxctx/bike/bikeApi/api/user"
)

func registRoute(route *gin.Engine) {

	// 用户相关
	usercontroller.RegistRoute(route)
	// 使用相关
	usecontroller.RegistRoute(route)
	// 钱包相关
	bagcontroller.RegistRoute(route)
	// 反馈
	feedbackcontroller.RegistRoute(route)
}
