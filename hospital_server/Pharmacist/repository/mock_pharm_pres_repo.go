package repository

import (
	"errors"

	"github.com/web1_group_project/hospital_server/Pharmacist"
	"github.com/web1_group_project/hospital_server/entity"
	"github.com/jinzhu/gorm"
	
)

//MockPrescriptionGormRepo Implements the user.UserRepository interface
type MockPrescriptionGormRepo struct {
	conn *gorm.DB
}

// NewMockPrescriptionGormRepo creates a new object of UserGormRepo
func NewMockPrescriptionGormRepo(db *gorm.DB) Pharmacist.PrescriptionRepository {
	return &PrescriptionGormRepo{conn: db}
}

//AddPrescription adds
func (p MockPrescriptionGormRepo) AddPrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {
	prec := prescription
	return prec, nil
}

//DeletePrescription deletes
func (p MockPrescriptionGormRepo) DeletePrescription(id uint) (*entity.Prescription, []error) {
	prec := entity.PrescriptionMock
	return &prec, nil
}

//Prescription retunrs
func (p MockPrescriptionGormRepo) Prescription(id uint) (*entity.Prescription, []error) {
	prec := entity.PrescriptionMock
	if id == 1 {
		return &prec, nil
	}

	return nil, []error{errors.New("Not Found")}
}

//Prescriptions returns
func (p MockPrescriptionGormRepo) Prescriptions() ([]entity.Prescription, []error) {
	prec := []entity.Prescription{entity.PrescriptionMock}

	return prec, nil
}

//GetPrescriptions gets
func (p MockPrescriptionGormRepo) GetPrescriptions() ([]entity.Prescription, []error) {
	prec := []entity.Prescription{entity.PrescriptionMock}

	return prec, nil
}

//UpdatePrescription updates
func (p MockPrescriptionGormRepo) UpdatePrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {

	prec := entity.PrescriptionMock
	return &prec, nil
}
