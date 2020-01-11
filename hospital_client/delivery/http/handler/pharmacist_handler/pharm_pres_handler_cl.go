package pharmacist_handler

import (
	"net/http"
	"strconv"
	"time"

	pharmacistData "github.com/fasikawkn/web1_group_project-1/hospital_client/data/pharmacist"
	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
)

func (ach *PharmProfHandler) Prescription(w http.ResponseWriter, r *http.Request) {
	prescs, _ := pharmacistData.Prescriptions()
	ach.tmpl.ExecuteTemplate(w, "pharm.pres.layout", prescs)
}
func (ach *PharmProfHandler) DeletePrescription(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		pharmacistData.DeletePrescription(uint(id))

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/prescription", http.StatusSeeOther)
	ach.tmpl.ExecuteTemplate(w, "pharm.prov.update.layout", nil)

}

func (ach *PharmProfHandler) PrescriptionUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		parms, _ := pharmacistData.GetPrescription(uint(id))

		ach.tmpl.ExecuteTemplate(w, "pharm.pres.update.layout", parms)

	} else if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))

		presc, _ := pharmacistData.GetPrescription(uint(id))

		pharms := entity.Prescription{}

		pharms.ID = presc.ID
		pharms.PatientName = presc.PatientName
		pharms.PatientId = presc.PatientId
		pharms.DoctorId = presc.DoctorId
		pharms.PrescribedDate = time.Now()
		pharms.MedicineName = presc.MedicineName
		pharms.Description = r.FormValue("description")
		pharms.GivenStatus = r.FormValue("givenstatus")
		pharms.PhrmacistId = sesion
		pharmacistData.PutPrescription(&pharms)

		http.Redirect(w, r, "/prescription", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/prescription", http.StatusSeeOther)
	}
	// ach.tmpl.ExecuteTemplate(w, "pharm.pres.update.layout", nil)
}
