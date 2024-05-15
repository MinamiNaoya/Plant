package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	
)



// 地域別のurlを生成 
func generate_url(area_code AREACODE) (url string){
	base_url := "https://www.jma.go.jp/bosai/forecast/data/forecast/"
	url = base_url + string(area_code)
	return url
}

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



func GetWeatherInfo(w http.ResponseWriter, r *http.Request) {

	area_code := GetAreaCode(w, r)
	
	url := generate_url(area_code)
	saveWeatherInfoJsonFile(url)
	
	fmt.Println("JSONデータをファイルに保存しました。", "weather/weather_info.json")
}