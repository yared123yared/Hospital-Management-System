package Prescription

import "github.com/getach1/web1/hospital/entity"

type PrescriptionService interface {
	Prescriptions() ([]entity.Prescription, []error)
	Prescription(id uint) (*entity.Prescription, []error)
	UpdatePrescription(prescription *entity.Prescription) (*entity.Prescription, []error)
	DeletePrescription(id uint) (*entity.Prescription, []error)
	StorePrescription(prescription *entity.Prescription) (*entity.Prescription, []error)
}
