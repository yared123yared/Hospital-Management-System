package service

import (
	"github.com/getach1/web1/hospital/entity"
	prescription "github.com/getach1/web1/hospital/prescribtion"
)

// PrescriptionService implements menu.PrescriptionService interface
type PrescriptionService struct {
	prescriptionRepo prescription.PrescriptionRepository
}

// NewPrescriptionService  returns a new PrescriptionService object
func NewPrescriptionService(prescriptionRepository prescription.PrescriptionRepository) prescription.PrescriptionService {
	return &PrescriptionService{prescriptionRepo: prescriptionRepository}
}

// Prescriptions returns all stored application prescriptions
func (us *PrescriptionService) Prescriptions() ([]entity.Prescription, []error) {
	usrs, errs := us.prescriptionRepo.Prescriptions()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// Prescription retrieves an application prescription by its id
func (us *PrescriptionService) Prescription(id uint) (*entity.Prescription, []error) {
	usr, errs := us.prescriptionRepo.Prescription(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UpdatePrescription updates  a given application prescription
func (us *PrescriptionService) UpdatePrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {
	usr, errs := us.prescriptionRepo.UpdatePrescription(prescription)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeletePrescription deletes a given application prescription
func (us *PrescriptionService) DeletePrescription(id uint) (*entity.Prescription, []error) {
	usr, errs := us.prescriptionRepo.DeletePrescription(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StorePrescription stores a given application prescription
func (us *PrescriptionService) StorePrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {
	usr, errs := us.prescriptionRepo.StorePrescription(prescription)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
