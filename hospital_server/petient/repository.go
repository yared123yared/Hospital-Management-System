package petient

import "github.com/getach1/web1/hospital/entity"

type PetientRepository interface {
	Petients() ([]entity.Petient, []error)
	Petient(id uint) (*entity.Petient, []error)
	UpdatePetient(petient *entity.Petient) (*entity.Petient, []error)
	DeletePetient(id uint) (*entity.Petient, []error)
	StorePetient(petient *entity.Petient) (*entity.Petient, []error)
}
