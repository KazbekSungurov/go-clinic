package sqlite

import (
	"clinic-api/repository"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type rep struct {
	*department
	*employee
	*patient
	*user
}

func NewRepository(db *sql.DB) repository.Repository {
	return &rep{
		department: initDepartment(db),
		employee:   initEmployee(db),
		patient:    initPatient(db),
		user:       initUser(db),
	}
}

func NewDB(source string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
