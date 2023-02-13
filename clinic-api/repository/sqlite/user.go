package sqlite

import (
	"clinic-api/models"
	"database/sql"
	"fmt"
)

type user struct {
	// todo logger
	db *sql.DB
}

func initUser(db *sql.DB) *user {
	return &user{
		db: db,
	}
}

func (u *user) CreateUser(m *models.User) (id uint, err error) {
	fmt.Println(m)
	return 0, err
}

func (u *user) GetUser(id uint) (m *models.User, err error) {
	fmt.Println(id)
	return nil, err
}

func (u *user) UpdateUser(m *models.User) error {
	fmt.Println(m)
	return nil
}

func (u *user) DeleteUser(id uint) error {
	fmt.Println(id)
	return nil
}
