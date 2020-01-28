 package Pharmacist_handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/web1_group_project/hospital_client/entity"
	//"github.com/web1_group_project/hospital_client/session"
	pharmacistData "github.com/web1_group_project/hospital_client/data/pharmacist"
)

func (ach *PharmProfHandler) ProHandler(w http.ResponseWriter, r *http.Request) {
	var sesion=1
	pharmacist, _ := pharmacistData.GetPharmacist(uint(sesion))
	ach.tmpl.ExecuteTemplate(w, "pharm.prof.layout", pharmacist)

}

// PharmProfile handle requests on route /admin
func (ach *PharmProfHandler) PharmProfileUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		parms, _ := pharmacistData.GetPharmacist(uint(id))

		ach.tmpl.ExecuteTemplate(w, "pharm_prof.layout", parms)

	} else if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))

		pharmacist, _ := pharmacistData.GetPharmacist(uint(id))

		pharms := entity.Pharmacist{}

		pharms.ID = pharmacist.ID
		pharms.Medicine = pharmacist.Medicine
		pharms.Prescription = pharmacist.Prescription
		pharms.User.RoleId = pharmacist.User.RoleId

		stt, _ := strconv.ParseUint(r.FormValue("id"), 10, 64)
		pharms.User.ID = uint(stt)
		pharms.User.FullName = r.FormValue("fullname")
		pharms.User.Phone = r.FormValue("phone")
		pharms.User.Address = r.FormValue("address")
		if r.FormValue("image") == "" {
			pharms.User.Image = r.FormValue("image2")
		} else {
			pharms.User.Image = r.FormValue("image")

		}
		pharms.User.Sex = r.FormValue("sex")
		pharms.User.Email = r.FormValue("email")
		pharms.User.BirthDate = time.Now()
		fmt.Println(r.FormValue("birthdate"))
		pharms.User.Description = r.FormValue("description")
		pharmacistData.PutPharmacist(&pharms)

		http.Redirect(w, r, "/prof", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/prof", http.StatusSeeOther)
	}

}
