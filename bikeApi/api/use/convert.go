package usecontroller

import (
	"github.com/swxctx/bike/model/bike"
)

func convertUseLog(ulog *bike.UseLog) *Detail {
	result := &Detail{
		Amount:    ulog.Amount,
		BikeId:    ulog.ById,
		LatiTude:  ulog.LatiTude,
		LongiTude: ulog.LongiTude,
		StartTs:   ulog.StartTs,
		EndTs:     ulog.EndTs,
	}
	return result
}
