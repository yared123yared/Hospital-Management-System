package PetientRepository

import (
	"github.com/jinzhu/gorm"
	"github.com/web1_group_project/hospital_server/entity"
)

// PrescriptionGormRepo Implements the menu.PrescriptionRepository interface
type PrescriptionGormRepo struct {
	conn *gorm.DB
}

// NewPrescriptionGormRepo creates a new object of PrescriptionGormRepo
func NewPrescriptionGormRepo(db *gorm.DB) *PrescriptionGormRepo {
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
