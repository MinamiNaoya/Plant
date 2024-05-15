package database

import (
	"fmt"
    "log"
    "database/sql"
	
    _ "github.com/mattn/go-sqlite3"

	"github.com/MinamiNaoya/Plant/database"
	"github.com/MinamiNaoya/Plant/common"
    
)

type PlantInfo struct {
	Name 			  string
	MinTemp           int
	MaxTemp           int
	Fertilizer        string
	Habitat           string
	WaterFreq     string
	SummerOrWinter    string
	PurchaseDate      string
}


/*植物のメタ情報を手動で入力*/
func inputPlantInfo() PlantInfo{
	var plant PlantInfo
	plant.Name = common.InputString("名前：")
	plant.MinTemp = common.InputInt("最低気温: ")
	plant.MaxTemp = common.InputInt("最高気温: ")
	plant.Fertilizer = common.InputString("肥料: ")
	plant.Habitat = common.InputString("生息地: ")
	plant.WaterFreq = common.InputString("水やり頻度: ")
	plant.SummerOrWinter = common.InputString("夏型または冬型: ")
	plant.PurchaseDate = common.InputString("植物の購入時期: ")
	
	return plant
}

func createPlantInfoTable(db *sql.DB) {
	var err error
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS plant_info (
		name TEXT,
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
}

func prepareInsertStatement(db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare("INSERT INTO plant_info (name, min_temp, max_temp, fertilizer, habitat, water_frequency, summer_or_winter, purchase_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	return stmt
}

func execStatement(stmt *sql.Stmt, plant PlantInfo){
	var err error
	_, err = stmt.Exec(plant.Name, plant.MinTemp, plant.MaxTemp, plant.Fertilizer, plant.Habitat, plant.WaterFreq, plant.SummerOrWinter, plant.PurchaseDate)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("植物の情報をデータベースに登録しました。")
}


func RegisterPlantInfo() {
	plant := inputPlantInfo()

	dbPath := "./plant.db"
	dbDriver := "sqlite3"
	db, _ := database.SetupDB(dbDriver, dbPath)
	defer db.Close()
	
	createPlantInfoTable(db)
	stmt := prepareInsertStatement(db)
	defer stmt.Close()
	execStatement(stmt, plant)
}