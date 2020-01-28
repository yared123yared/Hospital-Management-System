package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/web1_group_project/hospital_server/Pharmacist"
	"github.com/web1_group_project/hospital_server/entity"
)

//MockPharmacistRepo struct
type MockPharmacistRepo struct {
	conn *gorm.DB
}

// NewMockPharmacistProfileGormRepo creates a new object of UserGormRepo
func NewMockPharmacistProfileGormRepo(db *gorm.DB) Pharmacist.PharmacistProfileRepository {
	return &MockPharmacistRepo{conn: db}
}

//Profile returns
func (p MockPharmacistRepo) Profile(id uint) (*entity.Pharmacist, []error) {
	prec := entity.PharmacistMock
	if id == 1 {
		return &prec, nil
	}

	return nil, []error{errors.New("Not Found")}
}

//Profiles return
func (p MockPharmacistRepo) Profiles() ([]entity.Pharmacist, []error) {

	prec := []entity.Pharmacist{entity.PharmacistMock}

	return prec, nil
}

//UpdateProfile updates
func (p MockPharmacistRepo) UpdateProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error) {

	prec := entity.PharmacistMock
	return &prec, nil
}

//DeleteProfile delets
func (p MockPharmacistRepo) DeleteProfile(id uint) (*entity.Pharmacist, []error) {

	prec := entity.PharmacistMock
	return &prec, nil
}

//AddProfile addes
func (p MockPharmacistRepo) AddProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error) {

	prec := pharmacist
	return prec, nil
}
