// generate by ginpt
package feedbackcontroller

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

// reportArgs 提交反馈请求参数
type reportArgs struct {
	// BaseParam 基础公共参数
	BaseParam
	ClientInfo     paramkits.ClientInfo
	CacheAPIKey    string
	CacheArguments []string
}

// Doreport 提交反馈
func Doreport(c *gin.Context) {
	params := reportArgs{
		CacheAPIKey: "feed_back:report",
	}
	c.Bind(&params.BaseParam)
	params.ClientInfo = paramkits.ParseClientInfo(c)

	result, hasError := doreport(c, &params)
	if hasError {
		kits.RenderError(c, result)
		return
	}
	kits.RenderSuccess(c, result)
}
