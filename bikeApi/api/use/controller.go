// generate by ginpt
package usecontroller

import (
	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/domego/gokits/log"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/middleware"
	"github.com/swxctx/bike/model/bike"
)

var _ = kits.RenderSuccess

type BaseParam bikemiddleware.BaseParam

// doHistoryLog 使用记录
func doHistoryLog(c *gin.Context, params *HistoryLogArgs) (result interface{}, hasError bool) {
	if params.BaseParam.Uid == 0 {
		return &errorkits.ErrorMessage{
			Code:    400101,
			Content: "登录已过期，请重新登录",
		}, true
	}
	listDetail := &UseDetail{}
	useLog := bike.GetUseLogWhere("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	for _, log := range useLog {
		listDetail.List = append(listDetail.List, convertUseLog(log))
	}
	return listDetail, false
}

// doStartUse 开始使用
func doStartUse(c *gin.Context, params *StartUseArgs) (result interface{}, hasError bool) {
	userMeta := bike.GetUserMetaFirst("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	if userMeta == nil {
		return &errorkits.ErrorMessage{
			Code:    400101,
			Content: "登录已过期，请重新登录",
		}, true
	}
	if len(userMeta.CardName) == 0 || len(userMeta.CardId) == 0 {
		return &errorkits.ErrorMessage{
			Code:    400102,
			Content: "请先进行身份实名认证",
		}, true
	}
	log.Infof("开始接入TCP服务...")
	log.Infof("开始发送指令(1)")
	log.Infof("指令发送成功")
	log.Infof("开锁成功")
	return
}
