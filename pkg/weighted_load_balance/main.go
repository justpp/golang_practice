package main

import (
	"giao/pkg/weighted_load_balance/src"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		_, _ = c.Writer.Write([]byte("hihi"))
	})

	r.StaticFile("/c", "./html/weighted_echars.html")

	weightedLoad := src.WeightRoundBalance{}
	weightedLoad.Add("192.168.1.2", 3)
	weightedLoad.Add("192.168.1.3", 3)
	weightedLoad.Add("192.168.1.4", 2)
	weightedLoad.Add("192.168.1.5", 1)
	r.GET("/c/data", func(c *gin.Context) {
		weightedLoad.Next()
		c.JSON(200, weightedLoad.GetCharsData())
	})

	err := r.Run(":9991")
	if err != nil {
		log.Fatalln(err)
	}

}
