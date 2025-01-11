package main

import (
	"log"

	"github.com/PhuPhuoc/curanest-auth-service/api"
)

// @title		Authentication Service
// @version	1.0
func main() {
	// port := ":8999"
	// db := mysql.ConnectDB()
	// if err_ping := db.Ping(); err_ping != nil {
	// 	log.Println("Cannot ping db: ", err_ping)
	// }
	// defer db.Close()

	server := api.InitServerTemp(":8001")

	if err_run_server := server.RunApp(); err_run_server != nil {
		log.Fatal("Cannot run app: ", err_run_server)
	}
}
