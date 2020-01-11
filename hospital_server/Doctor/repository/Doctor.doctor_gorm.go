package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/web1_group_project/hospital_server/Doctor"
	"github.com/web1_group_project/hospital_server/entity"
	//"github.com/yaredsolomon/webProgram1/hospital/request"
)

// AppointmentGormRepo Implements the request.AppointmentRepository interface
type AppointmentGormRepo struct {
	conn *gorm.DB
}

// NewAppointmentGormRepo creates a new object of AppointmentGormRepo
func NewAppointmentGormRepo(db *gorm.DB) Doctor.AppointmentRepository {
	return &AppointmentGormRepo{conn: db}
}

// Appointments return all Appointments from the database
func (appointRepo *AppointmentGormRepo) Appointments() ([]entity.Doctor, []error) {
	appointments := []entity.Doctor{}
	errs := appointRepo.conn.Preload("Profile").Preload("Prescription").Preload("Diagnosis").Preload("Appointment").Preload("Pharmacist").Find(&appointments).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return appointments, errs
}

// Appointment retrieves a Appointment by its id from the database
func (appointRepo *AppointmentGormRepo) Appointment(id uint) (*entity.Doctor, []error) {
	appointment := entity.Doctor{}
	errs := appointRepo.conn.Preload("Profile").Preload("Prescription").Preload("Diagnosis").Preload("Appointment").First(&appointment, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &appointment, errs
}

// UpdateAppointment updates a given Appointment in the database
func (appointRepo *AppointmentGormRepo) UpdateAppointment(appointment *entity.Doctor) (*entity.Doctor, []error) {
	fmt.Println("i am at the update method")
	apt := appointment
	fmt.Println("thise is the data that will be updated")
	fmt.Println(apt)

	errs := appointRepo.conn.Save(apt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	fmt.Println("i have done witht the updates")
	return apt, errs
}

// DeleteAppointment deletes a given appointment from the database
func (appointRepo *AppointmentGormRepo) DeleteAppointment(id uint) (*entity.Doctor, []error) {
	apt, errs := appointRepo.Appointment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = appointRepo.conn.Delete(apt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return apt, errs
}
