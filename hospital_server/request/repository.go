package request

import "github.com/getach1/web1/web1_group_project/hospital_server/entity"

type RequestRepository interface {
	Requests() ([]entity.Request, []error)
	Request(id uint) (*entity.Request, []error)
	UpdateRequest(request *entity.Request) (*entity.Request, []error)
	DeleteRequest(id uint) (*entity.Request, []error)
	StoreRequest(request *entity.Request) (*entity.Request, []error)
}
