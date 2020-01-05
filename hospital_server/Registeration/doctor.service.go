package Registeration


import	"github.com/yaredsolomon/webProgram1/hospital/entity"



// UserService specifies application user related services
type PatientService interface {
	Patientes() ([]entity.Petient, []error)
	Patient(id uint) (*entity.Petient, []error)
	UpdatePatient(user *entity.Petient) (*entity.Petient, []error)
	DeletePatient(id uint) (*entity.Petient, []error)
	StorePatient(user *entity.Petient) (*entity.Petient, []error)
}

// RoleService speifies application user role related services

