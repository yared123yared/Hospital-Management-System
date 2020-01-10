package service

import (
	"github.com/getach1/web1/web1_group_project/hospital_server/entity"
)

// PatientService implements Registeration.PatientService interface
type PatientService struct {
	patientRepo Registeration.PatientRepository
}

// NewPatientService  returns a new PatientService object
func NewPatientService(patientRepository Registeration.PatientRepository) Registeration.PatientService {
	return &PatientService{patientRepo: patientRepository}
}

// Patientes returns all stored application Patientes
func (ps *PatientService) Patientes() ([]entity.Petient, []error) {
	petientes, errs := ps.patientRepo.Patientes()
	if len(errs) > 0 {
		return nil, errs
	}
	return petientes, errs
}

// Patient retrieves an application Patient by its id
func (ps *PatientService) Patient(id uint) (*entity.Petient, []error) {
	pst, errs := ps.patientRepo.Patient(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// UpdatePatient updates  a given application Patient
func (ps *PatientService) UpdatePatient(petient *entity.Petient) (*entity.Petient, []error) {
	pst, errs := ps.patientRepo.UpdatePatient(petient)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// DeletePatient deletes a given application Patient
func (ps *PatientService) DeletePatient(id uint) (*entity.Petient, []error) {
	pst, errs := ps.patientRepo.DeletePatient(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// StoreUser stores a given application user
func (ps *PatientService) StorePatient(petient *entity.Petient) (*entity.Petient, []error) {
	pst, errs := ps.patientRepo.StorePatient(petient)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}
