package Repository

import (
	"github.com/getach1/web1/hospital/Appointment"
	"github.com/getach1/web1/hospital/entity"
	"github.com/jinzhu/gorm"
)

// AppointmentGormRepo Implements the menu.AppointmentRepository interface
type AppointmentGormRepo struct {
	conn *gorm.DB
}

// NewAppointmentGormRepo creates a new object of AppointmentGormRepo
func NewAppointmentGormRepo(db *gorm.DB) Appointment.AppointmentRepository {
	return &AppointmentGormRepo{conn: db}
}

// Appointments return all appointments from the database
func (appointmentRepo *AppointmentGormRepo) Appointments() ([]entity.Appointment, []error) {
	appointments := []entity.Appointment{}
	errs := appointmentRepo.conn.Find(&appointments).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return appointments, errs
}

// Appointment retrieves a appointment by its id from the database
func (appointmentRepo *AppointmentGormRepo) Appointment(id uint) (*entity.Appointment, []error) {
	appointment := entity.Appointment{}

	errs := appointmentRepo.conn.Debug().Find(&appointment, "id=?", id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &appointment, errs
}

// UpdateAppointment updates a given appointment in the database
func (appointmentRepo *AppointmentGormRepo) UpdateAppointment(appointment *entity.Appointment) (*entity.Appointment, []error) {
	usr := appointment
	errs := appointmentRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteAppointment deletes a given appointment from the database
func (appointmentRepo *AppointmentGormRepo) DeleteAppointment(id uint) (*entity.Appointment, []error) {
	usr, errs := appointmentRepo.Appointment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = appointmentRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreAppointment stores a new appointment into the database
func (appointmentRepo *AppointmentGormRepo) StoreAppointment(appointment *entity.Appointment) (*entity.Appointment, []error) {
	usr := appointment
	errs := appointmentRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
