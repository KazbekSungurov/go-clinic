package service

import (
	"clinic-api/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

type patient struct {
	User
	// todo logger
}

func newPatientService(u User) error {
	//todo
	return nil
}

func (p *patient) Validate(m *models.Patient) error {
	if err := p.User.Validate(&m.User); err != nil {
		return err
	}

	return validation.ValidateStruct(m,
		validation.Field(&m.Diagnosis, validation.Length(5, 300), validation.By(IsSQL)),
	)
}

func (p *patient) Create(m *models.Patient) (id uint, err error) {
	if err := p.Validate(m); err != nil {
		return 0, err
	}

	panic("implement me")
}

func (p *patient) Update(m *models.Patient) error {
	if err := p.Validate(m); err != nil {
		return err
	}

	panic("implement me")
}

func (p *patient) Get(id uint) (m *models.Patient, err error) {
	//TODO implement me
	panic("implement me")
}
