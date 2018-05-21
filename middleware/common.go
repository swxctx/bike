package middleware

import (
	"regexp"
	"strings"

	"github.com/domego/gokits"
	"github.com/domego/gokits/log"
)

func GenPassword(password string) string {
	return utils.MD5String([]byte(strings.TrimSpace(password)))
}

const (
	regxpValue = "^((13[0-9])|(15[^4])|(16[0-9])|(18[0-9])|(17[0-8])|(19[0-9])|(147,145))\\d{8}$"
)

// CheckPhone 正则校验手机号码
func CheckPhone(phone string) bool {
	r, err := regexp.Compile(strings.TrimSpace(regxpValue))
	if err != nil {
		log.Errorf("phone regexp compile error: %v", err)
		return false
	}
	return r.MatchString(strings.TrimSpace(phone))
	// 没有找到正则配置，则跳过校验
	return false
}
