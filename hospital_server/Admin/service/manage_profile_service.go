package service

import (
	"github.com/fasikawkn/web1_group_project-1/hospital_server/Admin"
	"github.com/fasikawkn/web1_group_project-1/hospital_server/entity"
)

// ManageDoctorsService implements Admin.ManageDoctorsService interface
type ManageProfileService struct {
	mpRepo Admin.ManageProfileRepository
}

//NewManageDoctorsService returns new ManageDoctorsService object
func NewManageProfileService(mdsLR Admin.ManageProfileRepository) Admin.ManageProfileService {
	return &ManageProfileService{mpRepo: mdsLR}
}

// Doctors returns list of Doctors
func (mpSer *ManageProfileService) Profiles() ([]entity.Profile, []error) {
	prf, errs := mpSer.mpRepo.Profiles()
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// Doctor returns a Doctor object with a given id
func (mpSer *ManageProfileService) Profile(id uint) (*entity.Profile, []error) {
	prf, errs := mpSer.mpRepo.Profile(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// UpdateDoctor updates a Doctor with new data
func (mpSer *ManageProfileService) UpdateProfile(user *entity.Profile) (*entity.Profile, []error) {
	prf, errs := mpSer.mpRepo.UpdateProfile(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// DeleteDoctor delete a Doctor by its id
func (mpSer *ManageProfileService) DeleteProfile(id uint) (*entity.Profile, []error) {
	prf, errs := mpSer.mpRepo.DeleteProfile(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// StoreDoctor persists new Doctor information
func (mpSer *ManageProfileService) StoreProfile(user *entity.Profile) (*entity.Profile, []error) {

	prf, errs := mpSer.mpRepo.StoreProfile(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}
