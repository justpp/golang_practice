package wx

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"log"
	"sort"
	"strings"
)

var token = ""

func init() {
	env := getConfig()
	token, _ = env.Get("token").(string)
}

func WxServer() {
	router := gin.Default()
	//router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	log.Println("Wechat Service: Start!")
	router.GET("/check", WxCheckSignature)

	router.Run(":80")
	log.Println("Wechat Service: Stop!")
}

func WxCheckSignature(c *gin.Context) {
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	signature := c.Query("signature")
	echostr := c.Query("echostr")
	signatureGen := makeSignature(timestamp, nonce)
	if signatureGen != signature {
		log.Println("微信公众号接入失败")
		return
	}
	log.Println("微信公众号接入成功")
	_, _ = c.Writer.WriteString(echostr)
}

func makeSignature(timestamp, nonce string) string { //本地计算signature
	si := []string{token, timestamp, nonce}
	sort.Strings(si)            //字典序排序
	str := strings.Join(si, "") //组合字符串
	s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
	return fmt.Sprintf("%x", s.Sum(nil))
}

func getConfig() *viper.Viper {
	vp := viper.New()
	vp.SetConfigName("env")
	vp.AddConfigPath("wx/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil
	}
	return vp
}
