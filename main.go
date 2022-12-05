package main

import "giao/practice/server"

func main() {
	//r := gin.New()
	//r.GET("/profile", func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, gin.H{"234": "234"})
	//})
	//r.GET("/home", func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, gin.H{"home": "home"})
	//})
	//r.Run(":9999")
	server.Server()
}
