package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/*
Return the database connection
*/
func GetDbConn() (*sql.DB, error) {
	dbHostname := "db" // Utilisez le nom d'hôfvfvte 'db' au lieu de 'localhost'
	_dbHostname, isSet := os.LookupEnv("DB_HOSTNAME")
	if isSet {
		dbHostname = _dbHostname
	}
	db, err := sql.Open("mysql", fmt.Sprintf("tomyj:tomyj123@tcp(%s:3306)/db", dbHostname))
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	return db, nil
}
