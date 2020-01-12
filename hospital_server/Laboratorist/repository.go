package Laboratorist

import "github.com/fasikawkn/web1_group_project/hospital_server/entity"

//DiagnosisRepository  specifies Medicine database
type DiagnosisRepository interface {
	Diagnosis(id uint) (*entity.Diagnosis, []error)
	Diagnosiss() ([]entity.Diagnosis, []error)
	UpdateDiagnosis(diagnosis *entity.Diagnosis) (*entity.Diagnosis, []error)
	DeleteDiagnosis(id uint) (*entity.Diagnosis, []error)
	AddDiagnosis(diagnosis *entity.Diagnosis) (*entity.Diagnosis, []error)
	GetDiagnosiss() ([]entity.Diagnosis, []error)
}

//LabratoristProfileRepository specifies Medicine database
type LabratoristProfileRepository interface {
	Profile(id uint) (*entity.Laboratorist, []error)
	Profiles() ([]entity.Laboratorist, []error)
	UpdateProfile(labratorist *entity.Laboratorist) (*entity.Laboratorist, []error)
	DeleteProfile(id uint) (*entity.Laboratorist, []error)
	AddProfile(labratorist *entity.Laboratorist) (*entity.Laboratorist, []error)
}
