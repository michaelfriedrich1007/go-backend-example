package services

import (
	"compartilhatech/internal/application/dto"
	"compartilhatech/internal/application/service_interface"
	"compartilhatech/internal/domain/entities"
	"compartilhatech/internal/infra/database/sqlc/queries"
	"context"
	"database/sql"
	"fmt"
)

type PersonService struct {
	db               *sql.DB
	PersonRepository []entities.Person
}

func NewPersonService(db *sql.DB) service_interface.PersonService {
	return &PersonService{
		db:               db,
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

	//s.PersonRepository = append(s.PersonRepository, *p)

	dbConn := queries.New(s.db)

	err := dbConn.InsertPerson(context.Background(), queries.InsertPersonParams{
		ID:        p.ID,
		Name:      p.Name,
		Age:       sql.NullInt32{Int32: int32(p.Age), Valid: true},
		Active:    p.Active,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	return p, nil

}

func (s *PersonService) List() ([]entities.Person, error) {
	dbConn := queries.New(s.db)

	p, err := dbConn.GetPersons(context.Background())
	if err != nil {
		return nil, err
	}

	persons := []entities.Person{}
	for _, v := range p {
		persons = append(persons, entities.Person{
			ID:        v.ID,
			Name:      v.Name,
			Age:       int(v.Age.Int32),
			Active:    v.Active,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return persons, nil
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
