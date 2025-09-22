package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"

	_ "github.com/lib/pq" // PostgreSQL driver
)
func ConnectDB() (*sql.DB, error) {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: could not read .env file: %v", err)
	}

	viper.AutomaticEnv() 

	connStr := viper.GetString("DATABASE_URL")
	if connStr == "" {
		log.Fatalf("DATABASE_URL not set in environment or .env file")
		return nil, fmt.Errorf("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
		return nil, err
	}

	fmt.Println("Database connection established successfully")
	return db, nil
}


func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing the database: %v", err)
	} else {
		fmt.Println("Database connection closed successfully")
	}
}