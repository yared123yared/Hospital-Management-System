package PetientRepository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// DoctorGormRepo Implements the menu.DoctorRepository interface
type MockoctorGormRepo struct {
	conn *gorm.DB
}

// NewDoctorGormRepo creates a new object of DoctorGormRepo
func NewMockDoctorGormRepo(db *gorm.DB) petient.DoctorRepository {
	return &MockoctorGormRepo{conn: db}
}

// Doctors return all doctors from the database
func (doctorRepo *MockoctorGormRepo) Doctors() ([]entity.Doctor, []error) {
	doctors := []entity.Doctor{entity.DoctorMock}
	return doctors, nil

}

// Doctor retrieves a doctor by its id from the database
func (doctorRepo *MockoctorGormRepo) Doctor(id uint) (*entity.Doctor, []error) {
	doctor := entity.DoctorMock
	if id == 1 {
		return &doctor, nil
	}
	return nil, []error{errors.New("Not found")}
}
