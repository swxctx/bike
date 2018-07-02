package bagcontroller

import (
	"time"

	"github.com/swxctx/bike/model/bike"
)

func convertTopLog(top *bike.TopLog) *TopDetail {
	result := &TopDetail{
		Count: top.Count / 100,
		TopTs: getTimeFormat(top.TopTs),
		Phone: top.Phone,
	}
	if top.Status == 0 {
		result.Status = "未处理"
	} else {
		result.Status = "已处理"
	}
	return result
}

func getTimeFormat(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05")
}
