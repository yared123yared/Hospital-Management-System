package service

import (
	"github.com/web1_group_project/hospital_server/Laboratorist"
	"github.com/web1_group_project/hospital_server/entity"

)

//LabratoristProfileService Implements the user.UserRepository interface
type LabratoristProfileService struct {
	labrProfSrv Laboratorist.LabratoristProfileRepository
}

//NewPharmacistProfileService creates a new object of UserGormRepo
func NewLabratoristProfileService(laborProfRepository Laboratorist.LabratoristProfileRepository) Laboratorist.LabratoristProfileService {
	return &LabratoristProfileService{labrProfSrv: laborProfRepository}
}

func (p LabratoristProfileService) Profile(id uint) (*entity.Laboratorist, []error) {
	presc, errs := p.labrProfSrv.Profile(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p LabratoristProfileService) Profiles() ([]entity.Laboratorist, []error) {
	presc, errs := p.labrProfSrv.Profiles()
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p LabratoristProfileService) UpdateProfile(laboratorist *entity.Laboratorist) (*entity.Laboratorist, []error) {
	presc, errs := p.labrProfSrv.UpdateProfile(laboratorist)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p LabratoristProfileService) DeleteProfile(id uint) (*entity.Laboratorist, []error) {
	presc, errs := p.labrProfSrv.DeleteProfile(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p LabratoristProfileService) AddProfile(laboratorist *entity.Laboratorist) (*entity.Laboratorist, []error) {
	presc, errs := p.labrProfSrv.AddProfile(laboratorist)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}
