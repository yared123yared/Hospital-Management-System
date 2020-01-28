package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/web1_group_project/hospital_server/Laboratorist"
	"github.com/web1_group_project/hospital_server/entity"
)

//MockLabratoristProfileGormRepo rep
type MockLabratoristProfileGormRepo struct {
	conn *gorm.DB
}

// NewMockLabratoristProfileGormRepo creates a new object of UserGormRepo
func NewMockLabratoristProfileGormRepo(db *gorm.DB) Laboratorist.LabratoristProfileRepository {
	return &MockLabratoristProfileGormRepo{conn: db}
}

//Profile returns
func (p MockLabratoristProfileGormRepo) Profile(id uint) (*entity.Laboratorist, []error) {
	prec := entity.LaboratoristMock
	if id == 1 {
		return &prec, nil
	}

	return nil, []error{errors.New("Not Found")}
}

//Profiles return
func (p MockLabratoristProfileGormRepo) Profiles() ([]entity.Laboratorist, []error) {

	prec := []entity.Laboratorist{entity.LaboratoristMock}

	return prec, nil
}

//UpdateProfile updates
func (p MockLabratoristProfileGormRepo) UpdateProfile(laboratorist *entity.Laboratorist) (*entity.Laboratorist, []error) {

	prec := entity.LaboratoristMock
	return &prec, nil
}

//DeleteProfile delets
func (p MockLabratoristProfileGormRepo) DeleteProfile(id uint) (*entity.Laboratorist, []error) {

	prec := entity.LaboratoristMock
	return &prec, nil
}

//AddProfile addes
func (p MockLabratoristProfileGormRepo) AddProfile(laboratorist *entity.Laboratorist) (*entity.Laboratorist, []error) {

	prec := laboratorist
	return prec, nil
}
