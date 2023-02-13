package sqlite

import (
	"clinic-api/models"
	"database/sql"
	"fmt"
)

type department struct {
	// todo logger
	db *sql.DB
}

func initDepartment(db *sql.DB) *department {
	return &department{
		db: db,
	}
}

func (d *department) CreateDepartment(m *models.Department) (id uint, err error) {
	fmt.Println(m)
	return 0, err
}

func (d *department) GetDepartment(id uint) (m *models.Department, err error) {
	fmt.Println(id)
	return nil, err
}

func (d *department) UpdateDepartment(m *models.Department) error {
	fmt.Println(m)
	return nil
}

func (d *department) DeleteDepartment(id uint) error {
	fmt.Println(id)
	return nil
}
