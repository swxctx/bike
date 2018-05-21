package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.xiaoenai.net/xserver/utils"

	kits "github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/domego/ginkits/params"
	"github.com/domego/gokits/log"
	"github.com/gin-gonic/gin"
)

func AccessTokenTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer log.Tracef("end access token handler")
		log.Tracef("start access token handler")

		// gat all params
		params := paramkits.GetParams(c)

		accessToken := params["access_token"]
		if !paramkits.Required(c, "access_token", accessToken) {
			return
		}
		// 解析 uid
		uid, err := GetUidFromAccessToken(accessToken)
		if err != nil {
			kits.RenderError(c, &errorkits.ErrorMessage{
				Code:    errorkits.ErrorCodeArgumentValueInvalid,
				Message: "access_token无效",
				Content: "网络不给力",
			})
			return
		}
		// set uid to form
		c.Request.Form["uid"] = []string{strconv.FormatInt(int64(uid), 10)}
		log.Tracef("post form: %+v", c.Request.PostForm)
		c.Next()
	}
}

// GetUidFromXeaAccessToken 解析uid
func GetUidFromAccessToken(accessToken string) (uid int32, err error) {
	ss := strings.Split(accessToken, "x")
	if len(ss) <= 0 {
		err = fmt.Errorf("invalid xea_access_token, %s", accessToken)
		return
	}
	i, err := strconv.ParseInt(ss[0], 35, 32)
	if err != nil {
		return
	}
	uid = int32(i)
	return
}

// GenAccessToken 生成 access token
func GenAccessToken(uid int32) string {
	return strconv.FormatInt(int64(uid), 35) + "x" + utils.RandString(24)
}
