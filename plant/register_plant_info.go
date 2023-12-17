package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
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
	yamlFile, err := os.ReadFile("plant/config.yml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return
	}
	plant := inputPlantInfo()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.Database.User,
		config.Database.Password,
		config.Server.Host,
		config.Server.Port,
	)
	
	// データベースに接続
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// データベースの作成
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS plant")
	if err != nil {
		fmt.Println("Error creating database:", err)
		return
	}
	// データベースを選択
	_, err = db.Exec("USE plant")
	if err != nil {
		log.Fatal(err)
	}
	// テーブルの作成
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS plant_info (
		min_temp INT,
		max_temp INT,
		fertilizer VARCHAR(255),
		habitat VARCHAR(255),
		water_frequency VARCHAR(255),
		summer_or_winter VARCHAR(255),
		purchase_date DATE
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
	stmt.Exec(plant.MinTemp, plant.MaxTemp, plant.Fertilizer, plant.Habitat, plant.WaterFreq, plant.SummerOrWinter, plant.PurchaseDate)

	fmt.Println("植物の情報をデータベースに登録しました。")
}