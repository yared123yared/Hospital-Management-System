package Admin

import (
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageDoctorsRepository specifies application Doctor related database operations
type ManageDoctorsRepository interface {
	Doctors() ([]entity.Doctor, []error)
	Doctor(id uint) (*entity.Doctor, []error)
	UpdateDoctor(user *entity.Doctor) (*entity.Doctor, []error)
	DeleteDoctor(id uint) (*entity.Doctor, []error)
	StoreDoctor(user *entity.Doctor) (*entity.Doctor, []error)
}

// ManagManagePatientsRepositoryeDoctorsRepository specifies application Patient related database operations
type ManagePatientsRepository interface {
	Patients() ([]entity.Petient, []error)
	Patient(id uint) (*entity.Petient, []error)
	UpdatePatient(user *entity.Petient) (*entity.Petient, []error)
	DeletePatient(id uint) (*entity.Petient, []error)
	StorePatient(user *entity.Petient) (*entity.Petient, []error)
}

// ManageLaboratoristsRepository specifies application Laboratorist related database operations
type ManageLaboratoristsRepository interface {
	Laboratorsts() ([]entity.Laboratorist, []error)
	Laboratorst(id uint) (*entity.Laboratorist, []error)
	UpdateLaboratorst(user *entity.Laboratorist) (*entity.Laboratorist, []error)
	DeleteLaboratorst(id uint) (*entity.Laboratorist, []error)
	StoreLaboratorst(user *entity.Laboratorist) (*entity.Laboratorist, []error)
}

// ManagePharmasistsRepository specifies application Pharmasist related database operations
type ManagePharmasistsRepository interface {
	Pharmasists() ([]entity.Pharmacist, []error)
	Pharmasist(id uint) (*entity.Pharmacist, []error)
	UpdatePharmasist(user *entity.Pharmacist) (*entity.Pharmacist, []error)
	DeletePharmasist(id uint) (*entity.Pharmacist, []error)
	StorePharmasist(user *entity.Pharmacist) (*entity.Pharmacist, []error)
}

// ManageAppointmetRepository specifies application Appointment related database operations
type ManageAppointmetRepository interface {
	Appointments() ([]entity.Appointment, []error)
	Appointment(id uint) (*entity.Appointment, []error)
	UpdateAppointment(user *entity.Appointment) (*entity.Appointment, []error)
	DeleteAppointment(id uint) (*entity.Appointment, []error)
	StoreAppointment(user *entity.Appointment) (*entity.Appointment, []error)
}

// ManageProfileRepository specifies application Profile related database operations
// type ManageProfileRepository interface {
// 	Profiles() ([]entity.Profile, []error)
// 	Profile(id uint) (*entity.Profile, []error)
// 	UpdateProfile(user *entity.Profile) (*entity.Profile, []error)
// 	DeleteProfile(id uint) (*entity.Profile, []error)
// 	StoreProfile(user *entity.Profile) (*entity.Profile, []error)
// }
//ManageProfileRepository
type UserRepository interface {
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	UserByEmail(email string) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	DeleteUser(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
	PhoneExists(phone string) bool
	EmailExists(email string) bool
	UserRoles(*entity.User) ([]entity.Role, []error)
}
