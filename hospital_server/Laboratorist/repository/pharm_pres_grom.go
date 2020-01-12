package repository

import (
	"fmt"

	"github.com/fasikawkn/web1_group_project/hospital_server/Laboratorist"
	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
	"github.com/jinzhu/gorm"
)

//DiagnosisGormRepo Implements the user.UserRepository interface
type DiagnosisGormRepo struct {
	conn *gorm.DB
}

// NewMedicineGormRepo creates a new object of UserGormRepo
func NewDiagnosisGormRepo(db *gorm.DB) Laboratorist.DiagnosisRepository {
	return &DiagnosisGormRepo{conn: db}
}

func (p DiagnosisGormRepo) AddDiagnosis(diagnosis *entity.Diagnosis) (*entity.Diagnosis, []error) {
	prec := diagnosis
	errs := p.conn.Create(prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}

func (p DiagnosisGormRepo) DeleteDiagnosis(id uint) (*entity.Diagnosis, []error) {
	prec, errs := p.Diagnosis(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = p.conn.Delete(prec, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}

func (p DiagnosisGormRepo) Diagnosis(id uint) (*entity.Diagnosis, []error) {
	prec := entity.Diagnosis{}
	errs := p.conn.First(&prec, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &prec, errs
}
func (p DiagnosisGormRepo) Diagnosiss() ([]entity.Diagnosis, []error) {
	prec := []entity.Diagnosis{}
	errs := p.conn.Find(&prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}

func (p DiagnosisGormRepo) GetDiagnosiss() ([]entity.Diagnosis, []error) {
	prec := []entity.Diagnosis{}
	errs := p.conn.Find(&prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	fmt.Println(prec)
	prescriptions := []entity.Diagnosis{}
	for index := 0; index < len(prec); index++ {
		if prec[index].Reponse == "" {
			prescriptions = append(prescriptions, prec[index])
		}
	}

	return prescriptions, errs
}

func (p DiagnosisGormRepo) UpdateDiagnosis(diagnosis *entity.Diagnosis) (*entity.Diagnosis, []error) {

	prec := diagnosis
	errs := p.conn.Save(prec).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prec, errs
}
