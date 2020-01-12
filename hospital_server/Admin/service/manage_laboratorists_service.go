package service

import (
	"github.com/fasikawkn/web1_group_project-1/hospital_server/Admin"
	"github.com/fasikawkn/web1_group_project-1/hospital_server/entity"
)

// ManageLaboratoristsService implements Admin.ManageLaboratoristsService interface
type ManageLaboratoristsService struct {
	mlRepo Admin.ManageLaboratoristsRepository
}

//NewManageLaboratoristsService returns new ManageLaboratoristsService object
func NewManageLaboratoristsService(mdsLR Admin.ManageLaboratoristsRepository) Admin.ManageLaboratoristsService {
	return &ManageLaboratoristsService{mlRepo: mdsLR}
}

// Laboratorsts returns list of Laboratorsts
func (mlSer *ManageLaboratoristsService) Laboratorsts() ([]entity.Laboratorist, []error) {
	prf, errs := mlSer.mlRepo.Laboratorsts()
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// Laboratorst returns a Laboratorst object with a given id
func (mlSer *ManageLaboratoristsService) Laboratorst(id uint) (*entity.Laboratorist, []error) {
	prf, errs := mlSer.mlRepo.Laboratorst(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// UpdateLaboratorst updates a Laboratorst with new data
func (mlSer *ManageLaboratoristsService) UpdateLaboratorst(user *entity.Laboratorist) (*entity.Laboratorist, []error) {
	prf, errs := mlSer.mlRepo.UpdateLaboratorst(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// DeleteLaboratorst delete a Laboratorst by its id
func (mlSer *ManageLaboratoristsService) DeleteLaboratorst(id uint) (*entity.Laboratorist, []error) {
	prf, errs := mlSer.mlRepo.DeleteLaboratorst(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// StoreLaboratorst persists new Laboratorst information
func (mlSer *ManageLaboratoristsService) StoreLaboratorst(user *entity.Laboratorist) (*entity.Laboratorist, []error) {
	prf, errs := mlSer.mlRepo.StoreLaboratorst(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}
