package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	
	// Connection string parameters
	server := "localhost:9042"
	
	database := "AuthorsDb"

	connString := fmt.Sprintf("server=%s;database=%s;integrated security=false;encrypt=disable;connection+timeout=30", server, database)

	// Create a SQL Server connection
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error opening connection:", err.Error())
		return
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err.Error())
		return
	}

	fmt.Println("Connected to the database successfully")
}

func GetDB() *gorm.DB {
	return db
}
