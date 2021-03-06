package keymanage

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/teamssix/cf/pkg/cloud"
	"github.com/teamssix/cf/pkg/util/cmdutil"
	"strings"
)

var AddKeyCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"update", "a"},
	Short:   "添加密钥 (Add Key)",
	Long: "添加密钥到数据库, 如果想要更新密钥的标记,也可以使用这个方法 " +
		"(Add Key to framwork database, if you want to update key remark or name, you can also use this method)",
	Run: func(cmd *cobra.Command, args []string) {
		AddOrUpdate()
	},
}

// AddOrUpdate 增加或更新密钥, 如果密钥已经存在, 则更新, 否则增加. 通过 AccessKeyID 来判断是否已经存在.
func AddOrUpdate() {
	cloudConfigList, _ := cmdutil.ReturnCloudProviderList()
	var qs = []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "请为当前使用的 Key 输入名称 (Please input name for current using key)"},
			Validate: survey.Required,
		},
		{
			Name: "remark",
			Prompt: &survey.Input{
				Message: "请为当前使用的 Key 输入备注 (可选) " +
					"(Please input remark for current using key)[Optional]",
			},
		},
		{
			Name: "platform",
			Prompt: &survey.Select{
				Message: "Key 所属的云服务平台",
				Options: cloudConfigList,
			},
		},
		{
			Name:     "AccessKeyId",
			Prompt:   &survey.Input{Message: "Access Key Id (必须 Required):"},
			Validate: survey.Required,
		},
		{
			Name:     "AccessKeySecret",
			Prompt:   &survey.Input{Message: "Access Key Secret (必须 Required):"},
			Validate: survey.Required,
		},
		{
			Name:   "STSToken",
			Prompt: &survey.Input{Message: "STS Token (可选 Optional):"},
		},
	}

	// Generate the new config struct named cred to receive the inputted values.
	cred := struct {
		Name            string `survey:"name"`
		Remark          string `survey:"remark"`
		Platform        string `survey:"platform"`
		AccessKeyId     string `survey:"AccessKeyId"`
		AccessKeySecret string `survey:"AccessKeySecret"`
		STSToken        string `survey:"STSToken"`
	}{}
	survey.Ask(qs, &cred)

	key := Key{
		Name:     cred.Name,
		Remark:   cred.Remark,
		Platform: cred.Platform,
		Config: cloud.Config{
			AccessKeyId:     strings.TrimSpace(cred.AccessKeyId),
			AccessKeySecret: strings.TrimSpace(cred.AccessKeySecret),
			STSToken:        strings.TrimSpace(cred.STSToken),
		},
	}

	// Make user to check
	PrintKeysTable([]Key{key})
	promot := &survey.Confirm{
		Message: "以上信息是否正确 (make sure correctness) "}
	sure := true // Break out
	survey.AskOne(promot, &sure)
	if sure {
		KeyDb.Where("access_key_id = ?", key.AccessKeyId).Save(&key)
	}
}
