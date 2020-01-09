package repository

import (
	"github.com/fasikawkn/web1_group_project_me/hospital/Pharmacist/medicine"
	"github.com/fasikawkn/web1_group_project_me/hospital/entity"
	"github.com/jinzhu/gorm"
)

//MedicineGormRepo Implements the user.UserRepository interface
type MedicineGormRepo struct {
	conn *gorm.DB
}

// NewMedicineGormRepo creates a new object of UserGormRepo
func NewMedicineGormRepo(db *gorm.DB) medicine.MedicineRepository {
	return &MedicineGormRepo{conn: db}
}

func (m MedicineGormRepo) AddMedicine(medicine *entity.Medicine) (*entity.Medicine, []error) {
	mdn := medicine
	errs := m.conn.Create(mdn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return mdn, errs
}

func (m MedicineGormRepo) DeleteMedicine(id uint) (*entity.Medicine, []error) {
	mdn, errs := m.Medicine(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = m.conn.Delete(mdn, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return mdn, errs
}

func (m MedicineGormRepo) Medicine(id uint) (*entity.Medicine, []error) {
	medicine := entity.Medicine{}
	errs := m.conn.First(&medicine, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &medicine, errs
}

func (m MedicineGormRepo) Medicines() ([]entity.Medicine, []error) {
	medicnes := []entity.Medicine{}
	errs := m.conn.Find(&medicnes).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return medicnes, errs
}

func (m MedicineGormRepo) UpdateMedicine(medicine *entity.Medicine) (*entity.Medicine, []error) {

	mdn := medicine
	errs := m.conn.Save(mdn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return mdn, errs
}
