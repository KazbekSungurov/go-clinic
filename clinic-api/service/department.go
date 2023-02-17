package service

import (
	"clinic-api/models"
	"clinic-api/repository"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/exp/slices"
	"time"
)

var (
	errTimeIsNil = errors.New("заданные промежутки времени некорректны")
)

type department struct {
	Employee
	repository.Repository
	// todo logger
}

func initDepartmentService(r repository.Repository, e Employee) Department {
	return &department{
		Employee:   e,
		Repository: r,
	}
}

func (d *department) Validate(m *models.Department) error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Name, validation.Required, validation.By(IsLetterHyphenSpaces), validation.Length(2, 100), validation.By(IsSQL)),
	)
}

func (d *department) Create(m *models.Department) (id uint, err error) {
	return d.Repository.CreateDepartment(m)
}

func (d *department) Update(m *models.Department) error {
	return d.Repository.UpdateDepartment(m)
}

func (d *department) Get(id uint) (m *models.Department, err error) {
	return d.Repository.GetDepartment(id)
}

func (d *department) Delete(id uint) error {
	// todo employeeId
	return d.Repository.DeleteDepartment(id)
}

func (d *department) AddEmployee(id, employeeId uint) error {
	// todo
	panic("")
}

func (d *department) DeleteEmployee(id, employeeId uint) error {
	// todo employeeId
	panic("implement me")
}

func (d *department) ScheduleOperationBuild(m *models.Department, from, to time.Time) error {
	if from.IsZero() || to.IsZero() {
		return errTimeIsNil
	}

	if from.After(to) {
		from, to = to, from
	}

	for i := from; i.Before(to.AddDate(0, 0, 1)); i = i.AddDate(0, 0, 1) {
		if slices.Contains[time.Weekday](m.ScheduleOperation.WeekDays, i.Weekday()) {
			m.ScheduleOperation.Appointments = append(m.ScheduleOperation.Appointments, models.Appointment{Day: i})
		}
	}

	return nil
}

func (d *department) ScheduleConsultBuild(m *models.Department, from, to time.Time) error {
	if from.IsZero() || to.IsZero() {
		return errTimeIsNil
	}

	if from.After(to) {
		from, to = to, from
	}

	for i := from; i.Before(to.AddDate(0, 0, 1)); i = i.AddDate(0, 0, 1) {
		if slices.Contains[time.Weekday](m.ScheduleConsult.WeekDays, i.Weekday()) {
			for x := m.ScheduleConsult.StartAt; x.Before(m.ScheduleConsult.EndAt.Add(time.Minute - m.ScheduleConsult.ExaminationDuration)); x = x.Add(m.ScheduleConsult.ExaminationDuration) {
				if x.After(m.ScheduleConsult.BreakStart) && x.Before(m.ScheduleConsult.BreakStart.Add(m.ScheduleConsult.BreakDuration)) {
					continue
				}
				m.ScheduleConsult.Appointments = append(m.ScheduleConsult.Appointments, models.Appointment{Day: time.Date(i.Year(), i.Month(), i.Day(), x.Hour(), x.Minute(), 0, 0, time.Local)})
			}
		}
	}

	return nil
}
