package cmd

import (
	"giao/pkg/e_book/src"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Flags().StringVarP(&url, "url", "u", "https://wap.x33xs.com/33xs/392/392417/", "小说目录页")
	cmd.Flags().IntVarP(&g, "g", "g", 20, "限制协程数")
}

var (
	url string
	g   int
	cmd = &cobra.Command{
		Short: "下载小说",
		Run: func(cmd *cobra.Command, args []string) {
			var tool = src.EBook{
				G: g,
			}
			tool.Run(url)
		},
	}
)

func Execute() error {
	return cmd.Execute()
}
