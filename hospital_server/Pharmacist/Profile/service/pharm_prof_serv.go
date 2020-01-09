package service

import (
	"github.com/fasikawkn/web1_group_project_me/hospital/Pharmacist/Profile"
	"github.com/fasikawkn/web1_group_project_me/hospital/entity"
)

//PharmacistProfilesService Implements the user.UserRepository interface
type PharmacistProfileService struct {
	pharmProfSrv Profile.PharmacistProfileRepository
}

//NewPharmacistProfileService creates a new object of UserGormRepo
func NewPharmacistProfileService(pharmProfRepository Profile.PharmacistProfileRepository) Profile.PharmacistProfileService {
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
