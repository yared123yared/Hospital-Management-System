package service

import (
	"github.com/getach1/web1/hospital/entity"
	"github.com/getach1/web1/hospital/petient"
)

// PetientService implements menu.PetientService interface
type PetientService struct {
	petientRepo petient.PetientRepository
}

// NewPetientService  returns a new PetientService object
func NewPetientService(petientRepository petient.PetientRepository) petient.PetientService {
	return &PetientService{petientRepo: petientRepository}
}

// Petients returns all stored application petients
func (us *PetientService) Petients() ([]entity.Petient, []error) {
	usrs, errs := us.petientRepo.Petients()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// Petient retrieves an application petient by its id
func (us *PetientService) Petient(id uint) (*entity.Petient, []error) {
	usr, errs := us.petientRepo.Petient(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UpdatePetient updates  a given application petient
func (us *PetientService) UpdatePetient(petient *entity.Petient) (*entity.Petient, []error) {
	usr, errs := us.petientRepo.UpdatePetient(petient)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeletePetient deletes a given application petient
func (us *PetientService) DeletePetient(id uint) (*entity.Petient, []error) {
	usr, errs := us.petientRepo.DeletePetient(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StorePetient stores a given application petient
func (us *PetientService) StorePetient(petient *entity.Petient) (*entity.Petient, []error) {
	usr, errs := us.petientRepo.StorePetient(petient)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}