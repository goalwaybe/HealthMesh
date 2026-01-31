package pkg

import (
	"github.com/smartwalle/alipay/v3"
	"medi-biz/common/config"
	"strconv"
)

func AliPagePay(OutTradeNo string, TotalAmount float64) string {
	data := config.ConfigData.Alipay
	privateKey := data.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	appId := data.AppId
	client, err := alipay.New(appId, privateKey, false)
	if err != nil {
		return ""
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = data.NotifyUrl
	p.ReturnURL = data.ReturnUrl
	p.Subject = "医疗项目"
	p.OutTradeNo = OutTradeNo
	p.TotalAmount = strconv.FormatFloat(TotalAmount, 'f', 1, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		return ""
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	return payURL
}
