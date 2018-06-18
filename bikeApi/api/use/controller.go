// generate by ginpt
package usecontroller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/domego/ginkits"
	"github.com/domego/ginkits/errors"
	"github.com/domego/gokits/log"
	"github.com/gin-gonic/gin"

	"github.com/swxctx/bike/bikeApi/cfg"
	"github.com/swxctx/bike/middleware"
	"github.com/swxctx/bike/model/bike"
	"github.com/swxctx/bike/push"
)

var _ = kits.RenderSuccess

type BaseParam bikemiddleware.BaseParam

// doHistoryLog 使用记录
func doHistoryLog(c *gin.Context, params *HistoryLogArgs) (result interface{}, hasError bool) {
	if params.BaseParam.Uid == 0 {
		return &errorkits.ErrorMessage{
			Code:    400101,
			Content: "登录已过期，请重新登录",
		}, true
	}
	listDetail := &UseDetail{}
	useLog := bike.GetUseLogWhere("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	for _, log := range useLog {
		listDetail.List = append(listDetail.List, convertUseLog(log))
	}
	return listDetail, false
}

type PushInfo struct {
	Ip  string `json:"ip"`
	Msg string `json:"msg"`
}

// MarshalPush json编码pushinfo
func MarshalPush(data *PushInfo) string {
	result, err := json.Marshal(data)
	if err != nil {
		log.Errorf("push info json marshal error: %v", err)
	}
	return string(result)
}

// doStartUse 开始使用
func doStartUse(c *gin.Context, params *StartUseArgs) (result interface{}, hasError bool) {
	userMeta := bike.GetUserMetaFirst("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	if userMeta == nil {
		return &errorkits.ErrorMessage{
			Code:    400101,
			Content: "登录已过期，请重新登录",
		}, true
	}
	if len(userMeta.CardName) == 0 || len(userMeta.CardId) == 0 {
		return &errorkits.ErrorMessage{
			Code:    400102,
			Content: "请先进行身份实名认证",
		}, true
	}
	userBag := bike.GetUserBagFirst("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	if userBag == nil {
		return &errorkits.ErrorMessage{
			Code:    400102,
			Content: "余额不足，请充值",
		}, true
	}
	if userBag.Balance < 100 {
		return &errorkits.ErrorMessage{
			Code:    400102,
			Content: "余额不足，请充值",
		}, true
	}
	// 查询单车信息
	bikeInfo := bike.GetBikeFirst("`deleted`=0 AND `by_id`=?", params.BikeId)
	if bikeInfo == nil {
		return &errorkits.ErrorMessage{
			Code:    400103,
			Content: "网络出错，稍候再试",
		}, true
	}
	if bikeInfo.Status == 1 {
		return &errorkits.ErrorMessage{
			Code:    400104,
			Content: "此单车暂时不可用，请选择其他车辆",
		}, true
	}

	// message(0:->开锁)
	pushinfo := &PushInfo{}
	pushinfo.Msg = "0"
	// bike ip
	pushinfo.Ip = fmt.Sprintf("%s:%s", bikeInfo.Ip, bikeInfo.Port)
	message := MarshalPush(pushinfo)
	log.Infof("TCP服务开始发送指令进行单车开锁...")
	log.Infof("Bike ip:%s", pushinfo.Ip)
	// cfg.Conf.PushIp 推送服务ip
	err := push.DoTcpClient(cfg.Conf.PushIp, message)
	if err != nil {
		log.Errorf("Lock On err->%v", err)
		return &errorkits.ErrorMessage{
			Code:    400105,
			Content: "网络出错，请稍后再试",
		}, true
	}

	// 存储记录
	useLog := bike.NewUseLog()
	useLog.Uid = params.BaseParam.Uid
	useLog.ById = params.BikeId
	useLog.LatiTude = params.PointLat
	useLog.LongiTude = params.PointLng
	useLog.StartTs = time.Now().Unix()
	useLog.Insert()
	return
}

// doFinishUse 结束使用
func doFinishUse(c *gin.Context, params *FinishUseArgs) (result interface{}, hasError bool) {
	useLog := bike.GetUseLogFirst("`deleted`=0 AND `by_id`=? ORDER BY `created_at` DESC", params.BikeId)
	if useLog == nil {
		return &errorkits.ErrorMessage{
			Code:    400106,
			Content: "网络不给力，请重试",
		}, true
	}
	if useLog != nil {
		ts := time.Now().Unix()
		useLog.EndTs = ts
		// 本次骑行消耗金额数
		useLog.Amount = 100
		useLog.Update()
	}
	useBag := bike.GetUserBagFirst("`deleted`=0 AND `uid`=?", params.BaseParam.Uid)
	if useBag != nil {
		useBag.Balance -= 100
		useBag.Update()
	}
	bikeInfo := bike.GetBikeFirst("`deleted`=0 AND `by_id`=?", params.BikeId)
	if bikeInfo != nil {
		bikeInfo.LatiTude = params.PointLat
		bikeInfo.LongiTude = params.PointLng
		bikeInfo.Update()
	}

	// message(0:->开锁)
	pushinfo := &PushInfo{}
	pushinfo.Msg = "1"
	// bike ip
	pushinfo.Ip = fmt.Sprintf("%s:%s", bikeInfo.Ip, bikeInfo.Port)
	message := MarshalPush(pushinfo)
	log.Infof("TCP服务开始发送指令进行单车上锁...")
	log.Infof("Bike ip:%s", pushinfo.Ip)
	// cfg.Conf.PushIp 推送服务ip
	err := push.DoTcpClient(cfg.Conf.PushIp, message)
	if err != nil {
		log.Errorf("Lock Off err->%v", err)
	}
	return useLog, false
}
