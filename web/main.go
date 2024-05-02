package main

import (
	"net/http"

	"github.com/MinamiNaoya/Plant/weather"
	"github.com/gin-gonic/gin"
)



func main(){	
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.POST("/getAreaCode", weather.GetAreaCode)
	
	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run()
}