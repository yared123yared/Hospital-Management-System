package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	//"github.com/yaredsolomon/webProgram1/hospital/request"

	"github.com/web1_group_project/hospital_server/Doctor"
	"github.com/web1_group_project/hospital_server/entity"
)

// AppointmentGormRepo Implements the request.AppointmentRepository interface
type MockAppointmentGormRepo struct {
	conn *gorm.DB
}

// NewAppointmentGormRepo creates a new object of AppointmentGormRepo
func NewMockAppointmentGormRepo(db *gorm.DB) Doctor.AppointmentRepository {
	return &MockAppointmentGormRepo{conn: db}
}

// Appointments return all Appointments from the database
func (appointRepo *MockAppointmentGormRepo) Appointments() ([]entity.Doctor, []error) {
	appointments := []entity.Doctor{entity.DoctorMock}
	return appointments, nil
}

// Appointment retrieves a Appointment by its id from the database
func (appointRepo *MockAppointmentGormRepo) Appointment(id uint) (*entity.Doctor, []error) {
	appointment := entity.Doctor{}
	if id == 1 {
		return &appointment, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateAppointment updates a given Appointment in the database
func (appointRepo *MockAppointmentGormRepo) UpdateAppointment(appointment *entity.Doctor) (*entity.Doctor, []error) {
	prec := entity.DoctorMock
	return &prec, nil
}

// DeleteAppointment deletes a given appointment from the database
func (appointRepo *MockAppointmentGormRepo) DeleteAppointment(id uint) (*entity.Doctor, []error) {
	prec := entity.DoctorMock
	return &prec, nil
}
