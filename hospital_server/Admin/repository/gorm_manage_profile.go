package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/fasikawkn/web1_group_project-1/hospital_server/Admin"
	"github.com/fasikawkn/web1_group_project-1/hospital_server/entity"
)

// ManageProfileRepository implements Admin.ManageProfileRepository interface
type ManageProfileRepository struct {
	conn *gorm.DB
}

// ManageProfileRepository returns new object of ManageProfileRepository
func NewManageProfileRepository(db *gorm.DB) Admin.ManageProfileRepository {
	return &ManageProfileRepository{conn: db}
}

// Doctors return all doctors stored in the databasee
func (mpRepo *ManageProfileRepository) Profiles() ([]entity.Profile, []error) {
	prfs := []entity.Profile{}
	errs := mpRepo.conn.Find(&prfs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prfs, errs
}

// Doctor retrieves a doctor from the database by its id
func (mpRepo *ManageProfileRepository) Profile(id uint) (*entity.Profile, []error) {
	prfs := entity.Profile{}
	errs := mpRepo.conn.First(&prfs, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &prfs, errs
}

// UpdateDoctor updats a given doctor in the database
func (mpRepo *ManageProfileRepository) UpdateProfile(user *entity.Profile) (*entity.Profile, []error) {
	prf := user
	errs := mpRepo.conn.Save(prf).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prf, errs
}

// DeleteDoctor deletes a given doctor from the database
func (mpRepo *ManageProfileRepository) DeleteProfile(id uint) (*entity.Profile, []error) {
	prf, errs := mpRepo.Profile(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = mpRepo.conn.Delete(prf, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prf, errs
}

// StoreDoctor stores a given doctor in the database
func (mpRepo *ManageProfileRepository) StoreProfile(user *entity.Profile) (*entity.Profile, []error) {
	prf := user
	errs := mpRepo.conn.Create(prf).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prf, errs
}
