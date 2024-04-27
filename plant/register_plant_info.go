package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type PlantInfo struct {
	MinTemp           int
	MaxTemp           int
	Fertilizer        string
	Habitat           string
	WaterFreq     string
	SummerOrWinter    string
	PurchaseDate      string
}
type Config struct{
	Database DatabaseConfig
	Server ServerConfig
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
// YAMLデータをマッピングするための構造体
type DatabaseConfig struct {
	User string `yaml:"mysqluser"`
	Password string `yaml:"password"`
	Name string `yaml:"name"`
}

func inputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func inputInt(prompt string) int {
	var input int
	fmt.Print(prompt)
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	return input
}


func inputPlantInfo() PlantInfo{
	var plant PlantInfo
	plant.MinTemp = inputInt("最低気温: ")
	plant.MaxTemp = inputInt("最高気温: ")
	plant.Fertilizer = inputString("肥料: ")
	plant.Habitat = inputString("生息地: ")
	plant.WaterFreq = inputString("水やり頻度: ")
	plant.SummerOrWinter = inputString("夏型または冬型: ")
	plant.PurchaseDate = inputString("植物の購入時期: ")
	
	return plant
}



func main() {
	plant := inputPlantInfo()

	dbPath := "./plant.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS plant_info (
		min_temp INT,
		max_temp INT,
		fertilizer TEXT,
		habitat TEXT,
		water_frequency TEXT,
		summer_or_winter TEXT,
		purchase_date TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}
	
	stmt, err := db.Prepare("INSERT INTO plant_info (min_temp, max_temp, fertilizer, habitat, water_frequency, summer_or_winter, purchase_date) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	

	// クエリの実行
	_, err = stmt.Exec(plant.MinTemp, plant.MaxTemp, plant.Fertilizer, plant.Habitat, plant.WaterFreq, plant.SummerOrWinter, plant.PurchaseDate)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("植物の情報をデータベースに登録しました。")
}