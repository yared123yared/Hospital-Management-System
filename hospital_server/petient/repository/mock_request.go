package PetientRepository

import (
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// RequestGormRepo Implements the menu.RequestRepository interface
type MockRequestGormRepo struct {
	conn *gorm.DB
}

// NewRequestGormRepo creates a new object of RequestGormRepo
func NewMockRequestGormRepo(db *gorm.DB) petient.RequestRepository {
	return &RequestGormRepo{conn: db}
}

// Requests return all requests from the database
func (requestRepo *MockRequestGormRepo) Requests() ([]entity.Request, []error) {
	requests := []entity.Request{entity.RequestMock}
	return requests, nil
}

// Request retrieves a request by its id from the database
func (requestRepo *MockRequestGormRepo) Request(id uint) (*entity.Request, []error) {
	request := entity.RequestMock
	if id == 1 {
		return &request, nil
	}
	return &request, []error{errors.New("Not Found")}
}

/*
// UpdateRequest updates a given request in the database
func (requestRepo *RequestGormRepo) UpdateRequest(request *entity.Request) (*entity.Request, []error) {
	usr := request
	errs := requestRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteRequest deletes a given request from the database
func (requestRepo *RequestGormRepo) DeleteRequest(id uint) (*entity.Request, []error) {
	usr, errs := requestRepo.Request(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = requestRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
*/
// StoreRequest stores a new request into the database
func (requestRepo *MockRequestGormRepo) StoreRequest(request *entity.Request) (*entity.Request, []error) {
	req := entity.RequestMock
	return &req, nil
}
