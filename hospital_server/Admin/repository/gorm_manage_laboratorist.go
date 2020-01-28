package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageLaboratoristsRepository implements Admin.ManageLaboratoristsRepository interface
type ManageLaboratoristsRepository struct {
	conn *gorm.DB
}

// NewManageLaboratoristsRepository returns new object of ManageLaboratoristsRepository
func NewManageLaboratoristsRepository(db *gorm.DB) Admin.ManageLaboratoristsRepository {
	return &ManageLaboratoristsRepository{conn: db}
}

// Laboratorsts return all Laboratorst stored in the databasee
func (mlRepo *ManageLaboratoristsRepository) Laboratorsts() ([]entity.Laboratorist, []error) {
	labs := []entity.Laboratorist{}
	errs := mlRepo.conn.Preload("User").Preload("Diagnosis").Find(&labs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return labs, errs
}

// Laboratorst retrieves a Laboratorst from the database by its id
func (mlRepo *ManageLaboratoristsRepository) Laboratorst(id uint) (*entity.Laboratorist, []error) {
	labs := entity.Laboratorist{}
	errs := mlRepo.conn.Preload("User").Preload("Diagnosis").First(&labs, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &labs, errs
}

// UpdateLaboratorst updats a given Laboratorst in the database
func (mlRepo *ManageLaboratoristsRepository) UpdateLaboratorst(user *entity.Laboratorist) (*entity.Laboratorist, []error) {
	lab := user
	errs := mlRepo.conn.Save(lab).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return lab, errs
}

// DeleteLaboratorst deletes a given Laboratorst from the database
func (mlRepo *ManageLaboratoristsRepository) DeleteLaboratorst(id uint) (*entity.Laboratorist, []error) {
	lab, errs := mlRepo.Laboratorst(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = mlRepo.conn.Delete(lab, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return lab, errs
}

// StoreLaboratorst stores a given Laboratorst in the database
func (mlRepo *ManageLaboratoristsRepository) StoreLaboratorst(user *entity.Laboratorist) (*entity.Laboratorist, []error) {
	lab := user
	errs := mlRepo.conn.Create(lab).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return lab, errs
}
