package gmsr

import (
	"encoding/json"

	"github.com/FloatTech/zbputils/binary"
	"github.com/FloatTech/zbputils/web"
	"github.com/tidwall/gjson"
)

type userinfo struct {
	Name               string `json:"Name"`
	Level              string `json:"Level"`
	ExpPercent         string `json:"ExpPercent"`
	Class              string `json:"Class"`
	CharacterImageURL  string `json:"CharacterImageURL"`
	LegionLevel        string `json:"LegionLevel"`
	LegionRank         string `json:"LegionRank"`
	LegionPower        string `json:"LegionPower"`
	LegionCoinsPerDay  string `json:"LegionCoinsPerDay"`
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
