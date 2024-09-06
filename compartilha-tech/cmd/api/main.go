package main

import (
	"compartilhatech/internal/application/services"
	"compartilhatech/internal/interface/api/controllers"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	personService := services.NewPersonService()

	controllers.NewPersonController(mux, personService)

	port := ":3333"

	fmt.Println("Starting server in port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Error:", err)
	}

}
