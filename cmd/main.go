package main

import (
	"fmt"
	"os"
	"umkmgubernur/internal/db"
)

func main() {
 dbConn, err := db.ConnectDB()
 if err != nil {
  fmt.Printf("Failed to connect to database: %v\n", err)
  os.Exit(1)
 }
 defer db.CloseDB(dbConn)
 fmt.Println("server running and DB connected!")
}