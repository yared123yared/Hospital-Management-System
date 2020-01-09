package repository

import (
	"github.com/fasikawkn/web1_group_project_me/hospital/Pharmacist/prescription"
	"github.com/fasikawkn/web1_group_project_me/hospital/entity"
)

//PrescriptionService Implements the user.UserRepository interface
type PrescriptionService struct {
	prescSrv prescription.PrescriptionRepository
}

//NewMedicineGormRepo creates a new object of UserGormRepo
func NewPrescriptionService(prescRepository prescription.PrescriptionRepository) prescription.PrescriptionService {
	return &PrescriptionService{prescRepository}
}

func (p PrescriptionService) AddPrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {
	presc, errs := p.prescSrv.AddPrescription(prescription)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p PrescriptionService) DeletePrescription(id uint) (*entity.Prescription, []error) {
	presc, errs := p.prescSrv.DeletePrescription(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p PrescriptionService) Prescription(id uint) (*entity.Prescription, []error) {
	presc, errs := p.prescSrv.Prescription(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p PrescriptionService) Prescriptions() ([]entity.Prescription, []error) {
	presc, errs := p.prescSrv.Prescriptions()
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p PrescriptionService) UpdatePrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {

	presc, errs := p.prescSrv.UpdatePrescription(prescription)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}
