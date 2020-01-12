package Admin

import (
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageDotorsService specifies application Doctor related services
type ManageDoctorsService interface {
	Doctors() ([]entity.Doctor, []error)
	Doctor(id uint) (*entity.Doctor, []error)
	UpdateDoctor(user *entity.Doctor) (*entity.Doctor, []error)
	DeleteDoctor(id uint) (*entity.Doctor, []error)
	StoreDoctor(user *entity.Doctor) (*entity.Doctor, []error)
}

// ManagePatientsService specifies application Patient related services
type ManagePatientsService interface {
	Patients() ([]entity.Petient, []error)
	Patient(id uint) (*entity.Petient, []error)
	UpdatePatient(user *entity.Petient) (*entity.Petient, []error)
	DeletePatient(id uint) (*entity.Petient, []error)
	StorePatient(user *entity.Petient) (*entity.Petient, []error)
}

// ManageLaboratoristsService specifies application Laboratorist related services
type ManageLaboratoristsService interface {
	Laboratorsts() ([]entity.Laboratorist, []error)
	Laboratorst(id uint) (*entity.Laboratorist, []error)
	UpdateLaboratorst(user *entity.Laboratorist) (*entity.Laboratorist, []error)
	DeleteLaboratorst(id uint) (*entity.Laboratorist, []error)
	StoreLaboratorst(user *entity.Laboratorist) (*entity.Laboratorist, []error)
}

// ManagePharmasistsService specifies application Pharmasist related services
type ManagePharmasistsService interface {
	Pharmasists() ([]entity.Pharmacist, []error)
	Pharmasist(id uint) (*entity.Pharmacist, []error)
	UpdatePharmasist(user *entity.Pharmacist) (*entity.Pharmacist, []error)
	DeletePharmasist(id uint) (*entity.Pharmacist, []error)
	StorePharmasist(user *entity.Pharmacist) (*entity.Pharmacist, []error)
}

// ManageAppointmetService specifies application Appointment related services
type ManageAppointmetService interface {
	Appointments() ([]entity.Appointment, []error)
	Appointment(id uint) (*entity.Appointment, []error)
	UpdateAppointment(user *entity.Appointment) (*entity.Appointment, []error)
	DeleteAppointment(id uint) (*entity.Appointment, []error)
	StoreAppointment(user *entity.Appointment) (*entity.Appointment, []error)
}

// ManageProfileService specifies application Profile related services
type ManageProfileService interface {
	Profiles() ([]entity.Profile, []error)
	Profile(id uint) (*entity.Profile, []error)
	UpdateProfile(user *entity.Profile) (*entity.Profile, []error)
	DeleteProfile(id uint) (*entity.Profile, []error)
	StoreProfile(user *entity.Profile) (*entity.Profile, []error)
}
