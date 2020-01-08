package request

import "github.com/getach1/web1/hospital/entity"

type RequestService interface {
	Requests() ([]entity.Request, []error)
	Request(id uint) (*entity.Request, []error)
	UpdateRequest(request *entity.Request) (*entity.Request, []error)
	DeleteRequest(id uint) (*entity.Request, []error)
	StoreRequest(request *entity.Request) (*entity.Request, []error)
}
