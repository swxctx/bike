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
	result.AccessToken = bikemiddleware.GenAccessToken(user.Id)
	return result
}

func convertUserProfile(user *bike.User, meta *bike.UserMeta) *UserProfile {
	result := &UserProfile{
		Username: user.Username,
		Phone:    user.Phone,
		Email:    user.Email,
		Sex:      meta.Sex,
		Avatar:   meta.Avatar,
		Bio:      meta.Bio,
		CardName: meta.CardName,
		CardId:   meta.CardId,
	}
	return result
}
