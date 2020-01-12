package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/web1_group_project/hospital_server/Laboratorist"
	"github.com/web1_group_project/hospital_server/entity"
)

//LabratoristProfileGormRepo rep
type LabratoristProfileGormRepo struct {
	conn *gorm.DB
}

// NewPharmacistProfileGormRepo creates a new object of UserGormRepo
func NewLabratoristProfileGormRepo(db *gorm.DB) Laboratorist.LabratoristProfileRepository {
	return &LabratoristProfileGormRepo{conn: db}
}

func (p LabratoristProfileGormRepo) Profile(id uint) (*entity.Laboratorist, []error) {
	prec := entity.Laboratorist{}
	errs := p.conn.Preload("Profile").Preload("Diagnosis").First(&prec, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &prec, errs
}

func (p LabratoristProfileGormRepo) Profiles() ([]entity.Laboratorist, []error) {

	prec := []entity.Laboratorist{}
	errs := p.conn.Preload("Profile").Preload("Diagnosis").Find(&prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	return prec, errs
}

func (p LabratoristProfileGormRepo) UpdateProfile(laboratorist *entity.Laboratorist) (*entity.Laboratorist, []error) {

	prec := laboratorist
	errs := p.conn.Save(prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}

func (p LabratoristProfileGormRepo) DeleteProfile(id uint) (*entity.Laboratorist, []error) {

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

func (p LabratoristProfileGormRepo) AddProfile(labratorist *entity.Laboratorist) (*entity.Laboratorist, []error) {

	prec := labratorist
	errs := p.conn.Create(prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}
