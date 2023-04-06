package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

const host = "168.138.33.211:8892"

func init() {
	setEnv()

}

// Server  微信server
func Server() {
	router := gin.Default()
	router.GET("/check", CheckSignature)
	router.POST("/check", ReceiveMsg)
	router.GET("/fetch_code", fetchCode)
	router.GET("/get_userinfo", getUserInfo)
	router.GET("/set_menu", setMenu)
	router.GET("/del_menu", delMenu)
	router.LoadHTMLGlob("./template/*")

	log.Println("Wechat Service: Start!")
	router.Run(":8892")
	log.Println("Wechat Service: Stop!")
}
