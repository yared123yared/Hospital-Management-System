package Pharmacist

import (
	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
)

//MedicineService medicine related service
type MedicineService interface {
	Medicine(id uint) (*entity.Medicine, []error)
	Medicines() ([]entity.Medicine, []error)
	UpdateMedicine(medicine *entity.Medicine) (*entity.Medicine, []error)
	DeleteMedicine(id uint) (*entity.Medicine, []error)
	AddMedicine(medicine *entity.Medicine) (*entity.Medicine, []error)
	GetMedicines(addedby uint) ([]entity.Medicine, []error)
}

//PrescriptionService  specifies Medicine service
type PrescriptionService interface {
	Prescription(id uint) (*entity.Prescription, []error)
	Prescriptions() ([]entity.Prescription, []error)
	UpdatePrescription(medicine *entity.Prescription) (*entity.Prescription, []error)
	DeletePrescription(id uint) (*entity.Prescription, []error)
	AddPrescription(medicine *entity.Prescription) (*entity.Prescription, []error)
	GetPrescriptions() ([]entity.Prescription, []error)
}

//PharmacistProfileService specifies Medicine database
type PharmacistProfileService interface {
	Profile(id uint) (*entity.Pharmacist, []error)
	Profiles() ([]entity.Pharmacist, []error)
	UpdateProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error)
	DeleteProfile(id uint) (*entity.Pharmacist, []error)
	AddProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error)
}
