package pharmacist_handler

import (
	"fmt"

	"html/template"
	"net/http"

	pharmacistData "github.com/fasikawkn/web1_group_project-1/hospital_client/data/pharmacist"

	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
	"github.com/fasikawkn/web1_group_project/hospital_server/session"
)

var sesion uint = session.GetPharmSession()

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
	fmt.Println("Dashboard")
	num := pharmacistData.GetMedsNumber(sesion)
	num2 := pharmacistData.GetPrescsNumber(sesion)
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
