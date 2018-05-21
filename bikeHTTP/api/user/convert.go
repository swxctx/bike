package usercontroller

import (
	"github.com/swxctx/bike/middleware"
	"github.com/swxctx/bike/model/bike"
)

func convertLoginResult(user *bike.User) *LoginResult {
	result := &LoginResult{
		Username: user.Username,
		Phone:    user.Phone,
		Email:    user.Email,
	}
	result.AccessToken = middleware.GenAccessToken(user.Id)
	return result
}
