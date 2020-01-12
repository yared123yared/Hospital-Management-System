package service

import (
	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageDManageAppointmetServiceoctorsService implements Admin.ManageAppointmetService interface
type ManageAppointmetService struct {
	maRepo Admin.ManageAppointmetRepository
}

//NewManageAppointmetService returns new ManageAppointmetService object
func NewManageAppointmetService(mdsLR Admin.ManageAppointmetRepository) Admin.ManageAppointmetService {
	return &ManageAppointmetService{maRepo: mdsLR}
}

// Appointments returns list of Appointments
func (maSrv *ManageAppointmetService) Appointments() ([]entity.Appointment, []error) {
	apps, errs := maSrv.maRepo.Appointments()
	if len(errs) > 1 {
		return nil, errs
	}
	return apps, errs
}

// Appointments returns an Appointment object with a given id
func (maSrv *ManageAppointmetService) Appointment(id uint) (*entity.Appointment, []error) {
	app, errs := maSrv.maRepo.Appointment(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return app, errs
}

// UpdateAppointment updates an Appointment with new data
func (maSrv *ManageAppointmetService) UpdateAppointment(user *entity.Appointment) (*entity.Appointment, []error) {
	app, errs := maSrv.maRepo.UpdateAppointment(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return app, errs
}

// DeleteAppointment delete an Appointment by its id
func (maSrv *ManageAppointmetService) DeleteAppointment(id uint) (*entity.Appointment, []error) {
	app, errs := maSrv.maRepo.DeleteAppointment(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return app, errs
}

// StoreAppointment persists new Appointment information
func (maSrv *ManageAppointmetService) StoreAppointment(user *entity.Appointment) (*entity.Appointment, []error) {
	app, errs := maSrv.maRepo.StoreAppointment(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return app, errs
}
