package misc

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/teamssix/cf/pkg/cloud"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "关于作者 (About me)",
	Long:  `关于作者 (About me)`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Print(`
嗨, 我是 TeamsSix，很开心你能找到这儿，你可以在下面的平台中找到并关注我。
Hi, I'm TeamsSix and I'm glad you've found this place. You can find and follow me on the social platforms and links below.

`)
		data := [][]string{
			{"@teamssix", "TeamsSix", "teamssix.com", "github.com/teamssix", "wiki.teamssix.com", "狼组安全团队 @wgpsec"},
		}
		var header = []string{"推特 (Twitter)", "微信公众号 (WeChat Official Accounts)", "博客 (Blog)", "Github", "云安全知识库 T Wiki", "所属团队 (Organization)"}
		var td = cloud.TableData{Header: header, Body: data}
		cloud.PrintTable(td, "")
		color.Print(`
如果你使用着感觉还不错，记得给个 Star 哦<gray>（另外 T Wiki 是我自己在维护的云安全知识库，如果你想加入云安全交流群，那么在 T Wiki 中可以找到）</>
Hopefully this repository will reach 1k stars. <bold>I can do all this through him who gives me strength.</>
`)
	},
}
