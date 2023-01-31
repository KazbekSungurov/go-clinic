package models

import "time"

const (
	Male   Gender = "муж"
	Female Gender = "жен"
)

type Gender string

type PersonalData struct {
	Phone      string `json:"phone,omitempty"`
	Email      string `json:"email,omitempty"`
	Address    string `json:"address,omitempty"`
	Polyclinic string `json:"polyclinic,omitempty"`
}

type User struct {
	ID           uint      `json:"id,omitempty"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	MiddleName   string    `json:"middle_name,omitempty"`
	Gender       Gender    `json:"gender,omitempty"`
	BirthDate    time.Time `json:"birth_date"`
	Details      string    `json:"details,omitempty"`
	PersonalData `json:"personal_data"`
}

type Patient struct {
	ID        uint   `json:"id,omitempty"`
	Diagnosis string `json:"diagnosis,omitempty"`
	User      `json:"user"`
}

type Employee struct {
	ID   uint `json:"id,omitempty"`
	User `json:"user"`
}

type ScheduleDetails struct {
	WeekDays            []time.Weekday `json:"week_days,omitempty"`
	StartAt             time.Time      `json:"start_at"`
	EndAt               time.Time      `json:"end_at"`
	ExaminationDuration time.Duration  `json:"examination_duration,omitempty"`
	BreakStart          time.Time      `json:"break_start"`
	BreakDuration       time.Duration  `json:"break_duration,omitempty"`
}

type Appointment struct {
	Day      time.Time `json:"day"`
	Patients []Patient `json:"patients,omitempty"`
}

type ScheduleOperation struct {
	ID              uint          `json:"id,omitempty"`
	Appointments    []Appointment `json:"appointments,omitempty"`
	ScheduleDetails `json:"schedule_details"`
}

type ScheduleConsult struct {
	ID              uint          `json:"id,omitempty"`
	Appointments    []Appointment `json:"appointments,omitempty"`
	ScheduleDetails `json:"schedule_details"`
}

type Department struct {
	ID                uint       `json:"id,omitempty"`
	Name              string     `json:"name,omitempty"`
	Employees         []Employee `json:"employees,omitempty"`
	ScheduleConsult   `json:"schedule_consult"`
	ScheduleOperation `json:"schedule_operation"`
}
