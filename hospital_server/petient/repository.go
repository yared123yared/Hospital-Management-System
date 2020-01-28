package petient

import "github.com/web1_group_project/hospital_server/entity"

type PetientRepository interface {
	Petients() ([]entity.Petient, []error)
	Petient(id uint) (*entity.Petient, []error)
	Petient2(id uint) (*entity.Petient, []error)

	UpdatePetient(petient *entity.Petient) (*entity.Petient, []error)
	DeletePetient(id uint) (*entity.Petient, []error)
	StorePetient(petient *entity.Petient) (*entity.Petient, []error)
}

type AppointmentRepository interface {
	Appointments() ([]entity.Appointment, []error)
	Appointment(id uint) (*entity.Appointment, []error)
}

type PrescriptionRepository interface {
	Prescriptions() ([]entity.Prescription, []error)
	Prescription(id uint) (*entity.Prescription, []error)
}

type RequestRepository interface {
	Requests() ([]entity.Request, []error)
	Request(id uint) (*entity.Request, []error)
	StoreRequest(request *entity.Request) (*entity.Request, []error)
}
type DoctorRepository interface {
	Doctors() ([]entity.Doctor, []error)
	Doctor(id uint) (*entity.Doctor, []error)
}
type AdminRepository interface {
	Admins() ([]entity.Admin, []error)
}
