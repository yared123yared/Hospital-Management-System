package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManagePatientsRepository implements Admin.ManagePatientsRepository interface
type MockManagePatientsRepository struct {
	conn *gorm.DB
}

// NewManagePatientsRepository returns new object of ManagePatientsRepository
func NewMockManagePatientsRepository(db *gorm.DB) Admin.ManagePatientsRepository {
	return &MockManagePatientsRepository{conn: db}
}

// Laboratorsts return all Laboratorst stored in the databasee
func (mpRepo *MockManagePatientsRepository) Patients() ([]entity.Petient, []error) {
	pats := []entity.Petient{entity.PetientMock}
	return pats, nil
}

// Laboratorst retrieves a Laboratorst from the database by its id
func (mpRepo *MockManagePatientsRepository) Patient(id uint) (*entity.Petient, []error) {
	pats := entity.PetientMock
	if id == 1 {
		return &pats, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateLaboratorst updats a given Laboratorst in the database
func (mpRepo *MockManagePatientsRepository) UpdatePatient(user *entity.Petient) (*entity.Petient, []error) {
	prec := entity.PetientMock
	return &prec, nil
}

// DeleteLaboratorst deletes a given Laboratorst from the database
func (mpRepo *MockManagePatientsRepository) DeletePatient(id uint) (*entity.Petient, []error) {
	prec := entity.PetientMock
	return &prec, nil
}

// StoreLaboratorst stores a given Laboratorst in the database
func (mpRepo *MockManagePatientsRepository) StorePatient(user *entity.Petient) (*entity.Petient, []error) {
	prec := user
	return prec, nil
}
