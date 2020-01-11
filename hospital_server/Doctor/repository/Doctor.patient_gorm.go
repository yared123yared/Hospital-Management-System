package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/yaredsolomon/webProgram1/hospital/Registeration"
	"github.com/yaredsolomon/webProgram1/hospital/entity"
)

// UserGormRepo Implements the menu.UserRepository interface
type PatientGormRepo struct {
	conn *gorm.DB
}

// NewPatientGormRepo creates a new object of PatientGormRepo
func NewPatientGormRepo(db *gorm.DB) Registeration.PatientRepository {
	return &PatientGormRepo{conn: db}
}

// Patientes return all users from the database
func (patientRepo *PatientGormRepo) Patientes() ([]entity.Petient, []error) {
	patients := []entity.Petient{}
	errs := patientRepo.conn.Preload("Profile").Preload("Prescription").Preload("Diagnosis").Preload("Appointment").Find(&patients).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return patients, errs
}

// Patient retrieves a Patient by its id from the database
func (patientRepo *PatientGormRepo) Patient(id uint) (*entity.Petient, []error) {
	patient := entity.Petient{}
	errs := patientRepo.conn.Preload("Profile").Preload("Prescription").Preload("Diagnosis").Preload("Appointment").First(&patient, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &patient, errs
}

// UpdatePatient updates a given patient in the database
func (patientRepo *PatientGormRepo) UpdatePatient(patient *entity.Petient) (*entity.Petient, []error) {
	pst := patient
	errs := patientRepo.conn.Save(pst).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// DeletePatient deletes a given patient from the database
func (patientRepo *PatientGormRepo) DeletePatient(id uint) (*entity.Petient, []error) {
	pst, errs := patientRepo.Patient(id)
	if len(errs) > 0 {
		return nil, errs
	}
	proId := pst.Uuid
	profile, err := patientRepo.Profile(proId)
	if len(err) > 0 {
		return nil, err
	}

	fmt.Println(proId)
	errs = patientRepo.conn.Delete(pst, id).GetErrors()
	errs = patientRepo.conn.Delete(profile, proId).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// StoreUser stores a new user into the database
func (patientRepo *PatientGormRepo) StorePatient(patient *entity.Petient) (*entity.Petient, []error) {
	pst := patient
	errs := patientRepo.conn.Create(pst).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

func (patientRepo *PatientGormRepo) Profile(id uint) (*entity.Profile, []error) {
	profile := entity.Profile{}
	errs := patientRepo.conn.First(&profile, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &profile, errs
}
