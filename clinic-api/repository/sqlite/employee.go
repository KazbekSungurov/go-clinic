package sqlite

import (
	"clinic-api/models"
	"database/sql"
	"fmt"
)

type employee struct {
	// todo logger
	db *sql.DB
}

func initEmployee(db *sql.DB) *employee {
	return &employee{
		db: db,
	}
}

func (e *employee) CreateEmployee(m *models.Employee) (id uint, err error) {
	fmt.Println(m)
	return 0, err
}

func (e *employee) GetEmployee(id uint) (m *models.Employee, err error) {
	fmt.Println(id)
	return nil, err
}

func (e *employee) UpdateEmployee(m *models.Employee) error {
	fmt.Println(m)
	return nil
}

func (e *employee) DeleteEmployee(id uint) error {
	fmt.Println(id)
	return nil
}
