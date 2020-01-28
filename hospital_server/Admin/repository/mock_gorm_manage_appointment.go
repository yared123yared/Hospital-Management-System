package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageAppointmetRepository implements Admin.ManageAppointmetRepository interface
type MockManageAppointmetRepository struct {
	conn *gorm.DB
}

// NewManageAppointmetRepository returns new object of ManageAppointmetRepository
func NewMockManageAppointmetRepository(db *gorm.DB) Admin.ManageAppointmetRepository {
	return &MockManageAppointmetRepository{conn: db}
}

// Appointments return all Appointments stored in the databasee
func (maRepo *MockManageAppointmetRepository) Appointments() ([]entity.Appointment, []error) {
	apps := []entity.Appointment{entity.AppointmentMock}
	return apps, nil
}

// Appointment retrieves a Appointment from the database by its id
func (maRepo *MockManageAppointmetRepository) Appointment(id uint) (*entity.Appointment, []error) {
	apps := entity.AppointmentMock
	if id == 1 {
		return &apps, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateAppointment updats a given Appointment in the database
func (maRepo *MockManageAppointmetRepository) UpdateAppointment(user *entity.Appointment) (*entity.Appointment, []error) {
	prec := entity.AppointmentMock
	return &prec, nil
}

// DeleteAppointment deletes a given Appointment from the database
func (maRepo *MockManageAppointmetRepository) DeleteAppointment(id uint) (*entity.Appointment, []error) {
	prec := entity.AppointmentMock
	return &prec, nil
}

// StoreAppointment stores a given Appointment in the database
func (maRepo *MockManageAppointmetRepository) StoreAppointment(user *entity.Appointment) (*entity.Appointment, []error) {
	prec := user
	return prec, nil
}
