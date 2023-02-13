package repository

import (
	"clinic-api/models"
)

type Repository interface {
	CreateUser(m *models.User) (id uint, err error)
	GetUser(id uint) (m *models.User, err error)
	UpdateUser(m *models.User) error
	DeleteUser(id uint) error

	CreateEmployee(m *models.Employee) (id uint, err error)
	GetEmployee(id uint) (m *models.Employee, err error)
	UpdateEmployee(m *models.Employee) error
	DeleteEmployee(id uint) error

	CreatePatient(m *models.Patient) (id uint, err error)
	GetPatient(id uint) (m *models.Patient, err error)
	UpdatePatient(m *models.Patient) error
	DeletePatient(id uint) error

	CreateDepartment(m *models.Department) (id uint, err error)
	GetDepartment(id uint) (m *models.Department, err error)
	UpdateDepartment(m *models.Department) error
	DeleteDepartment(id uint) error
}
