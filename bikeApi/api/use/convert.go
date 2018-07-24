package usecontroller

import (
	"fmt"
	"time"

	"github.com/swxctx/bike/model/bike"
)

func convertUseLog(ulog *bike.UseLog) *Detail {
	result := &Detail{
		Amount:    fmt.Sprintf("%d元", ulog.Amount/100),
		BikeId:    ulog.ById,
		LatiTude:  ulog.LatiTude,
		LongiTude: ulog.LongiTude,
		StartTs:   getTimeFormat(ulog.StartTs),
		EndTs:     getTimeFormat(ulog.EndTs),
	}
	all := (ulog.EndTs - ulog.StartTs) / 60
	if all == 0 {
		all += 1
	}
	result.AllTs = fmt.Sprintf("%d分钟", all)
	return result
}

func getTimeFormat(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05")
}
