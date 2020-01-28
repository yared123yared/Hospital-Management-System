package service

import (
	"github.com/web1_group_project/hospital_server/Pharmacist"
	"github.com/web1_group_project/hospital_server/entity"
)

//PharmacistProfilesService Implements the user.UserRepository interface
type PharmacistProfileService struct {
	pharmProfSrv Pharmacist.PharmacistProfileRepository
}

//NewPharmacistProfileService creates a new object of UserGormRepo
func NewPharmacistProfileService(pharmProfRepository Pharmacist.PharmacistProfileRepository) Pharmacist.PharmacistProfileService {
	return &PharmacistProfileService{pharmProfSrv: pharmProfRepository}
}

func (p PharmacistProfileService) Profile(id uint) (*entity.Pharmacist, []error) {
	presc, errs := p.pharmProfSrv.Profile(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p PharmacistProfileService) Profiles() ([]entity.Pharmacist, []error) {
	presc, errs := p.pharmProfSrv.Profiles()
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p PharmacistProfileService) UpdateProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error) {
	presc, errs := p.pharmProfSrv.UpdateProfile(pharmacist)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p PharmacistProfileService) DeleteProfile(id uint) (*entity.Pharmacist, []error) {
	presc, errs := p.pharmProfSrv.DeleteProfile(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p PharmacistProfileService) AddProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error) {
	presc, errs := p.pharmProfSrv.AddProfile(pharmacist)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}
