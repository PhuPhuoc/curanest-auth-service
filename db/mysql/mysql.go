package mysql

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sqlx.DB, string) {
	err := godotenv.Load()
	if err != nil {
		log.Println("(dev) Error loading .env file")
	}
	appport := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=True&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Println(dsn)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("connect db: ", err)
	}
	return db, appport
}
