package service

import (
	"clinic-api/models"
	"clinic-api/repository"
)

type employee struct {
	User
	repository.Repository
	// todo logger
}

func initEmployeeService(r repository.Repository, u User) Employee {
	return &employee{
		User:       u,
		Repository: r,
	}
}

func (e *employee) Create(m *models.Employee) (id uint, err error) {
	if _, err := e.User.Create(&m.User); err != nil {
		return 0, nil
	}

	return e.Repository.CreateEmployee(m)
}

func (e *employee) Update(m *models.Employee) error {
	if err := e.User.Update(&m.User); err != nil {
		return err
	}

	return e.Repository.UpdateEmployee(m)
}

func (e *employee) Get(id uint) (m *models.Employee, err error) {
	return e.Repository.GetEmployee(id)
}

func (e *employee) Delete(id uint) error {
	return e.Repository.DeleteEmployee(id)
}
