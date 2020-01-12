package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/fasikawkn/web1_group_project-1/hospital_server/Admin"
	"github.com/fasikawkn/web1_group_project-1/hospital_server/entity"
)

// ManageDoctorsRepository implements Admin.ManagePharmasistsRepository interface
type ManagePharmasistsRepository struct {
	conn *gorm.DB
}

// NewManagePharmasistsRepository returns new object of ManagePharmasistsRepository
func NewManagePharmasistsRepository(db *gorm.DB) Admin.ManagePharmasistsRepository {
	return &ManagePharmasistsRepository{conn: db}
}

// Doctors return all doctors stored in the databasee
func (mpRepo *ManagePharmasistsRepository) Pharmasists() ([]entity.Pharmacist, []error) {
	phas := []entity.Pharmacist{}
	errs := mpRepo.conn.Preload("Profile").Preload("Prescription").Preload("Medicine").Find(&phas).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return phas, errs
}

// Doctor retrieves a doctor from the database by its id
func (mpRepo *ManagePharmasistsRepository) Pharmasist(id uint) (*entity.Pharmacist, []error) {
	phas := entity.Pharmacist{}
	errs := mpRepo.conn.Preload("Profile").Preload("Prescription").Preload("Medicine").First(&phas, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &phas, errs
}

// UpdateDoctor updats a given doctor in the database
func (mpRepo *ManagePharmasistsRepository) UpdatePharmasist(user *entity.Pharmacist) (*entity.Pharmacist, []error) {
	pha := user
	errs := mpRepo.conn.Save(pha).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pha, errs
}

// DeleteDoctor deletes a given doctor from the database
func (mpRepo *ManagePharmasistsRepository) DeletePharmasist(id uint) (*entity.Pharmacist, []error) {
	pha, errs := mpRepo.Pharmasist(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = mpRepo.conn.Where("uuid=?", id).Delete(&pha).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	fmt.Println("Deleting pahrmacist")
	return pha, errs
}

// StoreDoctor stores a given doctor in the database
func (mpRepo *ManagePharmasistsRepository) StorePharmasist(user *entity.Pharmacist) (*entity.Pharmacist, []error) {
	pha := user
	errs := mpRepo.conn.Create(pha).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pha, errs
}
