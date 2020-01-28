package Pharmacist_handler

import (
	"fmt"
	"html/template"
	"net/http"

	Pharmacist_data "github.com/web1_group_project/hospital_client/data/pharmacist"
	"github.com/web1_group_project/hospital_client/delivery/http/handler"
	"github.com/web1_group_project/hospital_client/entity"
	//"github.com/web1_group_project/hospital_client/session"
	//pharmacistData "github.com/web1_group_project/hospital_client/data/pharmacist/"
)

// AdminCategoryHandler handles category handler admin requests
type PharmProfHandler struct {
	tmpl              *template.Template
	UserHandler       *handler.UserHandler
	LogedInPharmacist *entity.Pharmacist
	csrfSignKey       []byte
}

// NewAdminCategoryHandler initializes and returns new AdminCateogryHandler
func NewPharmTempHandler(T *template.Template, userHandler *handler.UserHandler, csKey []byte) *PharmProfHandler {
	return &PharmProfHandler{
		tmpl:        T,
		UserHandler: userHandler,
		csrfSignKey: csKey,
	}
}

func (ach *PharmProfHandler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am about to fech data")
	var id uint = ach.UserHandler.LoggedInUser.ID

	pharmacist, err := Pharmacist_data.GetPharmacist(id)

	fmt.Println(ach.UserHandler.LoggedInUser.ID)
	if err != nil {
		panic(err)
	}

	ach.LogedInPharmacist = pharmacist
	fmt.Println(pharmacist)
	ach.tmpl.ExecuteTemplate(w, "main.layout", nil)

}

func (ach *PharmProfHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	// var sesion, _ = session.IsLogged(r)
	var sesion = 3
	fmt.Println("Dashboard")
	num := Pharmacist_data.GetMedsNumber(uint(sesion))
	num2 := Pharmacist_data.GetPrescsNumber(uint(sesion))
	var num3 int = num / 12
	var num4 int = num2 / 12
	dash := entity.Dash{

		Annual_one:  num,
		Monthly_one: num3,
		Annual_two:  num2,
		Monthly_two: num4,
	}
	ach.tmpl.ExecuteTemplate(w, "main.layout", dash)

}
