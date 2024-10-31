package services

import (
	"compartilhatech/internal/application/dto"
	"compartilhatech/internal/application/service_interface"
	"compartilhatech/internal/domain/entities"
	"compartilhatech/internal/infra/database/sqlc/queries"
	"context"
	"database/sql"
	"time"
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
	dbConn := queries.New(s.db)

	p, err := dbConn.GetPersonById(context.Background(), ID)
	if err != nil {
		return nil, err
	}

	return &entities.Person{
		ID:        p.ID,
		Name:      p.Name,
		Age:       int(p.Age.Int32),
		Active:    p.Active,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (s *PersonService) Update(ID string, data dto.UpdatePerson) (*entities.Person, error) {
	dbConn := queries.New(s.db)

	err := dbConn.UpdatePerson(context.Background(), queries.UpdatePersonParams{
		ID:        ID,
		Name:      sql.NullString{String: *data.Name, Valid: data.Name != nil},
		Age:       sql.NullInt32{Int32: int32(*data.Age), Valid: data.Age != nil},
		Active:    sql.NullBool{Bool: *data.Active, Valid: data.Active != nil},
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return s.GetById(ID)
}

func (s *PersonService) Delete(ID string) error {
	dbConn := queries.New(s.db)

	err := dbConn.DeletePerson(context.Background(), ID)
	if err != nil {
		return err
	}

	return nil
}
