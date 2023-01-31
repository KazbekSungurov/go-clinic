package service

import (
	"clinic-api/models"
	"time"
)

type Service struct {
	// todo repository
	// todo logger
	User
	Employee
	Patient
	Department
}

// todo New Service

type User interface {
	Validate(m *models.User) error
	Create(m *models.User) (id uint, err error)
	Update(m *models.User) error
	Delete(id uint) error
	Get(id uint) (m *models.User, err error)
}

type Employee interface {
	Create(m *models.Employee) (id uint, err error)
	Update(m *models.Employee) error
	Delete(id uint) error
	Get(id uint) (m *models.Employee, err error)
}

type Patient interface {
	Validate(m *models.Patient) error
	Create(m *models.Patient) (id uint, err error)
	Update(m *models.Patient) error
	Delete(id uint) error
	Get(id uint) (m *models.Patient, err error)
}

type Department interface {
	Validate(m *models.Department) error
	Create(m *models.Department) (id uint, err error)
	Update(m *models.Department) error
	Delete(id uint) error
	Get(id uint) (m *models.Department, err error)
	AddEmployee(id, employeeId uint) error
	DeleteEmployee(id, employeeId uint) error
	ScheduleOperationBuild(m *models.Department, from, to time.Time) error
	ScheduleConsultBuild(m *models.Department, from, to time.Time) error
}
