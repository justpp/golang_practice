package cmd

import (
	"giao/src/tour/cmd/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModelLower = iota + 1
	ModelUpper
	ModelUnderScoreToLowerCamelCase
	ModelUnderScoreToUpperCamelCase
	ModelCamelCaseToUnderScore
)

var (
	str  string
	mode int8
	desc = strings.Join([]string{
		"该命令支持以下类型单词转换:",
		"1:全部转小写",
		"2:全部转大写",
		"3:下划线转小驼峰",
		"4:下划线转大驼峰",
		"5:驼峰转下划线",
	}, "/n")
	wordCmd = &cobra.Command{
		Use:   "word",
		Short: "单词转换",
		Long:  desc,
		Run: func(cmd *cobra.Command, args []string) {
			var content string
			switch mode {
			case ModelLower:
				content = word.ToLower(str)
			case ModelUpper:
				content = word.ToUpper(str)
			case ModelUnderScoreToLowerCamelCase:
				content = word.UnderScoreToLowerCameCase(str)
			case ModelUnderScoreToUpperCamelCase:
				content = word.UnderScoreToUpperCamelCase(str)
			case ModelCamelCaseToUnderScore:
				content = word.CamelCaseToUnderScore(str)
			default:
				log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
			}
			log.Printf("输出结果: %s", content)
		},
	}
)

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
