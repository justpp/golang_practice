package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"sort"
	"strings"
)

var (
	token      = ""
	appId      = ""
	appSecret  = ""
	talkSecret = ""
)

func getConfig() *viper.Viper {
	vp := viper.New()
	vp.SetConfigName("env")
	vp.AddConfigPath("./")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil
	}
	return vp
}

func setEnv() {
	vp := getConfig()
	token, _ = vp.Get("token").(string)
	appId, _ = vp.Get("appId").(string)
	appSecret, _ = vp.Get("appSecret").(string)
	talkSecret, _ = vp.Get("talkSecret").(string)
}

func makeSignature(timestamp, nonce string) string { // 本地计算signature
	si := []string{token, timestamp, nonce}
	sort.Strings(si)            // 字典序排序
	str := strings.Join(si, "") // 组合字符串
	s := sha1.New()             // 返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      // WriteString函数将字符串数组str中的内容写入到s中
	return fmt.Sprintf("%x", s.Sum(nil))
}
