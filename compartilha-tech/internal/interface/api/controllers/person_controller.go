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
	mux.HandleFunc("GET /person", c.handleList)
	mux.HandleFunc("GET /person/{personID}", c.handleGetById)
	mux.HandleFunc("DELETE /person/{personID}", c.handleDelete)
	mux.HandleFunc("PATCH /person/{personID}", c.handleUpdate)

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

	responseJson(w, http.StatusCreated, person)
}

func (c *PersonController) handleList(w http.ResponseWriter, r *http.Request) {
	persons, err := c.service.List()
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseJson(w, http.StatusOK, persons)
}

func (c *PersonController) handleGetById(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("personID")

	person, err := c.service.GetById(ID)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if person == nil {
		fmt.Println("Person not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseJson(w, http.StatusOK, person)
}

func (c *PersonController) handleUpdate(w http.ResponseWriter, r *http.Request) {
	payload := dto.UpdatePerson{}

	ID := r.PathValue("personID")

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	person, err := c.service.Update(ID, payload)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseJson(w, http.StatusOK, person)
}

func (c *PersonController) handleDelete(w http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("personID")

	err := c.service.Delete(ID)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseJson(w, http.StatusNoContent, nil)
}

func responseJson(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(v)
}
