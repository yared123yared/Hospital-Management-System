package PetientRepository

import (
	"fmt"
	"github.com/getach1/web1/web1_group_project-master/hospital_server/entity"
	"github.com/getach1/web1/web1_group_project-master/hospital_server/petient"
	"github.com/jinzhu/gorm"
)

// PetientGormRepo Implements the menu.PetientRepository interface
type PetientGormRepo struct {
	conn *gorm.DB
}

// NewPetientGormRepo creates a new object of PetientGormRepo
func NewPetientGormRepo(db *gorm.DB) petient.PetientRepository {
	return &PetientGormRepo{conn: db}
}

// Petients return all petients from the database
func (petientRepo *PetientGormRepo) Petients() ([]entity.Petient, []error) {
	petients := []entity.Petient{}
	errs := petientRepo.conn.Debug().Preload("Profile").Preload("Prescription").Preload("Request").Preload("Appointment").Preload("Diagnosis").Find(&petients).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return petients, errs
}

// Petient retrieves a petient by its id from the database
func (petientRepo *PetientGormRepo) Petient(id uint) (*entity.Petient, []error) {
	petient := entity.Petient{}
	errs := petientRepo.conn.Debug().Preload("Profile").Preload("Prescription").Preload("Request").Preload("Appointment").Preload("Diagnosis").Find(&petient, id).GetErrors()
	fmt.Println(petient)

	if len(errs) > 0 {
		return nil, errs
	}
	return &petient, errs
}

// UpdatePetient updates a given petient in the database
func (petientRepo *PetientGormRepo) UpdatePetient(petient *entity.Petient) (*entity.Petient, []error) {
	usr := petient
	errs := petientRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeletePetient deletes a given petient from the database
func (petientRepo *PetientGormRepo) DeletePetient(id uint) (*entity.Petient, []error) {
	usr, errs := petientRepo.Petient(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = petientRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StorePetient stores a new petient into the database
func (petientRepo *PetientGormRepo) StorePetient(petient *entity.Petient) (*entity.Petient, []error) {
	usr := petient
	errs := petientRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
