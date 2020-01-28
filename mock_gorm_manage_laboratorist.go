package repository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageLaboratoristsRepository implements Admin.ManageLaboratoristsRepository interface
type MockManageLaboratoristsRepository struct {
	conn *gorm.DB
}

// NewManageLaboratoristsRepository returns new object of ManageLaboratoristsRepository
func NewMockManageLaboratoristsRepository(db *gorm.DB) Admin.ManageLaboratoristsRepository {
	return &MockManageLaboratoristsRepository{conn: db}
}

// Laboratorsts return all Laboratorst stored in the databasee
func (mlRepo *MockManageLaboratoristsRepository) Laboratorsts() ([]entity.Laboratorist, []error) {
	labs := []entity.Laboratorist{entity.LaboratoristMock}
	return labs, nil
}

// Laboratorst retrieves a Laboratorst from the database by its id
func (mlRepo *MockManageLaboratoristsRepository) Laboratorst(id uint) (*entity.Laboratorist, []error) {
	labs := entity.LaboratoristMock
	if id == 1 {
		return &labs, nil
	}
	return &labs, []error{errors.New("Not found")}
}

// UpdateLaboratorst updats a given Laboratorst in the database
func (mlRepo *MockManageLaboratoristsRepository) UpdateLaboratorst(user *entity.Laboratorist) (*entity.Laboratorist, []error) {
	prec := entity.LaboratoristMock
	return &prec, nil
}

// DeleteLaboratorst deletes a given Laboratorst from the database
func (mlRepo *MockManageLaboratoristsRepository) DeleteLaboratorst(id uint) (*entity.Laboratorist, []error) {
	prec := entity.LaboratoristMock
	return &prec, nil
}

// StoreLaboratorst stores a given Laboratorst in the database
func (mlRepo *MockManageLaboratoristsRepository) StoreLaboratorst(user *entity.Laboratorist) (*entity.Laboratorist, []error) {
	prec := user
	return prec, nil
}
