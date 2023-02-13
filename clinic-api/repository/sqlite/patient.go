package sqlite

import (
	"clinic-api/models"
	"database/sql"
	"fmt"
)

type patient struct {
	// todo logger
	db *sql.DB
}

func initPatient(db *sql.DB) *patient {
	return &patient{
		db: db,
	}
}

func (p *patient) CreatePatient(m *models.Patient) (id uint, err error) {
	fmt.Println(m)
	return 0, err
}

func (p *patient) GetPatient(id uint) (m *models.Patient, err error) {
	fmt.Println(id)
	return nil, err
}

func (p *patient) UpdatePatient(m *models.Patient) error {
	fmt.Println(m)
	return nil
}

func (p *patient) DeletePatient(id uint) error {
	fmt.Println(id)
	return nil
}
