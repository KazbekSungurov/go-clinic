package service

import (
	"clinic-api/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

type department struct {
	Employee
	// todo logger
}

func newDepartmentService(e Employee) error {
	// todo
	return nil
}

func (d *department) Validate(m *models.Department) error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Name, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(2, 100), validation.By(IsSQL)),
	)
}

func (d *department) Create(m *models.Department) (id uint, err error) {
	//TODO implement me
	panic("implement me")
}

func (d *department) Update(m *models.Department) error {
	//TODO implement me
	panic("implement me")
}

func (d *department) Get(id uint) (m *models.Department, err error) {
	//TODO implement me
	panic("implement me")
}
