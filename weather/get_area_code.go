package weather

import (
	"encoding/json"
	"net/http"
)


func GetAreaCode(w http.ResponseWriter, r *http.Request) (areaCode AREACODE){
	var requestData struct {
		Ken string `json:"ken"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	switch requestData.Ken {
	case "Hokkaido":
		areaCode = Hokkaido
	case "Aomori": 
		areaCode = Aomori
	case "Akita":
		areaCode = Akita
	case "Yamagata":
		areaCode = Yamagata
	case "Iwate":
		areaCode = Iwate
	case "Miyagi":
		areaCode = Miyagi
	case "Fukushima":
		areaCode = Fukushima
	case "Ibaraki":
		areaCode = Ibaraki
	case "Saitama":
		areaCode = Saitama
	case "Tochigi":
		areaCode = Tochigi
	case "Gunma":
		areaCode = Gunma
	case "Kanagawa":
		areaCode = Kanagawa
	case "Chiba":
		areaCode = Chiba
	case "Tokyo":
		areaCode = Tokyo
	case "Nigata":
		areaCode = Nigata
	case "Toyama":
		areaCode = Toyama
	case "Ishikawa":
		areaCode = Ishikawa
	case "Fukui":
		areaCode = Fukui
	case "Yamanashi":
		areaCode = Yamanashi
	case "Nagano":
		areaCode = Nagano
	case "Gifu":
		areaCode = Gifu
	case "Shizuoka":
		areaCode = Shizuoka
	case "Aichi":
		areaCode = Aichi
	case "Mie":
		areaCode = Mie
	case "Shiga":
		areaCode = Shiga
	case "Kyoto":
		areaCode = Kyoto
	case "Osaka":
		areaCode = Osaka
	case "Hyogo":
		areaCode = Hyogo
	case "Nara":
		areaCode = Nara
	case "Wakayama":
		areaCode = Wakayama
	case "Tottori":
		areaCode = Tottori
	case "Shimane":
		areaCode = Shimane
	case "Okayama":
		areaCode = Okayama
	case "Hiroshima":
		areaCode = Hiroshima
	case "Yamaguchi":
		areaCode = Yamaguchi
	case "Tokushima":
		areaCode = Tokushima
	case "Kagawa":
		areaCode = Kagawa
	case "Ehime":
		areaCode = Ehime
	case "Kochi":
		areaCode = Kochi
	case "Fukuoka":
		areaCode = Fukuoka
	case "Saga":
		areaCode = Saga
	case "Nagasaki":
		areaCode = Nagasaki
	case "Kumamoto":
		areaCode = Kumamoto
	case "Oita":
		areaCode = Oita
	case "Miyazaki":
		areaCode = Miyazaki
	case "Okinawa":
		areaCode = Okinawa
	default:
		http.Error(w, "Invalid ken", http.StatusBadRequest)
	}
	return areaCode
}