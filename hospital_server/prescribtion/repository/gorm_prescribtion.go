package repository

import (
	"github.com/getach1/web1/hospital/entity"
	"github.com/getach1/web1/hospital/prescribtion"
	"github.com/jinzhu/gorm"
)

// PrescriptionGormRepo Implements the menu.PrescriptionRepository interface
type PrescriptionGormRepo struct {
	conn *gorm.DB
}

// NewPrescriptionGormRepo creates a new object of PrescriptionGormRepo
func NewPrescriptionGormRepo(db *gorm.DB) Prescription.PrescriptionRepository {
	return &PrescriptionGormRepo{conn: db}
}

// Prescriptions return all prescriptions from the database
func (prescriptionRepo *PrescriptionGormRepo) Prescriptions() ([]entity.Prescription, []error) {
	prescriptions := []entity.Prescription{}
	errs := prescriptionRepo.conn.Find(&prescriptions).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prescriptions, errs
}

// Prescription retrieves a prescription by its id from the database
func (prescriptionRepo *PrescriptionGormRepo) Prescription(id uint) (*entity.Prescription, []error) {
	prescription := entity.Prescription{}
	errs := prescriptionRepo.conn.Debug().Find(&prescription, "id=?", id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &prescription, errs
}

// UpdatePrescription updates a given prescription in the database
func (prescriptionRepo *PrescriptionGormRepo) UpdatePrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {
	usr := prescription
	errs := prescriptionRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeletePrescription deletes a given prescription from the database
func (prescriptionRepo *PrescriptionGormRepo) DeletePrescription(id uint) (*entity.Prescription, []error) {
	usr, errs := prescriptionRepo.Prescription(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = prescriptionRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StorePrescription stores a new prescription into the database
func (prescriptionRepo *PrescriptionGormRepo) StorePrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {
	usr := prescription
	errs := prescriptionRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
