package service

import (
	"github.com/fasikawkn/web1_group_project-1/hospital_server/Admin"
	"github.com/fasikawkn/web1_group_project-1/hospital_server/entity"
)

// ManagePatientsService implements Admin.ManagePatientsService interface
type ManagePatientsService struct {
	mpRepo Admin.ManagePatientsRepository
}

//ManagePatientsService returns new ManagePatientsService object
func NewManagePatientsService(mdsLR Admin.ManagePatientsRepository) Admin.ManagePatientsService {
	return &ManagePatientsService{mpRepo: mdsLR}
}

// Doctors returns list of Doctors
func (mpSer *ManagePatientsService) Patients() ([]entity.Petient, []error) {
	prf, errs := mpSer.mpRepo.Patients()
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// Doctor returns a Doctor object with a given id
func (mpSer *ManagePatientsService) Patient(id uint) (*entity.Petient, []error) {
	prf, errs := mpSer.mpRepo.Patient(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// UpdateDoctor updates a Doctor with new data
func (mpSer *ManagePatientsService) UpdatePatient(user *entity.Petient) (*entity.Petient, []error) {
	prf, errs := mpSer.mpRepo.UpdatePatient(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// DeleteDoctor delete a Doctor by its id
func (mpSer *ManagePatientsService) DeletePatient(id uint) (*entity.Petient, []error) {
	prf, errs := mpSer.mpRepo.DeletePatient(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// StoreDoctor persists new Doctor information
func (mpSer *ManagePatientsService) StorePatient(user *entity.Petient) (*entity.Petient, []error) {
	prf, errs := mpSer.mpRepo.StorePatient(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}
