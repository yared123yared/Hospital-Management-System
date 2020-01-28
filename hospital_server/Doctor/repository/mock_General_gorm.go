package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	//"github.com/yaredsolomon/webProgram1/hospital/request"

	"github.com/web1_group_project/hospital_server/Doctor"
	"github.com/web1_group_project/hospital_server/entity"
)

// AppointmentGormRepo Implements the request.AppointmentRepository interface
type MockGeneralGormRepo struct {
	conn *gorm.DB
}

// NewGeneralGormRepo creates a new object of GeneralGormRepo
func NewMockGeneralGormRepo(db *gorm.DB) Doctor.GeneralRepository {
	return &MockGeneralGormRepo{conn: db}
}

// Pharmacists return all Pharmacists from the database
func (generalRepo *MockGeneralGormRepo) Pharmacists() ([]entity.Pharmacist, []error) {
	pharmacists := []entity.Pharmacist{entity.PharmacistMock}
	return pharmacists, []error{errors.New("Not found")}
}
func (generalRepo *MockGeneralGormRepo) Laboratorists() ([]entity.Laboratorist, []error) {
	laboratorists := []entity.Laboratorist{entity.LaboratoristMock}
	return laboratorists, []error{errors.New("Not found")}
}
func (generalRepo *MockGeneralGormRepo) Users(id int, password string) (*entity.Profile, []error) {

	users := entity.ProfileMock
	if id == 1 {
		return &users, nil
	}
	return nil, []error{errors.New("Not found")}

}
