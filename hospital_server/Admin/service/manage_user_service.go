package service

import (
	"github.com/web1_group_project/hospital_server/Admin"
	"github.com/web1_group_project/hospital_server/entity"
)

// ManageDoctorsService implements Admin.ManageDoctorsService interface

type UserService struct {
	userRepo Admin.UserRepository
}

//NewManageDoctorsService returns new ManageDoctorsService object
func NewUserService(mdsLR Admin.UserRepository) Admin.UserService {
	return &UserService{userRepo: mdsLR}
}

// Doctors returns list of Doctors
func (mpSer *UserService) Users() ([]entity.User, []error) {
	prf, errs := mpSer.userRepo.Users()
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// Doctor returns a Doctor object with a given id
func (mpSer *UserService) User(id uint) (*entity.User, []error) {
	prf, errs := mpSer.userRepo.User(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// UpdateDoctor updates a Doctor with new data
func (mpSer *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	prf, errs := mpSer.userRepo.UpdateUser(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// DeleteDoctor delete a Doctor by its id
func (mpSer *UserService) DeleteUser(id uint) (*entity.User, []error) {
	prf, errs := mpSer.userRepo.DeleteUser(id)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}

// StoreDoctor persists new Doctor information
func (mpSer *UserService) StoreUser(user *entity.User) (*entity.User, []error) {

	prf, errs := mpSer.userRepo.StoreUser(user)
	if len(errs) > 1 {
		return nil, errs
	}
	return prf, errs
}
// UserByEmail retrieves an application user by its email address
func (us *UserService) UserByEmail(email string) (*entity.User, []error) {
	usr, errs := us.userRepo.UserByEmail(email)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
// PhoneExists check if there is a user with a given phone number
func (us *UserService) PhoneExists(phone string) bool {
	exists := us.userRepo.PhoneExists(phone)
	return exists
}

// EmailExists checks if there exist a user with a given email address
func (us *UserService) EmailExists(email string) bool {
	exists := us.userRepo.EmailExists(email)
	return exists
}

// UserRoles returns list of roles a user has
func (us *UserService) UserRoles(user *entity.User) ([]entity.Role, []error) {
	userRoles, errs := us.userRepo.UserRoles(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}