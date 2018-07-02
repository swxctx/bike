// generate by ginpt
package usercontroller

import (
	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/domego/ginkits/params"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/model/bike"
)

var _ = kits.RenderSuccess

type BaseParam paramkits.BaseParam

// doUserList 用户列表
func doUserList(c *gin.Context, params *UserListArgs) {
	list := &UserList{}
	users := bike.GetUserWhere("`deleted`=0")
	for _, user := range users {
		if user != nil {
			meta := bike.GetUserMetaFirst("`deleted`=0 AND `uid`=?", user.Id)
			list.List = append(list.List, convertUserProfile(user, meta))
		}
	}
	kits.RenderSuccess(c, gin.H{
		"list": list,
	})
}

// doForbid 封禁用户
func doForbid(c *gin.Context, params *ForbidArgs) {
	user := bike.GetUserFirst("`deleted`=0 AND `id`=?", params.Id)
	if user == nil {
		kits.RenderError(c, &errorkits.ErrorMessage{
			Code:    100301,
			Content: "用户不存在",
		})
		return
	}
	user.Status = params.Status
	user.Update()
	return
}
