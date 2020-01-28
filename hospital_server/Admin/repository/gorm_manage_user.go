package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)


// ManageProfileRepository implements Admin.ManageProfileRepository interface
type UserRepository struct {
	conn *gorm.DB
}

// ManageProfileRepository returns new object of ManageProfileRepository
func NewUserRepository(db *gorm.DB) Admin.UserRepository {
	return &UserRepository{conn: db}
}

// Doctors return all doctors stored in the databasee
func (userRepo *UserRepository) Users() ([]entity.User, []error) {
	prfs := []entity.User{}
	errs := userRepo.conn.Find(&prfs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prfs, errs
}

// Doctor retrieves a doctor from the database by its id
func (userRepo *UserRepository) User(id uint) (*entity.User, []error) {
	prfs := entity.User{}
	errs := userRepo.conn.First(&prfs, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &prfs, errs
}

// UpdateDoctor updats a given doctor in the database
func (userRepo *UserRepository) UpdateUser(user *entity.User) (*entity.User, []error) {
	prf := user
	errs := userRepo.conn.Save(prf).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prf, errs
}

// DeleteDoctor deletes a given doctor from the database
func (userRepo *UserRepository) DeleteUser(id uint) (*entity.User, []error) {
	prf, errs := userRepo.User(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = userRepo.conn.Delete(prf, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prf, errs
}

// StoreDoctor stores a given doctor in the database
func (userRepo *UserRepository) StoreUser(user *entity.User) (*entity.User, []error) {
	prf := user
	errs := userRepo.conn.Create(prf).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return prf, errs
}
// PhoneExists check if a given phone number is found
func (userRepo *UserRepository) PhoneExists(phone string) bool {
	user := entity.User{}
	errs := userRepo.conn.Find(&user, "phone=?", phone).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

// EmailExists check if a given email is found
func (userRepo *UserRepository) EmailExists(email string) bool {
	user := entity.User{}
	errs := userRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

// UserRoles returns list of application roles that a given user has
func (userRepo *UserRepository) UserRoles(user *entity.User) ([]entity.Role, []error) {
	fmt.Println("thise is the userRoles gorm")
	userRoles := []entity.Role{}
	errs := userRepo.conn.Model(user).Related(&userRoles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	
	return userRoles, errs
}
// UserByEmail retrieves a user by its email address from the database
func (userRepo *UserRepository) UserByEmail(email string) (*entity.User, []error) {
	user := entity.User{}
	errs := userRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}