package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageDoctorsRepository implements Admin.ManageDoctorsRepository interface
type MockManageDoctorsRepository struct {
	conn *gorm.DB
}

// NewManageDoctorsRepository returns new object of ManageDoctorsRepository
func NewMockManageDoctorsRepository(db *gorm.DB) Admin.ManageDoctorsRepository {
	return &MockManageDoctorsRepository{conn: db}
}

// Doctors return all doctors stored in the databasee
func (mdRepo *MockManageDoctorsRepository) Doctors() ([]entity.Doctor, []error) {
	docs := []entity.Doctor{entity.DoctorMock}
	return docs, nil
}

// Doctor retrieves a doctor from the database by its id
func (mdRepo *MockManageDoctorsRepository) Doctor(id uint) (*entity.Doctor, []error) {
	docs := entity.DoctorMock
	if id == 1 {
		return &docs, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateDoctor updats a given doctor in the database
func (mdRepo *MockManageDoctorsRepository) UpdateDoctor(user *entity.Doctor) (*entity.Doctor, []error) {
	prec := entity.DoctorMock
	return &prec, nil
}

// DeleteDoctor deletes a given doctor from the database
func (mdRepo *MockManageDoctorsRepository) DeleteDoctor(id uint) (*entity.Doctor, []error) {
	prec := entity.DoctorMock
	return &prec, nil
}

// StoreDoctor stores a given doctor in the database
func (mdRepo *MockManageDoctorsRepository) StoreDoctor(user *entity.Doctor) (*entity.Doctor, []error) {
	prec := user
	return prec, nil
}
