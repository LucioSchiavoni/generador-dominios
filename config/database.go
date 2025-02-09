package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// dsn := "root:123@tcp(localhost:3308)/test?parseTime=true"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a la base de datos: %v", err)
	}
	fmt.Println("✅Connected to database!!")
	DB = db
	return nil

}
