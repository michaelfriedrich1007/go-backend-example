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

func (s *PersonService) List() ([]entities.Person, error) {
	return s.PersonRepository, nil
}

func (s *PersonService) GetById(ID string) (*entities.Person, error) {
	person := new(entities.Person)

	for _, p := range s.PersonRepository {
		if p.ID == ID {
			person = &p
			return person, nil
		}
	}

	return nil, nil
}

func (s *PersonService) Update(ID string, data dto.UpdatePerson) (*entities.Person, error) {
	for i, p := range s.PersonRepository {
		if p.ID == ID {
			if data.Name != nil {
				p.Name = *data.Name
			}
			if data.Age != nil {
				p.Age = *data.Age
			}
			if data.Active != nil {
				p.Active = *data.Active
			}

			s.PersonRepository[i] = p

			return &p, nil
		}
	}

	return nil, fmt.Errorf("not fount")
}

func (s *PersonService) Delete(ID string) error {
	for i, p := range s.PersonRepository {
		if p.ID == ID {
			s.PersonRepository = append(s.PersonRepository[:i], s.PersonRepository[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("not fount")

}
