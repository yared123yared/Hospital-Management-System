package service

import (
	//"github.com/yaredsolomon/webProgram1/hospital/entity"
	//"github.com/yaredsolomon/webProgram1/hospital/request"

	//"github.com/yaredsolomon/webProgram1/hospital/request"
	"github.com/web1_group_project/hospital_server/Doctor"
	"github.com/web1_group_project/hospital_server/entity"
)

// AppointmentService implements request.AppointmentService interface
type AppointmentService struct {
	appointRepo Doctor.AppointmentRepository
}

// NewAppointmentService  returns a new AppointmentService object
func NewAppointmentService(appointmentRepository Doctor.AppointmentRepository) Doctor.AppointmentService {
	return &AppointmentService{appointRepo: appointmentRepository}
}

// Appointments returns all stored application Appointments
func (as *AppointmentService) Appointments() ([]entity.Doctor, []error) {
	appointments, errs := as.appointRepo.Appointments()
	if len(errs) > 0 {
		return nil, errs
	}
	return appointments, errs
}

// Appointment retrieves an application Appointment by its id
func (as *AppointmentService) Appointment(id uint) (*entity.Doctor, []error) {
	apt, errs := as.appointRepo.Appointment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return apt, errs
}

// UpdateAppointment updates  a given application appointment
func (as *AppointmentService) UpdateAppointment(appointment *entity.Doctor) (*entity.Doctor, []error) {
	apt, errs := as.appointRepo.UpdateAppointment(appointment)
	if len(errs) > 0 {
		return nil, errs
	}
	return apt, errs
}

// DeleteAppointment deletes a given application appointment
func (as *AppointmentService) DeleteAppointment(id uint) (*entity.Doctor, []error) {
	apt, errs := as.appointRepo.DeleteAppointment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return apt, errs
}
