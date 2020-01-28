package repository

import (
	"errors"

	"github.com/web1_group_project/hospital_server/Laboratorist"
	"github.com/web1_group_project/hospital_server/entity"

	"github.com/jinzhu/gorm"
)

//DiagnosisGormRepo Implements the user.UserRepository interface
type MockDiagnosisGormRepo struct {
	conn *gorm.DB
}

// NewMedicineGormRepo creates a new object of UserGormRepo
func NewMockDiagnosisGormRepo(db *gorm.DB) Laboratorist.DiagnosisRepository {
	return &MockDiagnosisGormRepo{conn: db}
}

func (p MockDiagnosisGormRepo) AddDiagnosis(diagnosis *entity.Diagnosis) (*entity.Diagnosis, []error) {
	prec := diagnosis
	return prec, nil
}

func (p MockDiagnosisGormRepo) DeleteDiagnosis(id uint) (*entity.Diagnosis, []error) {
	prec := entity.DiagnosisMock
	return &prec, nil
}

func (p MockDiagnosisGormRepo) Diagnosis(id uint) (*entity.Diagnosis, []error) {
	prec := entity.DiagnosisMock
	if id == 1 {
		return &prec, nil
	}

	return nil, []error{errors.New("Not Found")}
}

func (p MockDiagnosisGormRepo) Diagnosiss() ([]entity.Diagnosis, []error) {
	prec := []entity.Diagnosis{entity.DiagnosisMock}

	return prec, nil
}

func (p MockDiagnosisGormRepo) GetDiagnosiss() ([]entity.Diagnosis, []error) {
	prec := []entity.Diagnosis{entity.DiagnosisMock}

	return prec, nil
}

func (p MockDiagnosisGormRepo) UpdateDiagnosis(diagnosis *entity.Diagnosis) (*entity.Diagnosis, []error) {
	prec := entity.DiagnosisMock
	return &prec, nil
}
