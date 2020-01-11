package PetientRepository

import (
	"fmt"
	"github.com/getach1/web1/web1_group_project-master/hospital_server/entity"
	"github.com/getach1/web1/web1_group_project-master/hospital_server/petient"
	"github.com/jinzhu/gorm"
)

// DoctorGormRepo Implements the menu.DoctorRepository interface
type DoctorGormRepo struct {
	conn *gorm.DB
}

// NewDoctorGormRepo creates a new object of DoctorGormRepo
func NewDoctorGormRepo(db *gorm.DB) petient.DoctorRepository {
	return &DoctorGormRepo{conn: db}
}

// Doctors return all doctors from the database
func (doctorRepo *DoctorGormRepo) Doctors() ([]entity.Doctor, []error) {
	fmt.Println("Getinng doctotrs ......... in repo")

	doctors := []entity.Doctor{}
	errs := doctorRepo.conn.Debug().Preload("Profile").Preload("Prescription").Preload("Diagnosis").Preload("Appointment").Preload("Diagnosis").Find(&doctors).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return doctors, errs
}

// Doctor retrieves a doctor by its id from the database
func (doctorRepo *DoctorGormRepo) Doctor(id uint) (*entity.Doctor, []error) {
	doctor := entity.Doctor{}
	errs := doctorRepo.conn.Debug().Preload("Profile").Preload("Prescription").Preload("Request").Preload("Appointment").Preload("Diagnosis").Find(&doctor, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &doctor, errs
}
