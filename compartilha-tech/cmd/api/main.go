package main

import (
	"compartilhatech/internal/application/services"
	"compartilhatech/internal/infra/database/sqlc"
	"compartilhatech/internal/interface/api/controllers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	dbConn := sqlc.NewDB()
	db, err := dbConn.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()

	personService := services.NewPersonService(db)

	controllers.NewPersonController(mux, personService)

	port := ":3333"

	fmt.Println("Starting server in port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Error:", err)
	}

}
