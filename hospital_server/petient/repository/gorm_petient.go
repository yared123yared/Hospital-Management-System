package PetientRepository

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// PetientGormRepo Implements the menu.PetientRepository interface
type PetientGormRepo struct {
	conn *gorm.DB
}

// NewPetientGormRepo creates a new object of PetientGormRepo
func NewPetientGormRepo(db *gorm.DB) petient.PetientRepository {
	return &PetientGormRepo{conn: db}
}

// Petients return all petients from the database
func (petientRepo *PetientGormRepo) Petients() ([]entity.Petient, []error) {
	petients := []entity.Petient{}
	errs := petientRepo.conn.Debug().Preload("User").Preload("Prescription").Preload("Request").Preload("Appointment").Preload("Diagnosis").Find(&petients).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return petients, errs
}

// Petient retrieves a petient by its id from the database
func (petientRepo *PetientGormRepo) Petient(id uint) (*entity.Petient, []error) {
	petient := entity.Petient{}
	errs := petientRepo.conn.Debug().Preload("User").Preload("Prescription").Preload("Request").Preload("Appointment").Preload("Diagnosis").Find(&petient, id).GetErrors()
	fmt.Println(petient)

	if len(errs) > 0 {
		return nil, errs
	}
	return &petient, errs
}
func (petientRepo *PetientGormRepo) Petient2(uuid uint) (*entity.Petient, []error) {
	petient := entity.Petient{}
	errs := petientRepo.conn.Debug().Preload("User").Preload("Prescription").Preload("Request").Preload("Appointment").Preload("Diagnosis").Find(&petient, "uuid=?", uuid).GetErrors()
	fmt.Println(petient)

	if len(errs) > 0 {
		return nil, errs
	}
	return &petient, errs
}

// UpdatePetient updates a given petient in the database
func (petientRepo *PetientGormRepo) UpdatePetient(petient *entity.Petient) (*entity.Petient, []error) {
	usr := petient
	errs := petientRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// // DeletePetient deletes a given petient from the database
// func (petientRepo *PetientGormRepo) DeletePetient(id uint) (*entity.Petient, []error) {
// 	usr, errs := petientRepo.Petient(id)
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	errs = petientRepo.conn.Delete(usr, id).GetErrors()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return usr, errs
// }

// StorePetient stores a new petient into the database
func (petientRepo *PetientGormRepo) StorePetient(petient *entity.Petient) (*entity.Petient, []error) {
	usr := petient
	errs := petientRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UpdatePetient updates a given petient in the database
func (petientRepo *PetientGormRepo) UpdatePetient(petient *entity.Petient) (*entity.Petient, []error) {
	usr := petient

	petientRepo.StoreRequest(usr.Request[len(usr.Request)-1])

	errs := petientRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeletePetient deletes a given petient from the database
func (petientRepo *PetientGormRepo) DeletePetient(id uint) (*entity.Petient, []error) {
	usr, errs := petientRepo.Petient(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = petientRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (petientRepo *PetientGormRepo) StoreRequest(request entity.Request) {
	usr := request
	_ = petientRepo.conn.Create(&usr).GetErrors()

	appointment := entity.Appointment{
		ID:          0,
		PatientId:   usr.PatientId,
		PatientName: usr.PatientName,
		DoctorId:    usr.DoctorId,
		Date:        time.Time{},
	}

	_ = petientRepo.conn.Create(&appointment).GetErrors()

	presc := entity.Prescription{
		ID:             0,
		PatientId:      usr.PatientId,
		PatientName:    usr.PatientName,
		DoctorId:       usr.DoctorId,
		PhrmacistId:    usr.ApprovedBy,
		PrescribedDate: time.Time{},
		MedicineName:   "",
		Description:    "",
		GivenStatus:    "",
		GivenDate:      time.Time{},
	}

	_ = petientRepo.conn.Create(&presc).GetErrors()

	diagnosis := entity.Diagnosis{
		ID:             0,
		PatientId:      usr.PatientId,
		PatientName:    usr.PatientName,
		DoctorId:       usr.DoctorId,
		LaboratoristId: usr.ApprovedBy,
		Description:    "",
		DiagonosesDate: time.Time{},
		Reponse:        "",
		ResponseDate:   time.Time{},
	}
	_ = petientRepo.conn.Create(&diagnosis).GetErrors()

}
