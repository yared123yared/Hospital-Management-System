package Doctor

import "github.com/web1_group_project/hospital_server/entity"

// UserRepository specifies application user related database operations
type PatientRepository interface {
	Patientes() ([]entity.Petient, []error)
	Patient(id uint) (*entity.Petient, []error)
	UpdatePatient(user *entity.Petient) (*entity.Petient, []error)
	DeletePatient(id uint) (*entity.Petient, []error)
	StorePatient(user *entity.Petient) (*entity.Petient, []error)
}
type AppointmentRepository interface {
	Appointments() ([]entity.Doctor, []error)
	Appointment(id uint) (*entity.Doctor, []error)
	UpdateAppointment(user *entity.Doctor) (*entity.Doctor, []error)
	DeleteAppointment(id uint) (*entity.Doctor, []error)
}
type GeneralRepository interface {
	Pharmacists() ([]entity.Pharmacist, []error)
	Laboratorists() ([]entity.Laboratorist, []error)
}
