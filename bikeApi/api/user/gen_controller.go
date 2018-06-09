// generate by ginpt
package usercontroller

import (
	"fmt"

	"github.com/domego/ginkits"
	"github.com/domego/ginkits/params"
	"github.com/domego/gokits"
	"github.com/domego/gokits/log"
	"github.com/gin-gonic/gin"
)

var _ = log.Tracef
var _ = kits.RenderSuccess
var _ = utils.IsEmpty
var _ = fmt.Sprintf

// RegisterArgs 注册请求参数
type RegisterArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	Username string `form:"username" json:"username"`
	Phone    string `form:"phone" json:"phone"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
	Sex      int32  `form:"sex" json:"sex"`
	Bio      string `form:"bio" json:"bio"`
}

// DoRegister 注册
func DoRegister(c *gin.Context) {
	params := RegisterArgs{
		CacheAPIKey: "user:register",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
		if utils.IsEmpty(params.Username) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "username is required",
			})
			return
		}

		if utils.IsEmpty(params.Phone) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "phone is required",
			})
			return
		}

		if utils.IsEmpty(params.Password) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "password is required",
			})
			return
		}

	}

	result, hasError := doRegister(c, &params)
	if hasError {
		kits.RenderError(c, result)
		return
	}
	kits.RenderSuccess(c, result)
}

// LoginArgs 登录请求参数
type LoginArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// DoLogin 登录
func DoLogin(c *gin.Context) {
	params := LoginArgs{
		CacheAPIKey: "user:login",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
		if utils.IsEmpty(params.Username) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "username is required",
			})
			return
		}

		if utils.IsEmpty(params.Password) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "password is required",
			})
			return
		}
	}

	result, hasError := doLogin(c, &params)
	if hasError {
		kits.RenderError(c, result)
		return
	}
	kits.RenderSuccess(c, result)
}

// GetProfileArgs 获取资料请求参数
type GetProfileArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string
}

// DoGetProfile 获取资料
func DoGetProfile(c *gin.Context) {
	params := GetProfileArgs{
		CacheAPIKey: "user:get_profile",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)

	result, hasError := doGetProfile(c, &params)
	if hasError {
		kits.RenderError(c, result)
		return
	}
	kits.RenderSuccess(c, result)
}

// AuthCardArgs 实名认证请求参数
type AuthCardArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	CardName string `form:"card_name" json:"card_name"`
	CardId   string `form:"card_id" json:"card_id"`
}

// DoAuthCard 实名认证
func DoAuthCard(c *gin.Context) {
	params := AuthCardArgs{
		CacheAPIKey: "user:auth_card",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
		if utils.IsEmpty(params.CardName) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "card_name is required",
			})
			return
		}

		if utils.IsEmpty(params.CardId) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "card_id is required",
			})
			return
		}
	}

	result, hasError := doAuthCard(c, &params)
	if hasError {
		kits.RenderError(c, result)
		return
	}
	kits.RenderSuccess(c, result)
}

// UpdateProfileArgs 更新资料请求参数
type UpdateProfileArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	Username string `form:"username" json:"username"`
	Phone    string `form:"phone" json:"phone"`
	Email    string `form:"email" json:"email"`
	Bio      string `form:"bio" json:"bio"`
	Sex      int32  `form:"sex" json:"sex"`
}

// DoUpdateProfile 更新资料
func DoUpdateProfile(c *gin.Context) {
	params := UpdateProfileArgs{
		CacheAPIKey: "user:update_profile",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
		if utils.IsEmpty(params.Username) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "username is required",
			})
			return
		}

		if utils.IsEmpty(params.Phone) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "phone is required",
			})
			return
		}

		if utils.IsEmpty(params.Email) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "email is required",
			})
			return
		}

		if utils.IsEmpty(params.Bio) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "bio is required",
			})
			return
		}

		if utils.IsEmpty(params.Sex) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "sex is required",
			})
			return
		}
	}

	result, hasError := doUpdateProfile(c, &params)
	if hasError {
		kits.RenderError(c, result)
		return
	}
	kits.RenderSuccess(c, result)
}

// UpdatePasswordArgs 修改密码请求参数
type UpdatePasswordArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string

	OldPassword string `form:"old_password" json:"old_password"`
	Password    string `form:"password" json:"password"`
}

// DoUpdatePassword 修改密码
func DoUpdatePassword(c *gin.Context) {
	params := UpdatePasswordArgs{
		CacheAPIKey: "user:update_password",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)
	c.Bind(&params)

	{
		if utils.IsEmpty(params.OldPassword) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "old_password is required",
			})
			return
		}

		if utils.IsEmpty(params.Password) {
			kits.RenderError(c, &kits.RespErrorMessage{
				Code:    kits.ErrorCodeArgumentLack,
				Message: "password is required",
			})
			return
		}
	}

	doUpdatePassword(c, &params)

}
