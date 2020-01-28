package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageDoctorsRepository implements Admin.ManagePharmasistsRepository interface
type MockManagePharmasistsRepository struct {
	conn *gorm.DB
}

// NewManagePharmasistsRepository returns new object of ManagePharmasistsRepository
func NewMockManagePharmasistsRepository(db *gorm.DB) Admin.ManagePharmasistsRepository {
	return &MockManagePharmasistsRepository{conn: db}
}

// Doctors return all doctors stored in the databasee
func (mpRepo *MockManagePharmasistsRepository) Pharmasists() ([]entity.Pharmacist, []error) {
	phas := []entity.Pharmacist{entity.PharmacistMock}
	return phas, nil
}

// Doctor retrieves a doctor from the database by its id
func (mpRepo *MockManagePharmasistsRepository) Pharmasist(id uint) (*entity.Pharmacist, []error) {
	phas := entity.PharmacistMock
	if id == 1 {
		return &phas, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateDoctor updats a given doctor in the database
func (mpRepo *MockManagePharmasistsRepository) UpdatePharmasist(user *entity.Pharmacist) (*entity.Pharmacist, []error) {
	prec := entity.PharmacistMock
	return &prec, nil
}

// DeleteDoctor deletes a given doctor from the database
func (mpRepo *MockManagePharmasistsRepository) DeletePharmasist(id uint) (*entity.Pharmacist, []error) {
	prec := entity.PharmacistMock
	return &prec, nil
}

// StoreDoctor stores a given doctor in the database
func (mpRepo *MockManagePharmasistsRepository) StorePharmasist(user *entity.Pharmacist) (*entity.Pharmacist, []error) {
	prec := user
	return prec, nil
}
