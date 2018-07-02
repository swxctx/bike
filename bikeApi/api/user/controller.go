// generate by ginpt
package usercontroller

import (
	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"

	"github.com/domego/gokits/log"
	"github.com/swxctx/bike/middleware"
	"github.com/swxctx/bike/model/bike"

	"github.com/gin-gonic/gin"
)

var _ = log.Tracef
var _ = kits.RenderSuccess

type BaseParam bikemiddleware.BaseParam

// doRegister 注册
func doRegister(c *gin.Context, params *RegisterArgs) (result interface{}, hasError bool) {
	// 1.查询该用户是否注册过
	isuser := bike.GetUserFirst("`deleted`=0 AND `username`=?", params.Username)
	if isuser != nil {
		return &errorkits.ErrorMessage{
			Code:    100101,
			Content: "用户名已经被注册了",
		}, true
	}
	// 2.校验手机号
	if !bikemiddleware.CheckPhone(params.Phone) {
		return &errorkits.ErrorMessage{
			Code:    100102,
			Content: "手机号不合法",
		}, true
	}
	// 3.校验用户名
	if len(params.Username) > 12 {
		return &errorkits.ErrorMessage{
			Code:    100103,
			Content: "用户名长度超出限制",
		}, true
	}
	// 4.校验密码
	if len(params.Password) > 16 || len(params.Password) < 6 {
		return &errorkits.ErrorMessage{
			Code:    100104,
			Content: "密码长度为6-16",
		}, true
	}
	user := bike.NewUser()
	user.Username = params.Username
	user.Phone = params.Phone
	user.Email = params.Email
	user.Password = bikemiddleware.GenPassword(params.Password)
	user.Insert()
	meta := bike.NewUserMeta()
	meta.Uid = user.Id
	meta.Bio = params.Bio
	meta.Sex = params.Sex
	meta.Insert()

	loginResult := convertLoginResult(user)
	return loginResult, false
}

// doLogin 登录
func doLogin(c *gin.Context, params *LoginArgs) (result interface{}, hasError bool) {
	pass := bikemiddleware.GenPassword(params.Password)
	user := bike.GetUserFirst("`deleted`=0 AND `username`=? AND `password`=?", params.Username, pass)
	if user == nil {
		return &errorkits.ErrorMessage{
			Code:    100201,
			Content: "请检查用户名与密码是否错误",
		}, true
	}
	loginResult := convertLoginResult(user)
	return loginResult, false
}

// doGetProfile 获取资料
func doGetProfile(c *gin.Context, params *GetProfileArgs) (result interface{}, hasError bool) {
	user := bike.GetUserFirst("`deleted`=0 AND `id`=?", params.BaseParam.Uid)
	meta := bike.GetUserMetaFirst("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	if user == nil || user == nil {
		return &errorkits.ErrorMessage{
			Code:    100301,
			Content: "暂时查询不到您的个人信息，请稍后再试",
		}, true
	}
	return convertUserProfile(user, meta), false
}

// doAuthCard 实名认证
func doAuthCard(c *gin.Context, params *AuthCardArgs) (result interface{}, hasError bool) {
	if len(params.CardId) != 18 || len(params.CardName) == 0 {
		return &errorkits.ErrorMessage{
			Code:    100401,
			Content: "您输入的身份证信息不合法，请重新输入",
		}, true
	}
	meta := bike.GetUserMetaFirst("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	if meta == nil {
		return &errorkits.ErrorMessage{
			Code:    100402,
			Content: "网络出错了，请稍后再试",
		}, true
	}
	if len(meta.CardId) != 0 && len(meta.CardName) != 0 {
		return &errorkits.ErrorMessage{
			Code:    100403,
			Content: "您已经认证过了，请勿重复提交",
		}, true
	}
	meta.CardId = params.CardId
	meta.CardName = params.CardName
	meta.Update()
	return
}

// doUpdateProfile 更新资料
func doUpdateProfile(c *gin.Context, params *UpdateProfileArgs) (result interface{}, hasError bool) {
	user := bike.GetUserFirst("`deleted`=0 AND `id`=?", params.BaseParam.Uid)
	meta := bike.GetUserMetaFirst("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	if user == nil || meta == nil {
		return &errorkits.ErrorMessage{
			Code:    100501,
			Content: "更新失败，请稍后再试",
		}, true
	}
	var (
		isUpadte bool
	)
	if user.Phone != params.Phone {
		if !bikemiddleware.CheckPhone(params.Phone) {
			return &errorkits.ErrorMessage{
				Code:    100102,
				Content: "手机号不合法",
			}, true
		}
		user.Phone = params.Phone
		isUpadte = true
	}
	if user.Username != params.Username {
		if len(params.Username) > 12 {
			return &errorkits.ErrorMessage{
				Code:    100502,
				Content: "用户名长度超出限制",
			}, true
		}
		isuser := bike.GetUserFirst("`deleted`=0 AND `username`=? AND `id`!=?", params.Username, params.BaseParam.Uid)
		if isuser != nil {
			return &errorkits.ErrorMessage{
				Code:    100101,
				Content: "用户名已经存在了",
			}, true
		}
		isUpadte = true
	}
	if user.Email != params.Email {
		user.Email = params.Email
		isUpadte = true
	}
	if meta.Bio != params.Bio && params.Bio != "暂未设置签名" {
		meta.Bio = params.Bio
		isUpadte = true
	}
	if meta.Sex != params.Sex {
		meta.Sex = params.Sex
		isUpadte = true
	}
	if isUpadte {
		user.Update()
		meta.Update()
	}
	return
}

// doUpdatePassword 修改密码
func doUpdatePassword(c *gin.Context, params *UpdatePasswordArgs) {
	user := bike.GetUserFirst("`deleted`=0 AND `id`=?", params.BaseParam.Uid)
	if user == nil {
		kits.RenderError(c, &errorkits.ErrorMessage{
			Code:    100102,
			Content: "网络不给力，请稍后再试",
		})
		return
	}
	pass := bikemiddleware.GenPassword(params.OldPassword)
	if user.Password != pass {
		kits.RenderError(c, &errorkits.ErrorMessage{
			Code:    100103,
			Content: "旧密码输入错误，请确认",
		})
		return
	}
	user.Password = bikemiddleware.GenPassword(params.Password)
	user.Update()
}
