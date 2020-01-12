package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageAppointmetRepository implements Admin.ManageAppointmetRepository interface
type ManageAppointmetRepository struct {
	conn *gorm.DB
}

// NewManageAppointmetRepository returns new object of ManageAppointmetRepository
func NewManageAppointmetRepository(db *gorm.DB) Admin.ManageAppointmetRepository {
	return &ManageAppointmetRepository{conn: db}
}

// Appointments return all Appointments stored in the databasee
func (maRepo *ManageAppointmetRepository) Appointments() ([]entity.Appointment, []error) {
	apps := []entity.Appointment{}
	errs := maRepo.conn.Find(&apps).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return apps, errs
}

// Appointment retrieves a Appointment from the database by its id
func (maRepo *ManageAppointmetRepository) Appointment(id uint) (*entity.Appointment, []error) {
	apps := entity.Appointment{}
	errs := maRepo.conn.First(&apps, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &apps, errs
}

// UpdateAppointment updats a given Appointment in the database
func (maRepo *ManageAppointmetRepository) UpdateAppointment(user *entity.Appointment) (*entity.Appointment, []error) {
	app := user
	errs := maRepo.conn.Save(app).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return app, errs
}

// DeleteAppointment deletes a given Appointment from the database
func (maRepo *ManageAppointmetRepository) DeleteAppointment(id uint) (*entity.Appointment, []error) {
	apps, errs := maRepo.Appointment(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = maRepo.conn.Delete(apps, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return apps, errs
}

// StoreAppointment stores a given Appointment in the database
func (maRepo *ManageAppointmetRepository) StoreAppointment(user *entity.Appointment) (*entity.Appointment, []error) {
	app := user
	errs := maRepo.conn.Create(app).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return app, errs
}
