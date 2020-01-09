package prescription

import (
	"github.com/fasikawkn/web1_group_project_me/hospital/entity"
)

//PrescriptionService  specifies Medicine service
type PrescriptionService interface {
	Prescription(id uint) (*entity.Prescription, []error)
	Prescriptions() ([]entity.Prescription, []error)
	UpdatePrescription(medicine *entity.Prescription) (*entity.Prescription, []error)
	DeletePrescription(id uint) (*entity.Prescription, []error)
	AddPrescription(medicine *entity.Prescription) (*entity.Prescription, []error)
}
