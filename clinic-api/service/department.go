package service

import (
	"clinic-api/models"
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

func (d *department) ScheduleOperationBuild(m *models.Department, from, to time.Time) error {
	if from.IsZero() || to.IsZero() {
		return errTimeIsNil
	}

	if from.After(to) {
		from, to = to, from
	}

	for i := from; i.Before(to); i.AddDate(0, 0, 1) {
		if slices.Contains[time.Weekday](m.WeekDays, i.Weekday()) {
			m.Appointments = append(m.Appointments, models.Appointment{Day: i})
			break
		}
	}

	return nil
}

//func (d *department) ScheduleConsultBuild(m *models.Department, from, to time.Time) error {
//	if from.IsZero() || to.IsZero() {
//		return errTimeIsNil
//	}
//
//	if from.After(to) {
//		from, to = to, from
//	}
//
//	for i := from; i.Before(to); i.AddDate(0, 0, 1) {
//		if slices.Contains[time.Weekday](m.WeekDays, i.Weekday()) {
//			var as []models.Appointment
//			m.ScheduleDetails.StartAt
//			m.ScheduleDetails.EndAt
//			m.ScheduleDetails.BreakStart
//			m.ScheduleDetails.BreakDuration
//			m.ScheduleDetails.ExaminationDuration
//			as = append(as)
//			m.Appointments = append(m.Appointments, models.Appointment{Day: i})
//			break
//		}
//	}
//
//	return nil
//}
