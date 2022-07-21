package gmsr

import (
	"encoding/json"

	"github.com/FloatTech/zbputils/binary"
	"github.com/FloatTech/zbputils/web"
	"github.com/tidwall/gjson"
)

type userinfo struct {
	Name       string  `json:"name"`
	Mid        string  `json:"mid"`
	Face       string  `json:"face"`
	Fans       int64   `json:"fans"`
	Regtime    int64   `json:"regtime"`
	Attentions []int64 `json:"attentions"`
}

// 获取详情
func search(uid string) (result userinfo, err error) {
	cardURL := "https://account.bilibili.com/api/member/getCardByMid?mid=" + uid
	data, err := web.GetData(cardURL)
	if err != nil {
		return
	}
	err = json.Unmarshal(binary.StringToBytes(gjson.ParseBytes(data).Get("card").Raw), &result)
	if err != nil {
		return
	}
	return
}
