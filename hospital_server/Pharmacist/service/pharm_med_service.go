package service

import (
	"github.com/web1_group_project/hospital_server/Pharmacist copy"
	"github.com/web1_group_project/hospital_server/entity"
)

// MedicineService implements menu.UserService interface
type MedicineService struct {
	medicineRepo Pharmacist.MedicineRepository
}

func NewMedicineService(userRepository Pharmacist.MedicineRepository) Pharmacist.MedicineService {
	return &MedicineService{medicineRepo: userRepository}
}

func (m MedicineService) Medicine(id uint) (*entity.Medicine, []error) {
	usr, errs := m.medicineRepo.Medicine(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (m MedicineService) Medicines() ([]entity.Medicine, []error) {
	usrs, errs := m.medicineRepo.Medicines()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}
func (m MedicineService) GetMedicines(addedby uint) ([]entity.Medicine, []error) {
	usrs, errs := m.medicineRepo.GetMedicines(addedby)
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

func (m MedicineService) UpdateMedicine(medicine *entity.Medicine) (*entity.Medicine, []error) {
	usr, errs := m.medicineRepo.UpdateMedicine(medicine)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (m MedicineService) DeleteMedicine(id uint) (*entity.Medicine, []error) {
	usr, errs := m.medicineRepo.DeleteMedicine(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (m MedicineService) AddMedicine(medicine *entity.Medicine) (*entity.Medicine, []error) {
	usr, errs := m.medicineRepo.AddMedicine(medicine)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
