// generate by ginpt
package feedbackcontroller

import (
	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/middleware"
	"github.com/swxctx/bike/model/bike"
)

var _ = kits.RenderSuccess

type BaseParam bikemiddleware.BaseParam

// doreport 提交反馈
func doreport(c *gin.Context, params *reportArgs) (result interface{}, hasError bool) {
	if params.BaseParam.Uid == 0 || len(params.Content) == 0 {
		return &errorkits.ErrorMessage{
			Code:    200101,
			Content: "操作不合法",
		}, true
	}
	back := bike.NewFeedback()
	back.Uid = params.BaseParam.Uid
	back.Content = params.Content
	back.Insert()
	return
}
