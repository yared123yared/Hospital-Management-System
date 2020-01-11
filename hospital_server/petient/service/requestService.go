package PetientService

import (
	"github.com/web1_group_project/hospital_server/entity"
	"github.com/web1_group_project/hospital_server/petient"
)

// RequestService implements menu.RequestService interface
type RequestService struct {
	requestRepo petient.RequestRepository
}

// NewRequestService  returns a new RequestService object
func NewRequestService(requestRepository petient.RequestRepository) petient.RequestService {
	return &RequestService{requestRepo: requestRepository}
}

// Requests returns all stored application requests
func (us *RequestService) Requests() ([]entity.Request, []error) {
	usrs, errs := us.requestRepo.Requests()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// Request retrieves an application request by its id
func (us *RequestService) Request(id uint) (*entity.Request, []error) {
	usr, errs := us.requestRepo.Request(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreRequest stores a given application request
func (us *RequestService) StoreRequest(request *entity.Request) (*entity.Request, []error) {
	usr, errs := us.requestRepo.StoreRequest(request)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
