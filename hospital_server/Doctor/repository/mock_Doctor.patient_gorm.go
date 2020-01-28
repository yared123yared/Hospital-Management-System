package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Doctor"
	"github.com/web1_group_project/hospital_server/entity"
)

// UserGormRepo Implements the menu.UserRepository interface
type MockPatientGormRepo struct {
	conn *gorm.DB
}

// NewPatientGormRepo creates a new object of PatientGormRepo
func NewMockPatientGormRepo(db *gorm.DB) Doctor.PatientRepository {
	return &MockPatientGormRepo{conn: db}
}

// Patientes return all users from the database
func (patientRepo *MockPatientGormRepo) Patientes() ([]entity.Petient, []error) {
	patients := []entity.Petient{entity.PetientMock}
	return patients, nil
}

// Patient retrieves a Patient by its id from the database
func (patientRepo *MockPatientGormRepo) Patient(id uint) (*entity.Petient, []error) {
	patient := entity.PetientMock
	if id == 1 {
		return &patient, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdatePatient updates a given patient in the database
func (patientRepo *MockPatientGormRepo) UpdatePatient(patient *entity.Petient) (*entity.Petient, []error) {
	prec := entity.PetientMock
	return &prec, nil
}

// DeletePatient deletes a given patient from the database
func (patientRepo *MockPatientGormRepo) DeletePatient(id uint) (*entity.Petient, []error) {
	prec := entity.PetientMock
	return &prec, nil
}

// StoreUser stores a new user into the database
func (patientRepo *MockPatientGormRepo) StorePatient(patient *entity.Petient) (*entity.Petient, []error) {
	prec := patient
	return prec, nil
}

func (patientRepo *MockPatientGormRepo) Profile(id uint) (*entity.Profile, []error) {
	profile := entity.ProfileMock
	if id == 1 {
		return &profile, nil
	}
	return nil, []error{errors.New("Not found")}
}
