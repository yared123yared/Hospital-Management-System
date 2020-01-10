package repository

import (
	"github.com/getach1/web1/web1_group_project/hospital_server/entity"
	"github.com/getach1/web1/web1_group_project/hospital_server/request"
	"github.com/jinzhu/gorm"
)

// RequestGormRepo Implements the menu.RequestRepository interface
type RequestGormRepo struct {
	conn *gorm.DB
}

// NewRequestGormRepo creates a new object of RequestGormRepo
func NewRequestGormRepo(db *gorm.DB) request.RequestRepository {
	return &RequestGormRepo{conn: db}
}

// Requests return all requests from the database
func (requestRepo *RequestGormRepo) Requests() ([]entity.Request, []error) {
	requests := []entity.Request{}
	errs := requestRepo.conn.Find(&requests).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return requests, errs
}

// Request retrieves a request by its id from the database
func (requestRepo *RequestGormRepo) Request(id uint) (*entity.Request, []error) {
	request := entity.Request{}
	errs := requestRepo.conn.Debug().Find(&request, "id=?", id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &request, errs
}

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

// StoreRequest stores a new request into the database
func (requestRepo *RequestGormRepo) StoreRequest(request *entity.Request) (*entity.Request, []error) {
	usr := request
	errs := requestRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
