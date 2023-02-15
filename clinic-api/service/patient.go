package service

import (
	"clinic-api/models"
	"clinic-api/repository"
	validation "github.com/go-ozzo/ozzo-validation"
)

type patient struct {
	User
	repository.Repository
	// todo logger
}

func initPatientService(r repository.Repository, u User) Patient {
	return &patient{
		User:       u,
		Repository: r,
	}
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

	return p.Repository.CreatePatient(m)
}

func (p *patient) Update(m *models.Patient) error {
	if err := p.Validate(m); err != nil {
		return err
	}

	return p.Repository.UpdatePatient(m)
}

func (p *patient) Get(id uint) (m *models.Patient, err error) {
	return p.Repository.GetPatient(id)
}

func (p *patient) Delete(id uint) error {
	return p.Repository.DeletePatient(id)
}
