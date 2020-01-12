package Laboratorist



import (
	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
)

//DiagnosisService  specifies Medicine service
type DiagnosisService interface {
	Diagnosis(id uint) (*entity.Diagnosis, []error)
	Diagnosiss() ([]entity.Diagnosis, []error)
	UpdateDiagnosis(diagnosis *entity.Diagnosis) (*entity.Diagnosis, []error)
	DeleteDiagnosis(id uint) (*entity.Diagnosis, []error)
	AddDiagnosis(diagnosis *entity.Diagnosis) (*entity.Diagnosis, []error)
	GetDiagnosiss() ([]entity.Diagnosis, []error)
}

//LabratoristProfileService specifies Medicine database
type LabratoristProfileService interface {
	Profile(id uint) (*entity.Laboratorist, []error)
	Profiles() ([]entity.Laboratorist, []error)
	UpdateProfile(labratorist *entity.Laboratorist) (*entity.Laboratorist, []error)
	DeleteProfile(id uint) (*entity.Laboratorist, []error)
	AddProfile(labratorist *entity.Laboratorist) (*entity.Laboratorist, []error)
}
