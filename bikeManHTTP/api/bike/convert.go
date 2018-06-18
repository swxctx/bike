package bikecontroller

import (
	"time"

	"github.com/swxctx/bike/model/bike"
)

func convertAlarmInfo(b *bike.BikeAlarm) *AlarmInfo {
	result := &AlarmInfo{
		Id:      b.Id,
		ById:    b.ById,
		AlarmTs: getTimeFormat(b.AlarmTs),
	}
	if b.MhCountOne == 1 || b.MhCountTwo == 1 {
		result.MhCount = "关闭"
	} else if b.MhCountOne == 0 && b.MhCountTwo == 0 {
		result.MhCount = "打开"
	} else {
		result.MhCount = "打开"
	}
	if b.SwCountOne == 1 {
		result.SwCount = "异常"
	} else {
		result.SwCount = "正常"
	}
	return result
}

func getTimeFormat(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05")
}
