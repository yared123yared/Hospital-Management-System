package service



import (
	"github.com/fasikawkn/web1_group_project-1/hospital_server/Pharmacist"
	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
)

//PrescriptionService Implements the user.UserRepository interface
type PrescriptionService struct {
	prescSrv Pharmacist.PrescriptionRepository
}

//NewMedicineGormRepo creates a new object of UserGormRepo
func NewPrescriptionService(prescRepository Pharmacist.PrescriptionRepository) Pharmacist.PrescriptionService {
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
func (p PrescriptionService) GetPrescriptions() ([]entity.Prescription, []error) {
	pres, errs := p.prescSrv.GetPrescriptions()
	if len(errs) > 0 {
		return nil, errs
	}
	return pres, errs

}
func (p PrescriptionService) UpdatePrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {

	presc, errs := p.prescSrv.UpdatePrescription(prescription)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}
