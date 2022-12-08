package wx

import (
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	setEnv()
}

// Server  微信server
func Server() {
	router := gin.Default()
	router.GET("/check", CheckSignature)
	router.GET("/fetch_code", fetchCode)
	router.GET("/get_userinfo", getUserInfo)
	router.LoadHTMLGlob("wx/template/*")

	log.Println("Wechat Service: Start!")
	router.Run(":80")
	log.Println("Wechat Service: Stop!")
}
