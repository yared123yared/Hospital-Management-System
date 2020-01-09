package Profile

import "github.com/fasikawkn/web1_group_project_me/hospital/entity"

//PharmacistProfileService specifies Medicine database
type PharmacistProfileService interface {
	Profile(id uint) (*entity.Pharmacist, []error)
	Profiles() ([]entity.Pharmacist, []error)
	UpdateProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error)
	DeleteProfile(id uint) (*entity.Pharmacist, []error)
	AddProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error)
}
