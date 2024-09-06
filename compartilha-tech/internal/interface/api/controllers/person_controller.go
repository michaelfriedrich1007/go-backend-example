package controllers

import (
	"compartilhatech/internal/application/dto"
	"compartilhatech/internal/application/service_interface"
	"encoding/json"
	"fmt"
	"net/http"
)

type PersonController struct {
	service service_interface.PersonService
}

func NewPersonController(mux *http.ServeMux, s service_interface.PersonService) *PersonController {
	c := PersonController{
		service: s,
	}

	mux.HandleFunc("POST /person", c.handleCreate)

	return &c
}

func (c *PersonController) handleCreate(w http.ResponseWriter, r *http.Request) {
	payload := dto.CreatePerson{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	person, err := c.service.Insert(payload)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(person)

	fmt.Println(payload)
}
