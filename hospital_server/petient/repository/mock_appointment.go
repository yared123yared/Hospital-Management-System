package PetientRepository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// AppointmentGormRepo Implements the menu.AppointmentRepository interface
type MockAppointmentGormRepo struct {
	conn *gorm.DB
}

// NewAppointmentGormRepo creates a new object of AppointmentGormRepo
func NewMockAppointmentGormRepo(db *gorm.DB) petient.AppointmentRepository {
	return &MockAppointmentGormRepo{conn: db}
}

// Appointments return all appointments from the database
func (appointmentRepo *MockAppointmentGormRepo) Appointments() ([]entity.Appointment, []error) {
	appointments := []entity.Appointment{entity.AppointmentMock}
	return appointments, nil
}

// Appointment retrieves a appointment by its id from the database
func (appointmentRepo *MockAppointmentGormRepo) Appointment(id uint) (*entity.Appointment, []error) {
	appointment := entity.AppointmentMock

	if id == 1 {
		return &appointment, nil
	}
	return nil, []error{errors.New("Not found")}
}
