package Appointment

import "github.com/getach1/web1/hospital/entity"

type AppointmentRepository interface {
	Appointments() ([]entity.Appointment, []error)
	Appointment(id uint) (*entity.Appointment, []error)
	UpdateAppointment(appointment *entity.Appointment) (*entity.Appointment, []error)
	DeleteAppointment(id uint) (*entity.Appointment, []error)
	StoreAppointment(appointment *entity.Appointment) (*entity.Appointment, []error)
}
