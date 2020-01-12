package service

import (
	"github.com/fasikawkn/web1_group_project/hospital_server/Laboratorist"
	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
)

//DiagnosisService Implements the user.UserRepository interface
type DiagnosisService struct {
	diagSrv Laboratorist.DiagnosisRepository
}

//NewMedicineGormRepo creates a new object of UserGormRepo
func NewDiagnosisService(diagRepository Laboratorist.DiagnosisRepository) Laboratorist.DiagnosisService {
	return &DiagnosisService{diagRepository}
}

func (p DiagnosisService) AddDiagnosis(diagnosis *entity.Diagnosis) (*entity.Diagnosis, []error) {
	presc, errs := p.diagSrv.AddDiagnosis(diagnosis)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p DiagnosisService) DeleteDiagnosis(id uint) (*entity.Diagnosis, []error) {
	presc, errs := p.diagSrv.DeleteDiagnosis(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p DiagnosisService) Diagnosis(id uint) (*entity.Diagnosis, []error) {
	presc, errs := p.diagSrv.Diagnosis(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p DiagnosisService) Diagnosiss() ([]entity.Diagnosis, []error) {
	presc, errs := p.diagSrv.Diagnosiss()
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p DiagnosisService) UpdateDiagnosis(diagnosis *entity.Diagnosis) (*entity.Diagnosis, []error) {

	presc, errs := p.diagSrv.UpdateDiagnosis(diagnosis)
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}

func (p DiagnosisService) GetDiagnosiss() ([]entity.Diagnosis, []error) {
	presc, errs := p.diagSrv.GetDiagnosiss()
	if len(errs) > 0 {
		return nil, errs
	}
	return presc, errs
}
