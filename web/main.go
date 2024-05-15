package main

import (

	"github.com/gin-gonic/gin"
	"github.com/MinamiNaoya/Plant/weather"
)


func main(){	
	weather_service := weather.WeatherService{}
	get_weather_info := func(ctx *gin.Context) {
    	weather_service.GetWeatherInfo(ctx.Writer, ctx.Request)
	}

	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.POST("/getWeatherInfo", 	get_weather_info)
	
	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run()
}