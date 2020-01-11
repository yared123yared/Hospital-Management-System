package PetientService

import (
	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// AppointmentService implements menu.AppointmentService interface
type AppointmentService struct {
	appointmentRepo petient.AppointmentRepository
}

// NewAppointmentService  returns a new AppointmentService object
func NewAppointmentService(appointmentRepository petient.AppointmentRepository) petient.AppointmentService {
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
