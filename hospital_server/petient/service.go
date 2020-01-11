package petient

import "github.com/getach1/web1/web1_group_project-master/hospital_server/entity"

type PetientService interface {
	Petients() ([]entity.Petient, []error)
	Petient(id uint) (*entity.Petient, []error)
	UpdatePetient(petient *entity.Petient) (*entity.Petient, []error)
	DeletePetient(id uint) (*entity.Petient, []error)
	StorePetient(petient *entity.Petient) (*entity.Petient, []error)
}

type AppointmentService interface {
	Appointments() ([]entity.Appointment, []error)
	Appointment(id uint) (*entity.Appointment, []error)
}


type PrescriptionService interface {
	Prescriptions() ([]entity.Prescription, []error)
	Prescription(id uint) (*entity.Prescription, []error)
}

type RequestService interface {
	Requests() ([]entity.Request, []error)
	Request(id uint) (*entity.Request, []error)
	StoreRequest(request *entity.Request) (*entity.Request, []error)
}
type DoctorService interface {
	Doctors() ([]entity.Doctor, []error)
	Doctor(id uint) (*entity.Doctor, []error)
}
