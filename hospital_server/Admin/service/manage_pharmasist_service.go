package service

import (
	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManagePharmasistsService implements Admin.ManagePharmasistsService interface
type ManagePharmasistsService struct {
	mpRepo Admin.ManagePharmasistsRepository
}

//NewManagePharmasistsService returns new ManagePharmasistsService object
func NewManagePharmasistsService(mdsLR Admin.ManagePharmasistsRepository) Admin.ManagePharmasistsService {
	return &ManagePharmasistsService{mpRepo: mdsLR}
}

// Doctors returns list of Doctors
func (mpSer *ManagePharmasistsService) Pharmasists() ([]entity.Pharmacist, []error) {
	prf, errs := mpSer.mpRepo.Pharmasists()
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// Doctor returns a Doctor object with a given id
func (mpSer *ManagePharmasistsService) Pharmasist(id uint) (*entity.Pharmacist, []error) {
	prf, errs := mpSer.mpRepo.Pharmasist(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// UpdateDoctor updates a Doctor with new data
func (mpSer *ManagePharmasistsService) UpdatePharmasist(user *entity.Pharmacist) (*entity.Pharmacist, []error) {
	prf, errs := mpSer.mpRepo.UpdatePharmasist(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// DeleteDoctor delete a Doctor by its id
func (mpSer *ManagePharmasistsService) DeletePharmasist(id uint) (*entity.Pharmacist, []error) {
	prf, errs := mpSer.mpRepo.DeletePharmasist(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// StoreDoctor persists new Doctor information
func (mpSer *ManagePharmasistsService) StorePharmasist(user *entity.Pharmacist) (*entity.Pharmacist, []error) {
	prf, errs := mpSer.mpRepo.StorePharmasist(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}
