package Service

import (
	appointment "github.com/getach1/web1/hospital/Appointment"
	"github.com/getach1/web1/hospital/entity"
)

// AppointmentService implements menu.AppointmentService interface
type AppointmentService struct {
	appointmentRepo appointment.AppointmentRepository
}

// NewAppointmentService  returns a new AppointmentService object
func NewAppointmentService(appointmentRepository appointment.AppointmentRepository) appointment.AppointmentService {
	return &AppointmentService{appointmentRepo: appointmentRepository}
}

// Appointments returns all stored application appointments
func (us *AppointmentService) Appointments() ([]entity.Appointment, []error) {
	usrs, errs := us.appointmentRepo.Appointments()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// Appointment retrieves an application appointment by its id
func (us *AppointmentService) Appointment(id uint) (*entity.Appointment, []error) {
	usr, errs := us.appointmentRepo.Appointment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UpdateAppointment updates  a given application appointment
func (us *AppointmentService) UpdateAppointment(appointment *entity.Appointment) (*entity.Appointment, []error) {
	usr, errs := us.appointmentRepo.UpdateAppointment(appointment)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteAppointment deletes a given application appointment
func (us *AppointmentService) DeleteAppointment(id uint) (*entity.Appointment, []error) {
	usr, errs := us.appointmentRepo.DeleteAppointment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreAppointment stores a given application appointment
func (us *AppointmentService) StoreAppointment(appointment *entity.Appointment) (*entity.Appointment, []error) {
	usr, errs := us.appointmentRepo.StoreAppointment(appointment)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
