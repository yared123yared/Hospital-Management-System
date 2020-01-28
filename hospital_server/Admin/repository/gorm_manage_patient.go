package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManagePatientsRepository implements Admin.ManagePatientsRepository interface
type ManagePatientsRepository struct {
	conn *gorm.DB
}

// NewManagePatientsRepository returns new object of ManagePatientsRepository
func NewManagePatientsRepository(db *gorm.DB) Admin.ManagePatientsRepository {
	return &ManagePatientsRepository{conn: db}
}

// Laboratorsts return all Laboratorst stored in the databasee
func (mpRepo *ManagePatientsRepository) Patients() ([]entity.Petient, []error) {
	pats := []entity.Petient{}
	errs := mpRepo.conn.Preload("Profile").Preload("Prescription").Preload("Diagnosis").Preload("Appointment").Find(&pats).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pats, errs
}

// Laboratorst retrieves a Laboratorst from the database by its id
func (mpRepo *ManagePatientsRepository) Patient(id uint) (*entity.Petient, []error) {
	pats := entity.Petient{}
	errs := mpRepo.conn.Preload("Profile").Preload("Prescription").Preload("Diagnosis").Preload("Appointment").First(&pats, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &pats, errs
}

// UpdateLaboratorst updats a given Laboratorst in the database
func (mpRepo *ManagePatientsRepository) UpdatePatient(user *entity.Petient) (*entity.Petient, []error) {
	pat := user
	errs := mpRepo.conn.Save(pat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pat, errs
}

// DeleteLaboratorst deletes a given Laboratorst from the database
func (mpRepo *ManagePatientsRepository) DeletePatient(id uint) (*entity.Petient, []error) {
	pat, errs := mpRepo.Patient(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = mpRepo.conn.Delete(pat, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pat, errs
}

// StoreLaboratorst stores a given Laboratorst in the database
func (mpRepo *ManagePatientsRepository) StorePatient(user *entity.Petient) (*entity.Petient, []error) {
	pat := user
	errs := mpRepo.conn.Preload("Profile").Preload("Prescription").Preload("Diagnosis").Preload("Appointment").Preload("Request").Create(pat).GetErrors()

	errs = mpRepo.conn.Create(pat.User).GetErrors()
	errs = mpRepo.conn.Create(pat.Prescription).GetErrors()
	errs = mpRepo.conn.Create(pat.Diagnosis).GetErrors()
	errs = mpRepo.conn.Create(pat.Appointment).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return pat, errs
}
