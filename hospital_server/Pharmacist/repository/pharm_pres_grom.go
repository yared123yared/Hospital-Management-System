package repository

import (
	"fmt"
	"github.com/web1_group_project/hospital_server/Pharmacist copy"
	"github.com/web1_group_project/hospital_server/entity"

	"github.com/jinzhu/gorm"
)

//PrescriptionGormRepo Implements the user.UserRepository interface
type PrescriptionGormRepo struct {
	conn *gorm.DB
}

// NewMedicineGormRepo creates a new object of UserGormRepo
func NewPrescriptionGormRepo(db *gorm.DB) Pharmacist.PrescriptionRepository {
	return &PrescriptionGormRepo{conn: db}
}

func (p PrescriptionGormRepo) AddPrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {
	prec := prescription
	errs := p.conn.Create(prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}

func (p PrescriptionGormRepo) DeletePrescription(id uint) (*entity.Prescription, []error) {
	prec, errs := p.Prescription(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = p.conn.Delete(prec, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}

func (p PrescriptionGormRepo) Prescription(id uint) (*entity.Prescription, []error) {
	prec := entity.Prescription{}
	errs := p.conn.First(&prec, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &prec, errs
}

func (p PrescriptionGormRepo) Prescriptions() ([]entity.Prescription, []error) {
	prec := []entity.Prescription{}
	errs := p.conn.Find(&prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}
func (p PrescriptionGormRepo) GetPrescriptions() ([]entity.Prescription, []error) {
	prec := []entity.Prescription{}
	errs := p.conn.Find(&prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	fmt.Println(prec)
	prescriptions := []entity.Prescription{}
	for index := 0; index < len(prec); index++ {
		if prec[index].GivenStatus == "" {
			prescriptions = append(prescriptions, prec[index])
		}
	}
	fmt.Println("pres", prescriptions)
	// if prec[0].GivenStatus == "" {
	// 	fmt.Println("NULL IS REF")
	// }
	return prescriptions, errs
}

func (p PrescriptionGormRepo) UpdatePrescription(prescription *entity.Prescription) (*entity.Prescription, []error) {

	prec := prescription
	errs := p.conn.Save(prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}
