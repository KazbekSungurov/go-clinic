package service

import (
	"clinic-api/models"
	"clinic-api/repository"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type user struct {
	// todo logger
	repository.Repository
}

func initUserService(r repository.Repository) User {
	return &user{
		Repository: r,
	}
}

func (u *user) Validate(m *models.User) error {
	return validation.ValidateStruct(m,
		validation.Field(&m.FirstName, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(2, 30), validation.By(IsSQL)),
		validation.Field(&m.LastName, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(2, 30), validation.By(IsSQL)),
		validation.Field(&m.MiddleName, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(2, 30), validation.By(IsSQL)),
		validation.Field(&m.Gender, validation.Required),
		validation.Field(&m.BirthDate, validation.Required, validation.By(IsValidBirthDate)),
		validation.Field(&m.Details, validation.By(IsSQL)),
		validation.Field(&m.Phone, validation.By(IsSQL), validation.By(IsPhone)),
		validation.Field(&m.Address, validation.By(IsLetterHyphenSpaces), validation.Length(3, 100), validation.By(IsSQL)),
		validation.Field(&m.Polyclinic, validation.Length(3, 100), validation.By(IsSQL)),
		validation.Field(&m.Email, validation.Required, is.Email, validation.By(IsSQL)),
	)
}

func (u *user) Create(m *models.User) (id uint, err error) {
	if err := u.Validate(m); err != nil {
		return 0, err
	}

	return u.Repository.CreateUser(m)
}

func (u *user) Update(m *models.User) error {
	if err := u.Validate(m); err != nil {
		return err
	}

	return u.Repository.UpdateUser(m)
}

func (u *user) Delete(id uint) error {
	return u.Repository.DeleteUser(id)
}

func (u *user) Get(id uint) (m *models.User, err error) {
	return u.Repository.GetUser(id)
}
