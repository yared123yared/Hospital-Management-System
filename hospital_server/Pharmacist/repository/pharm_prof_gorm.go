package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Pharmacist"
	"github.com/web1_group_project/hospital_server/entity"
)

//PharmacistProfileGormRepo rep
type PharmacistProfileGormRepo struct {
	conn *gorm.DB
}

// NewPharmacistProfileGormRepo creates a new object of UserGormRepo
func NewPharmacistProfileGormRepo(db *gorm.DB) Pharmacist.PharmacistProfileRepository {
	return &PharmacistProfileGormRepo{conn: db}
}

func (p PharmacistProfileGormRepo) Profile(id uint) (*entity.Pharmacist, []error) {
	prec := entity.Pharmacist{}
	errs := p.conn.Preload("User").Preload("Medicine").Preload("Prescription").First(&prec, "uuid=?", id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &prec, errs
}

func (p PharmacistProfileGormRepo) Profiles() ([]entity.Pharmacist, []error) {

	prec := []entity.Pharmacist{}
	errs := p.conn.Preload("User").Preload("Medicine").Preload("Prescription").Find(&prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return prec, errs
}

func (p PharmacistProfileGormRepo) UpdateProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error) {

	prec := pharmacist
	errs := p.conn.Save(prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}

func (p PharmacistProfileGormRepo) DeleteProfile(id uint) (*entity.Pharmacist, []error) {

	prec, errs := p.Profile(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = p.conn.Delete(prec, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return prec, errs

}

func (p PharmacistProfileGormRepo) AddProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error) {

	prec := pharmacist
	errs := p.conn.Create(prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}
