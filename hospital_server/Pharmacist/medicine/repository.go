package medicine

import "github.com/fasikawkn/web1_group_project_me/hospital/entity"

//MedicineRepository  specifies Medicine database
type MedicineRepository interface {
	Medicine(id uint) (*entity.Medicine, []error)
	Medicines() ([]entity.Medicine, []error)
	UpdateMedicine(medicine *entity.Medicine) (*entity.Medicine, []error)
	DeleteMedicine(id uint) (*entity.Medicine, []error)
	AddMedicine(medicine *entity.Medicine) (*entity.Medicine, []error)
}

//Prescription specifies Prescription database
type PrescriptionRepository interface {
	Prescription()
	Prescriptions()
	UpdatePrescription()
	DeletePrescription()
	AddPrescription()
}
