package weather

import (
	"testing"
)

var wether_code = map[string]string{"200": "曇りところにより雨", "202": "曇り 昼過ぎ 一時 雨", "205": "くもり"}
var area_code = map[string]string{"020030": "三八上北", "020010": "津軽", "020020": "下北"}

func TestGetWeatherInfo(){
	// 020000が青森のエリアコード
	url := "https://www.jma.go.jp/bosai/forecast/data/forecast/020000.json"

}