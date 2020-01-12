package Pharmacist

import "github.com/web1_group_project/hospital_server/entity"

//MedicineRepository  specifies Medicine database
type MedicineRepository interface {
	Medicine(id uint) (*entity.Medicine, []error)
	Medicines() ([]entity.Medicine, []error)
	UpdateMedicine(medicine *entity.Medicine) (*entity.Medicine, []error)
	DeleteMedicine(id uint) (*entity.Medicine, []error)
	AddMedicine(medicine *entity.Medicine) (*entity.Medicine, []error)
	GetMedicines(addedby uint) ([]entity.Medicine, []error)
}

//PharmacistProfileRepository specifies Medicine database
type PharmacistProfileRepository interface {
	Profile(id uint) (*entity.Pharmacist, []error)
	Profiles() ([]entity.Pharmacist, []error)
	UpdateProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error)
	DeleteProfile(id uint) (*entity.Pharmacist, []error)
	AddProfile(pharmacist *entity.Pharmacist) (*entity.Pharmacist, []error)
}

//PrescriptionRepository  specifies Medicine database
type PrescriptionRepository interface {
	Prescription(id uint) (*entity.Prescription, []error)
	Prescriptions() ([]entity.Prescription, []error)
	UpdatePrescription(medicine *entity.Prescription) (*entity.Prescription, []error)
	DeletePrescription(id uint) (*entity.Prescription, []error)
	AddPrescription(medicine *entity.Prescription) (*entity.Prescription, []error)
	GetPrescriptions() ([]entity.Prescription, []error)
}
