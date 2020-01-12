package PetientRepository

import (
	"github.com/jinzhu/gorm"
	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// AdminGormRepo Implements the menu.AdminRepository interface
type AdminGormRepo struct {
	conn *gorm.DB
}

// NewAdminGormRepo creates a new object of AdminGormRepo
func NewAdminGormRepo(db *gorm.DB) petient.AdminRepository {
	return &AdminGormRepo{conn: db}
}


// Admin retrieves a admin by its id from the database
func (adminRepo *AdminGormRepo) Admins() ([]entity.Admin, []error) {
	admin := []entity.Admin{}
	errs := adminRepo.conn.Debug().Find(&admin).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return admin, errs
}