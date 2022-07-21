package gmsr

import (
	"encoding/json"

	"github.com/FloatTech/zbputils/binary"
	"github.com/FloatTech/zbputils/web"
	"github.com/tidwall/gjson"
)

type userinfo struct {
	Name               string `json:"Name"`
	Level              int64  `json:"Level"`
	ExpPercent         int64  `json:"ExpPercent"`
	Class              int64  `json:"Class"`
	CharacterImageURL  int64  `json:"CharacterImageURL"`
	LegionLevel        int64  `json:"LegionLevel"`
	LegionRank         int64  `json:"LegionRank"`
	LegionPower        int64  `json:"LegionPower"`
	LegionCoinsPerDay  int64  `json:"LegionCoinsPerDay"`
	ServerClassRanking string `json:"ServerClassRanking"`
	ServerRank         string `json:"ServerRank"`
}

// 获取详情 CharacterData
func search(uid string) (result userinfo, err error) {
	cardURL := "https://api.maplestory.gg/v2/public/character/gms/" + uid
	data, err := web.GetData(cardURL)
	if err != nil {
		return
	}
	err = json.Unmarshal(binary.StringToBytes(gjson.ParseBytes(data).Get("CharacterData").Raw), &result)
	if err != nil {
		return
	}
	return
}
