package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var WETHER_CODE = map[string]string{"200": "曇りところにより雨", "202": "曇り 昼過ぎ 一時 雨", "205": "くもり"}
var AREA_CODE = map[string]string{"020030": "三八上北", "020010": "津軽", "020020": "下北"}

func saveWeatherInfoJsonFile(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	var data []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	filepath := filepath.Join(dir, "weather/weather_info.json")
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error createing file!: ", err)
		return
	}
	defer file.Close()
	
	jsonData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

}

// 地域別のurlを生成 
func generate_url() {
	
}
func main() {
	// 020000が青森のエリアコード
	url := "https://www.jma.go.jp/bosai/forecast/data/forecast/020000.json"
	saveWeatherInfoJsonFile(url)
	
	fmt.Println("JSONデータをファイルに保存しました。", "weather/weather_info.json")
}