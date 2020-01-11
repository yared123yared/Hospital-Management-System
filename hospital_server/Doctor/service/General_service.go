package service

import (
	//"github.com/yaredsolomon/webProgram1/hospital/entity"
	//"github.com/yaredsolomon/webProgram1/hospital/request"

	//"github.com/yaredsolomon/webProgram1/hospital/request"
	"github.com/web1_group_project/hospital_server/Doctor"
	"github.com/web1_group_project/hospital_server/entity"
)

// AppointmentService implements request.AppointmentService interface
type GeneralService struct {
	generalRepo Doctor.GeneralRepository
}

// NewAppointmentService  returns a new AppointmentService object
func NewGeneralService(generalRepository Doctor.GeneralRepository) Doctor.GeneralService {
	return &GeneralService{generalRepo: generalRepository}
}

// Appointments returns all stored application Appointments
func (gs *GeneralService) Pharmacists() ([]entity.Pharmacist, []error) {
	generals, errs := gs.generalRepo.Pharmacists()
	if len(errs) > 0 {
		return nil, errs
	}
	return generals, errs
}
func (gs *GeneralService) Laboratorists() ([]entity.Laboratorist, []error) {
	generals, errs := gs.generalRepo.Laboratorists()
	if len(errs) > 0 {
		return nil, errs
	}
	return generals, errs
}
func (gs *GeneralService) Users(id int, password string) (*entity.Profile, []error) {
	generals, errs := gs.generalRepo.Users(id, password)
	if len(errs) > 0 {
		return nil, errs
	}
	return generals, errs
}
