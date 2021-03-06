//generated by gpt
package bagcontroller

import (
	"time"
)

var _ = time.Now

type BalanceDetail struct {
	Balance int32        `json:"balance"`
	List    []*TopDetail `json:"list"`
}

type TopDetail struct {
	Count  int32  `json:"count"`
	Phone  string `json:"phone"`
	Status string `json:"status"`
	TopTs  string `json:"top_ts"`
}
