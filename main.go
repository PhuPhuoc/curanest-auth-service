package main

import (
	"log"

	"github.com/PhuPhuoc/curanest-auth-service/api"
	"github.com/PhuPhuoc/curanest-auth-service/db/mysql"
)

// @title		Authentication Service
// @version	1.0
func main() {
	db, port := mysql.ConnectDB()
	if err_ping := db.Ping(); err_ping != nil {
		log.Println("Cannot ping db: ", err_ping)
	}
	defer db.Close()

	server := api.InitServer(port, db)
	if err_run_server := server.RunApp(); err_run_server != nil {
		log.Fatal("Cannot run app: ", err_run_server)
	}
}
