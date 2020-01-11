package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/web1_group_project/hospital_server/Doctor"
	"github.com/web1_group_project/hospital_server/entity"
	//"github.com/yaredsolomon/webProgram1/hospital/request"
)

// AppointmentGormRepo Implements the request.AppointmentRepository interface
type GeneralGormRepo struct {
	conn *gorm.DB
}

// NewGeneralGormRepo creates a new object of GeneralGormRepo
func NewGeneralGormRepo(db *gorm.DB) Doctor.GeneralRepository {
	return &GeneralGormRepo{conn: db}
}

// Pharmacists return all Pharmacists from the database
func (generalRepo *GeneralGormRepo) Pharmacists() ([]entity.Pharmacist, []error) {
	pharmacists := []entity.Pharmacist{}
	errs := generalRepo.conn.Find(&pharmacists).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pharmacists, errs
}
func (generalRepo *GeneralGormRepo) Laboratorists() ([]entity.Laboratorist, []error) {
	laboratorists := []entity.Laboratorist{}
	errs := generalRepo.conn.Find(&laboratorists).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return laboratorists, errs
}
func (generalRepo *GeneralGormRepo) Users(id int, password string) (*entity.Profile, []error) {

	fmt.Println(" thise is the gorm method")
	fmt.Print(id)
	users := entity.Profile{}
	errs := generalRepo.conn.Where("id =?", id).First(&users).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	fmt.Println("user input")
	fmt.Println(password)
	fmt.Println("fetched data")
	fmt.Println(users.Password)

	if users.Password == password {
		fmt.Println("thise is the gorm result")
		fmt.Println(users)
		return &users, errs
	}
	return nil, nil

}
