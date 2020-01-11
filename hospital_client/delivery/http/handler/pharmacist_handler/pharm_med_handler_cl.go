package pharmacist_handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	pharmacistData "github.com/fasikawkn/web1_group_project-1/hospital_client/data/pharmacist"
	"github.com/fasikawkn/web1_group_project/hospital_server/entity"
)

func (ach *PharmProfHandler) CatHandler(w http.ResponseWriter, r *http.Request) {
	medicines, _ := pharmacistData.GetMedicines()
	ach.tmpl.ExecuteTemplate(w, "pharm_cat.layout", medicines)

}

func (ach *PharmProfHandler) AddNewCat(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		med := entity.Medicine{}
		med.CategoryName = r.FormValue("catname")
		med.MedicineName = r.FormValue("Medname")
		date := "2022-01-02"
		fmt.Println("Date s", date)
		// t, err := time.Parse("2002-01-06", date)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println("Date", t)
		med.ExpiredAt = time.Now()
		fmt.Println("catDate", r.FormValue("expireddate"))
		i, _ := strconv.ParseUint(r.FormValue("amount"), 10, 64)
		med.Amount = uint(i)
		med.AddedBy = sesion
		// fmt.Println(med.AddedBy)
		pharmacistData.PostMedicine(&med)
		http.Redirect(w, r, "/cat", http.StatusSeeOther)

	} else {
		ach.tmpl.ExecuteTemplate(w, "pharm.cat.new.layout", nil)
	}
}

func (ach *PharmProfHandler) UpdateCat(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		parms, _ := pharmacistData.GetMedicine(uint(id))

		ach.tmpl.ExecuteTemplate(w, "pharm.cat.update.layout", parms)

	} else if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))

		medicine, _ := pharmacistData.GetMedicine(uint(id))
		fmt.Println("medinen", medicine)

		pharms := entity.Medicine{}

		pharms.ID = medicine.ID
		pharms.AddedBy = sesion

		pharms.CategoryName = r.FormValue("catname")
		pharms.MedicineName = r.FormValue("Medname")
		pharms.ExpiredAt = time.Now()
		i, _ := strconv.ParseUint(r.FormValue("amount"), 10, 64)
		pharms.Amount = uint(i)

		pharmacistData.PutMedicine(&pharms)

		http.Redirect(w, r, "/cat", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/cat", http.StatusSeeOther)
	}

}
func (ach *PharmProfHandler) DleteMedicine(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		pharmacistData.DeleteMedicine(uint(id))

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/cat", http.StatusSeeOther)

}
func (ach *PharmProfHandler) UpdateProv(w http.ResponseWriter, r *http.Request) {

	ach.tmpl.ExecuteTemplate(w, "pharm.prov.update.layout", nil)

}
