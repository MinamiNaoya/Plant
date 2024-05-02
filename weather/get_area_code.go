package weather

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAreaCode(w http.ResponseWriter, r *http.Request) (areaCode AREACODE){
	var requestData struct {
		Ken string `json:"ken"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	switch requestData.Ken {
	case "北海道":
		areaCode = Hokkaido
	case "青森県": 
		areaCode = Aomori
	case "秋田県":
		areaCode = Akita
	case "山形県":
		areaCode = Yamagata
	case "岩手県":
		areaCode = Iwate
	case "宮城県":
		areaCode = Miyagi
	case "福島県":
		areaCode = Fukushima
	case "茨城県":
		areaCode = Ibaraki
	case "埼玉県":
		areaCode = Saitama
	case "栃木県":
		areaCode = Tochigi
	case "群馬県":
		areaCode = Gunma
	case "神奈川県":
		areaCode = Kanagawa
	case "千葉県":
		areaCode = Chiba
	case "東京都":
		areaCode = Tokyo
	case "新潟県":
		areaCode = Nigata
	case "富山県":
		areaCode = Toyama
	case "石川県":
		areaCode = Ishikawa
	case "福井県":
		areaCode = Fukui
	case "山梨県":
		areaCode = Yamanashi
	case "長野県":
		areaCode = Nagano
	case "岐阜県":
		areaCode = Gifu
	case "静岡県":
		areaCode = Shizuoka
	case "愛知県":
		areaCode = Aichi
	case "三重県":
		areaCode = Mie
	case "滋賀県":
		areaCode = Shiga
	case "京都府":
		areaCode = Kyoto
	case "大阪府":
		areaCode = Osaka
	case "兵庫県":
		areaCode = Hyogo
	case "奈良県":
		areaCode = Nara
	case "和歌山県":
		areaCode = Wakayama
	case "鳥取県":
		areaCode = Tottori
	case "島根県":
		areaCode = Shimane
	case "岡山県":
		areaCode = Okayama
	case "広島県":
		areaCode = Hiroshima
	case "山口県":
		areaCode = Yamaguchi
	case "徳島県":
		areaCode = Tokushima
	case "香川県":
		areaCode = Kagawa
	case "愛媛県":
		areaCode = Ehime
	case "高知県":
		areaCode = Kochi
	case "福岡県":
		areaCode = Fukuoka
	case "佐賀県":
		areaCode = Saga
	case "長崎県":
		areaCode = Nagasaki
	case "熊本県":
		areaCode = Kumamoto
	case "大分県":
		areaCode = Oita
	case "宮崎県":
		areaCode = Miyazaki
	case "沖縄県":
		areaCode = Okinawa
	default:
		http.Error(w, "Invalid ken", http.StatusBadRequest)
	}
	return areaCode
}