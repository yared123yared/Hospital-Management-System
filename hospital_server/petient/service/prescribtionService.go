package PetientService

import (
	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// PrescriptionService implements menu.PrescriptionService interface
type PrescriptionService struct {
	prescriptionRepo petient.PrescriptionRepository
}

// NewPrescriptionService  returns a new PrescriptionService object
func NewPrescriptionService(prescriptionRepository petient.PrescriptionRepository) petient.PrescriptionService {
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
