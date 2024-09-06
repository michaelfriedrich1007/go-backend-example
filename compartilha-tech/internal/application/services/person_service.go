package services

import (
	"compartilhatech/internal/application/dto"
	"compartilhatech/internal/application/service_interface"
	"compartilhatech/internal/domain/entities"
	"fmt"
)

type PersonService struct {
	PersonRepository []entities.Person
}

func NewPersonService() service_interface.PersonService {
	return &PersonService{
		PersonRepository: []entities.Person{},
	}
}

func (s *PersonService) Insert(data dto.CreatePerson) (*entities.Person, error) {
	p := entities.NewPerson()
	p.Name = data.Name
	p.Age = data.Age

	if data.Active != nil {
		p.Active = *data.Active
	}

	s.PersonRepository = append(s.PersonRepository, *p)

	return p, nil

}

func (s *PersonService) List() ([]*entities.Person, error) {
	fmt.Println("List")
	return nil, nil
}

func (s *PersonService) GetById(ID string) (*entities.Person, error) {
	fmt.Println("GetById")
	return nil, nil
}
