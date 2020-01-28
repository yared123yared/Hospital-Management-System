package PetientRepository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// PetientGormRepo Implements the menu.PetientRepository interface
type MockPetientGormRepo struct {
	conn *gorm.DB
}

// NewPetientGormRepo creates a new object of PetientGormRepo
func NewMockPetientGormRepo(db *gorm.DB) petient.PetientRepository {
	return &MockPetientGormRepo{conn: db}
}

// Petients return all petients from the database
func (petientRepo *MockPetientGormRepo) Petients() ([]entity.Petient, []error) {
	petients := []entity.Petient{entity.PetientMock}

	return petients, nil
}

// Petient retrieves a petient by its id from the database
func (petientRepo *MockPetientGormRepo) Petient(id uint) (*entity.Petient, []error) {
	petient := entity.PetientMock
	if id == 1 {
		return &petient, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdatePetient updates a given petient in the database
func (petientRepo *MockPetientGormRepo) UpdatePetient(petient *entity.Petient) (*entity.Petient, []error) {
	prec := entity.PetientMock
	return &prec, nil
}

// DeletePetient deletes a given petient from the database
func (petientRepo *MockPetientGormRepo) DeletePetient(id uint) (*entity.Petient, []error) {
	prec := entity.PetientMock
	return &prec, nil
}

// StorePetient stores a new petient into the database
func (petientRepo *MockPetientGormRepo) StorePetient(petient *entity.Petient) (*entity.Petient, []error) {
	pet := petient
	return pet, nil
}
