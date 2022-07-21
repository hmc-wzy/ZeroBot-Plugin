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

// 查成分的
func init() {
	// engine.OnRegex(`^>vup info\s?(.{1,25})$`).SetBlock(true).
	// 	Handle(func(ctx *zero.Ctx) {
	// 		keyword := ctx.State["regex_matched"].([]string)[1]
	// 		res, err := search(keyword)
	// 		if err != nil {
	// 			ctx.SendChain(message.Text("ERROR:", err))
	// 			return
	// 		}
	// 		id := strconv.FormatInt(res[0].Mid, 10)
	// 		// 获取详情
	// 		fo, err := fansapi(id)
	// 		if err != nil {
	// 			ctx.SendChain(message.Text("ERROR:", err))
	// 			return
	// 		}
	// 		ctx.SendChain(message.Text(
	// 			"search: ", fo.Mid, "\n",
	// 			"名字: ", fo.Uname, "\n",
	// 			"当前粉丝数: ", fo.Follower, "\n",
	// 			"24h涨粉数: ", fo.Rise, "\n",
	// 			"视频投稿数: ", fo.Video, "\n",
	// 			"直播间id: ", fo.Roomid, "\n",
	// 			"舰队: ", fo.GuardNum, "\n",
	// 			"直播总排名: ", fo.AreaRank, "\n",
	// 			"数据来源: ", "https://vtbs.moe/detail/", fo.Mid, "\n",
	// 			"数据获取时间: ", time.Now().Format("2006-01-02 15:04:05"),
	// 		))
	// 	})

	engine.OnFullMatch("zbphelloworld").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("少女祈祷中..."))

			ctx.SendChain(message.Text("vup已更新"))
		})
}
