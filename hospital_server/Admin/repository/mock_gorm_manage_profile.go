package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageProfileRepository implements Admin.ManageProfileRepository interface
type MockManageProfileRepository struct {
	conn *gorm.DB
}

// ManageProfileRepository returns new object of ManageProfileRepository
func NewMockManageProfileRepository(db *gorm.DB) Admin.ManageProfileRepository {
	return &MockManageProfileRepository{conn: db}
}

// Doctors return all doctors stored in the databasee
func (mpRepo *MockManageProfileRepository) Profiles() ([]entity.Profile, []error) {
	prfs := []entity.Profile{entity.ProfileMock}
	return prfs, nil
}

// Doctor retrieves a doctor from the database by its id
func (mpRepo *MockManageProfileRepository) Profile(id uint) (*entity.Profile, []error) {
	prfs := entity.ProfileMock
	if id == 1 {
		return &prfs, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateDoctor updats a given doctor in the database
func (mpRepo *MockManageProfileRepository) UpdateProfile(user *entity.Profile) (*entity.Profile, []error) {
	prec := entity.ProfileMock
	return &prec, nil
}

// DeleteDoctor deletes a given doctor from the database
func (mpRepo *MockManageProfileRepository) DeleteProfile(id uint) (*entity.Profile, []error) {
	prec := entity.ProfileMock
	return &prec, nil
}

// StoreDoctor stores a given doctor in the database
func (mpRepo *MockManageProfileRepository) StoreProfile(user *entity.Profile) (*entity.Profile, []error) {
	prec := user
	return prec, nil
}
