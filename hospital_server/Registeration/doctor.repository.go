package Registeration



import  "github.com/yaredsolomon/webProgram1/hospital/entity"


// UserRepository specifies application user related database operations
type PatientRepository interface {
	Patientes() ([]entity.Petient, []error)
	Patient(id uint) (*entity.Petient, []error)
	UpdatePatient(user *entity.Petient) (*entity.Petient, []error)
	DeletePatient(id uint) (*entity.Petient, []error)
	StorePatient(user *entity.Petient) (*entity.Petient, []error)
}

