package service

import (
	"github.com/fasikawkn/web1_group_project-1/hospital_server/Admin"
	"github.com/fasikawkn/web1_group_project-1/hospital_server/entity"
)

// ManageDoctorsService implements Admin.ManageDoctorsService interface
type ManageDoctorsService struct {
	mdRepo Admin.ManageDoctorsRepository
}

//NewManageDoctorsService returns new ManageDoctorsService object
func NewManageDoctorsService(mdsLR Admin.ManageDoctorsRepository) Admin.ManageDoctorsService {
	return &ManageDoctorsService{mdRepo: mdsLR}
}

// Doctors returns list of Doctors
func (mdSer *ManageDoctorsService) Doctors() ([]entity.Doctor, []error) {
	prf, errs := mdSer.mdRepo.Doctors()
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// Doctor returns a Doctor object with a given id
func (mdSer *ManageDoctorsService) Doctor(id uint) (*entity.Doctor, []error) {
	prf, errs := mdSer.mdRepo.Doctor(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// UpdateDoctor updates a Doctor with new data
func (mdSer *ManageDoctorsService) UpdateDoctor(user *entity.Doctor) (*entity.Doctor, []error) {
	prf, errs := mdSer.mdRepo.UpdateDoctor(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// DeleteDoctor delete a Doctor by its id
func (mdSer *ManageDoctorsService) DeleteDoctor(id uint) (*entity.Doctor, []error) {
	prf, errs := mdSer.mdRepo.DeleteDoctor(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// StoreDoctor persists new Doctor information
func (mdSer *ManageDoctorsService) StoreDoctor(user *entity.Doctor) (*entity.Doctor, []error) {
	prf, errs := mdSer.mdRepo.StoreDoctor(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}
