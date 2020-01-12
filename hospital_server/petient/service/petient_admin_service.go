package PetientService

import (
	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)
// AdminService implements menu.AdminService interface
type AdminService struct {
	adminRepo petient.AdminRepository
}

// NewAdminService  returns a new AdminService object
func NewAdminService(adminRepository petient.AdminRepository) petient.AdminService {
	return &AdminService{adminRepo: adminRepository}
}

// Admin retrieves an application admin by its id
func (us *AdminService) Admins() ([]entity.Admin, []error) {
	usr, errs := us.adminRepo.Admins()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
