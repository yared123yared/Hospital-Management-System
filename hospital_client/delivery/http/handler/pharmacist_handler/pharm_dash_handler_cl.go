package pharmacist_handler

import (
	"fmt"
	"github.com/web1_group_project/hospital_client/entity"
	"github.com/web1_group_project/hospital_client/session"

	"html/template"
	"net/http"

	pharmacistData "github.com/web1_group_project/hospital_client/data/pharmacist"
)

// AdminCategoryHandler handles category handler admin requests
type PharmProfHandler struct {
	tmpl *template.Template
}

// NewAdminCategoryHandler initializes and returns new AdminCateogryHandler
func NewPharmTempHandler(T *template.Template) *PharmProfHandler {
	return &PharmProfHandler{tmpl: T}
}

func (ach *PharmProfHandler) Index(w http.ResponseWriter, r *http.Request) {
	ach.tmpl.ExecuteTemplate(w, "main.layout", nil)

}

func (ach *PharmProfHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	var sesion, _ = session.IsLogged(r)
	fmt.Println("Dashboard")
	num := pharmacistData.GetMedsNumber(uint(sesion))
	num2 := pharmacistData.GetPrescsNumber(uint(sesion))
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
