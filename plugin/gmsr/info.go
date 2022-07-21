// Package bilibili 查询b站用户信息
package gmsr

import (
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var engine = control.Register("gmsr", &ctrl.Options[*zero.Ctx]{
	DisableOnDefault: false,
	Help: "gmsr\n" +
		"查询+角色名\n",
	PublicDataFolder: "Gmsr",
})

// 查询GMS角色信息
func init() {
	engine.OnRegex(`^查询\s*([A-Za-z0-9]+)$`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			keyword := ctx.State["regex_matched"].([]string)[1]
			u, err := search(keyword)
			if err != nil {
				//ctx.SendChain(message.Text("ERROR:", err))
				ctx.SendChain(message.Text("没找到角色名为【", keyword, "】的信息"))
				return
			}
			ctx.SendChain(message.Image(u.CharacterImageURL), message.Text(
				"名字: ", u.Name, "\n",
				"等级: ", u.Level, "\n",
				"当前经验百分比: ", u.ExpPercent, "\n",
				"职业: ", u.Class, "\n",
				"联盟等级: ", u.LegionLevel, "\n",
				"联盟排名: ", u.LegionRank, "\n",
				"联盟战斗力: ", u.LegionPower, "\n",
				"每日获得联盟币: ", u.LegionCoinsPerDay, "\n",
				"职业等级排名: ", u.ServerClassRanking, "\n",
				"总等级排名: ", u.ServerRank, "\n",
			))
		})

	engine.OnFullMatch("zbphelloworld").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("少女祈祷中..."))

			ctx.SendChain(message.Text("vup已更新"))
		})
}
