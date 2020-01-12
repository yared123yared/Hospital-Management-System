package PetientService

import (
	"fmt"
	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)
type DoctorService struct {
	doctorRepo petient.DoctorRepository
}

// NewPrescriptionService  returns a new PrescriptionService object
func NewDoctorService(doctorRepository petient.DoctorRepository) petient.DoctorService {
	return &DoctorService{doctorRepo: doctorRepository}
}
// Prescriptions returns all stored application prescriptions
func (us *DoctorService) Doctors() ([]entity.Doctor, []error) {
	fmt.Println("Getinng doctotrs ......... in service")

	usrs, errs := us.doctorRepo.Doctors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// Prescription retrieves an application prescription by its id
func (us *DoctorService) Doctor(id uint) (*entity.Doctor, []error) {
	usr, errs := us.doctorRepo.Doctor(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
