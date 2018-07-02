package usercontroller

import "github.com/swxctx/bike/model/bike"

func convertUserProfile(user *bike.User, meta *bike.UserMeta) *UserProfile {
	result := &UserProfile{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Email:    user.Email,
		Avatar:   meta.Avatar,
		Bio:      meta.Bio,
		CardName: meta.CardName,
		CardId:   meta.CardId,
	}
	if user.Status == 0 {
		result.Status = "正常"
	} else {
		result.Status = "封禁"
	}
	if meta.Sex == 1 {
		result.Sex = "男"
	} else {
		result.Sex = "女"
	}
	return result
}
