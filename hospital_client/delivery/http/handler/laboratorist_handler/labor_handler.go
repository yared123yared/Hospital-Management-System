package laboratorist_handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fasikawkn/web1_group_project/hospital_server/session"
	laborData "github.com/web1_group_project/hospital_client/data/laboratorist"
	"github.com/web1_group_project/hospital_client/entity"
)

var sesion uint = session.GetLaborSession()

// LaborProfHandler handles category handler admin requests
type LaborProfHandler struct {
	tmpl *template.Template
}

// NewAdminCategoryHandler initializes and returns new AdminCateogryHandler
func NewLaborTempHandler(T *template.Template) *LaborProfHandler {
	return &LaborProfHandler{tmpl: T}
}
func (ach *LaborProfHandler) LaborDashHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Dashboard")
	num := laborData.GetDiagsNumber(sesion)
	num2 := laborData.GetPrescsNumber(sesion)
	var num3 int = num / 12
	var num4 int = num2 / 12
	dash := entity.Dash{

		Annual_one:  num,
		Monthly_one: num3,
		Annual_two:  num2,
		Monthly_two: num4,
	}

	ach.tmpl.ExecuteTemplate(w, "labor.home.layout", dash)
}

func (ach *LaborProfHandler) LaborDiagnosisHandler(w http.ResponseWriter, r *http.Request) {
	diagnosiss, _ := laborData.GetDiagnosiss()
	log.Println("labor handelr diagnosis")
	ach.tmpl.ExecuteTemplate(w, "labor.diag.home.layout", diagnosiss)
}

func (ach *LaborProfHandler) LaborProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("labor handelr profile")
	labratorist, _ := laborData.GetLaboratorist(sesion)
	ach.tmpl.ExecuteTemplate(w, "labor.prof.layout", labratorist)
}

func (ach *LaborProfHandler) LaborProfileUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		parms, _ := laborData.GetLaboratorist(uint(id))

		ach.tmpl.ExecuteTemplate(w, "labor.prof.update.layout", parms)

	} else if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))

		laboratorist, _ := laborData.GetLaboratorist(uint(id))

		pharms := entity.Laboratorist{}

		pharms.ID = laboratorist.ID
		pharms.Diagnosis = laboratorist.Diagnosis
		pharms.Uuid = laboratorist.Uuid

		pharms.Profile.Role = laboratorist.Profile.Role

		stt, _ := strconv.ParseUint(r.FormValue("id"), 10, 64)
		pharms.Profile.ID = uint(stt)
		pharms.Profile.FullName = r.FormValue("fullname")
		pharms.Profile.Phone = r.FormValue("phone")
		pharms.Profile.Address = r.FormValue("address")
		if r.FormValue("image") == "" {
			pharms.Profile.Image = r.FormValue("image2")
		} else {
			pharms.Profile.Image = r.FormValue("image")

		}
		pharms.Profile.Sex = r.FormValue("sex")
		pharms.Profile.Email = r.FormValue("email")
		pharms.Profile.BirthDate = time.Now()
		fmt.Println(r.FormValue("birthdate"))
		pharms.Profile.Description = r.FormValue("description")

		laborData.PutLaboratorist(&pharms)

		http.Redirect(w, r, "/profileLabor", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/profileLabor", http.StatusSeeOther)
	}
}

func (ach *LaborProfHandler) LaborDiagnosisUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}
		parms, _ := laborData.GetDiagnosis(uint(id))
		log.Println("laboratorist", parms)
		ach.tmpl.ExecuteTemplate(w, "labor.diag.update.layout", parms)

	} else if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))

		diagns, _ := laborData.GetDiagnosis(uint(id))

		pharms := entity.Diagnosis{}

		pharms.ID = diagns.ID

		pharms.PatientName = diagns.PatientName
		pharms.PatientId = diagns.PatientId
		pharms.DoctorId = diagns.DoctorId
		pharms.LaboratoristId = sesion
		pharms.DiagonosesDate = time.Now()
		pharms.Description = r.FormValue("description")
		pharms.Reponse = r.FormValue("response")
		pharms.ResponseDate = time.Now()
		laborData.PutDiagnosis(&pharms)

		http.Redirect(w, r, "/diagnosisLabor", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/diagnosisLabor", http.StatusSeeOther)
	}
}
