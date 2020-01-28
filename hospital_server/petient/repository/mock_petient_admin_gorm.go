package PetientRepository

import (
	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// AdminGormRepo Implements the menu.AdminRepository interface
type MockAdminGormRepo struct {
	conn *gorm.DB
}

// NewAdminGormRepo creates a new object of AdminGormRepo
func NewMockAdminGormRepo(db *gorm.DB) petient.AdminRepository {
	return &MockAdminGormRepo{conn: db}
}

// Admin retrieves a admin by its id from the database
func (adminRepo *MockAdminGormRepo) Admins() ([]entity.Admin, []error) {
	admin := []entity.Admin{entity.AdminMock}
	return admin, nil
}
