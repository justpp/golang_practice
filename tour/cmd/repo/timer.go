package cmd

import (
	"giao/tour/cmd/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间",
	Long:  "时间处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Fatalf("输出结果: %v %v", nowTime.Format("2006-04-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTime string
var duration string

var calculateCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-04-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-04-02"
			}
			if space == 1 {
				layout = "2006-04-02 15:04"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime: %v", err)
		}
		log.Printf("输出结果: %v %v", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateCmd)

	calculateCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间，有效单位为时间戳或已格式化后的时间`)
	calculateCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µs"), "ms", "s", "m", "h"`)
}
