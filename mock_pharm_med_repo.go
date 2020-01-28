package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/web1_group_project/hospital_server/Pharmacist"
	"github.com/web1_group_project/hospital_server/entity"
)

//MockPharmMedRepo struct
type MockPharmMedRepo struct {
	conn *gorm.DB
}

// NewMockPharmMedRepo creates a new object of UserGormRepo
func NewMockPharmMedRepo(db *gorm.DB) Pharmacist.MedicineRepository {
	return &MockPharmMedRepo{conn: db}
}

//GetMedicines returns
func (p MockPharmMedRepo) GetMedicines(addedby uint) ([]entity.Medicine, []error) {
	prec := []entity.Medicine{entity.MedicineMock}

	return prec, nil
}

//Medicine returns
func (p MockPharmMedRepo) Medicine(id uint) (*entity.Medicine, []error) {
	prec := entity.MedicineMock
	if id == 1 {
		return &prec, nil
	}

	return nil, []error{errors.New("Not Found")}
}

//Medicines return
func (p MockPharmMedRepo) Medicines() ([]entity.Medicine, []error) {

	prec := []entity.Medicine{entity.MedicineMock}

	return prec, nil
}

//UpdateMedicine updates
func (p MockPharmMedRepo) UpdateMedicine(medicine *entity.Medicine) (*entity.Medicine, []error) {

	prec := entity.MedicineMock
	return &prec, nil
}

//DeleteMedicine delets
func (p MockPharmMedRepo) DeleteMedicine(id uint) (*entity.Medicine, []error) {

	prec := entity.MedicineMock
	return &prec, nil
}

//AddMedicine addes
func (p MockPharmMedRepo) AddMedicine(medicine *entity.Medicine) (*entity.Medicine, []error) {

	prec := medicine
	return prec, nil
}
