package cmd

import (
	"giao/pkg/e_book/src"
	"giao/pkg/e_book/src/sites"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Flags().StringVarP(&url, "url", "u", "https://m2.ddyueshu.com/wapbook/11082821.html", "小说目录页")
	cmd.Flags().IntVarP(&g, "g", "g", 20, "限制协程数")
}

var (
	url string
	g   int
	cmd = &cobra.Command{
		Short: "下载小说",
		Run: func(cmd *cobra.Command, args []string) {
			var site = sites.New().InitImport().GetSite(url)
			var tool = src.EBook{
				G:    g,
				Site: site,
			}
			tool.Run(url)
		},
	}
)

func Execute() error {
	return cmd.Execute()
}
