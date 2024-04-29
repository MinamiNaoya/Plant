package database

import (
	"database/sql"
	"log"
)


func SetupDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func selectFromTable(db *sql.DB, tableName string){
	var err error
	var stmt *sql.Stmt

	query := "SELECT * FROM " + tableName
	stmt, err = db.Prepare(query)
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()
	
}