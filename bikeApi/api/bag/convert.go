package bagcontroller

import (
	"github.com/swxctx/bike/model/bike"
)

func convertTopLog(top *bike.TopLog) *TopDetail {
	result := &TopDetail{
		Count:  top.Count,
		Status: top.Status,
		TopTs:  top.TopTs,
	}
	return result
}
