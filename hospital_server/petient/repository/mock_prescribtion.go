package PetientRepository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// PrescriptionGormRepo Implements the menu.PrescriptionRepository interface
type MockPrescriptionGormRepo struct {
	conn *gorm.DB
}

// NewPrescriptionGormRepo creates a new object of PrescriptionGormRepo
func NewMockPrescriptionGormRepo(db *gorm.DB) petient.PrescriptionRepository {
	return &MockPrescriptionGormRepo{conn: db}
}

// Prescriptions return all prescriptions from the database
func (prescriptionRepo *MockPrescriptionGormRepo) Prescriptions() ([]entity.Prescription, []error) {
	prescriptions := []entity.Prescription{entity.PrescriptionMock}
	return prescriptions, nil
}

// Prescription retrieves a prescription by its id from the database
func (prescriptionRepo *MockPrescriptionGormRepo) Prescription(id uint) (*entity.Prescription, []error) {
	prescription := entity.PrescriptionMock
	if id == 1 {
		return &prescription, nil
	}
	return nil, []error{errors.New("Not found")}
}
