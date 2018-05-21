// generate by ginpt
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/bikeHTTP/api/user"
)

func registRoute(route *gin.Engine) {

	// 用户相关
	usercontroller.RegistRoute(route)
}
