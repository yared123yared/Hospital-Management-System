package medicine

import (
	"github.com/fasikawkn/web1_group_project_me/hospital/entity"
)

//MedicineService medicine related service
type MedicineService interface {
	Medicine(id uint) (*entity.Medicine, []error)
	Medicines() ([]entity.Medicine, []error)
	UpdateMedicine(medicine *entity.Medicine) (*entity.Medicine, []error)
	DeleteMedicine(id uint) (*entity.Medicine, []error)
	AddMedicine(medicine *entity.Medicine) (*entity.Medicine, []error)
}
