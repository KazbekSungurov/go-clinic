package service

import (
	"clinic-api/models"
)

type employee struct {
	User
	// todo logger
}

func newEmployeeService(u User) error {
	// todo
	return nil
}

func (e *employee) Create(m *models.Employee) (id uint, err error) {
	if _, err := e.User.Create(&m.User); err != nil {
		return 0, nil
	}

	//TODO implement me
	panic("implement me")
}

func (e *employee) Update(m *models.Employee) error {
	if err := e.User.Update(&m.User); err != nil {
		return err
	}

	//TODO implement me
	panic("implement me")
}

func (e *employee) Get(id uint) (m *models.Employee, err error) {
	//TODO implement me
	panic("implement me")
}
