package handler

import (
	"fmt"
	entity "github.com/getach1/web1/Project/entities"
	"github.com/getach1/web1/Project/entities/data"
	"html/template"
	"strconv"
	"time"

	"net/http"
)

var tmpl = template.Must(template.ParseGlob("template/*"))

func Profile(w http.ResponseWriter, _ *http.Request) {
	//petient:=Petient{1 ,"Getachew","Tebikew","prescription.png","Addis Ababa","xy@z.com","+1113444",time.Now()}
	petient := entity.Petient{}
	var err error
	petient, err = data.FetchPetient(11)
	data.CheckErr(err)
	err = tmpl.ExecuteTemplate(w, "patient.view.profile.update", petient)
	data.CheckErr(err)
}

func Doctors(w http.ResponseWriter, _ *http.Request) {
	doctors := []entity.Doctor{}
	doctors, _ = data.FetchDoctors()
	err := tmpl.ExecuteTemplate(w, "patient.view.doctor", doctors)
	data.CheckErr(err)
}
func Appointment(w http.ResponseWriter, _ *http.Request) {
	petient := entity.Petient{}
	petient, _ = data.FetchPetient(11)
	err := tmpl.ExecuteTemplate(w, "patient.view.appointment", petient)
	data.CheckErr(err)
}
func Prescription(w http.ResponseWriter, _ *http.Request) {
	petient := entity.Petient{}
	petient, _ = data.FetchPetient(11)
	err := tmpl.ExecuteTemplate(w, "patient.view.prescription", petient)
	fmt.Println(petient.Appointment)
	data.CheckErr(err)
}

func Request(w http.ResponseWriter, r *http.Request) {
	petient := entity.Petient{}

	petient, _ = data.FetchPetient(11)
	err := tmpl.ExecuteTemplate(w, "patient.view.request", petient)
	data.CheckErr(err)
}
func SendRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		fmt.Println("hello")
		id, err := strconv.Atoi(idRaw)
		data.CheckErr(err)
		user := entity.Petient{}
		user, err = data.FetchPetient(id)
		data.CheckErr(err)
		user, _ = data.FetchPetient(id)
		err = tmpl.ExecuteTemplate(w, "patient.view.request", user)
		data.CheckErr(err)
	} else if r.Method == http.MethodPost {
		user := entity.Petient{}
		id, _ := strconv.Atoi(r.FormValue("id"))
		var err error
		user, err = data.FetchPetient(id)
		data.CheckErr(err)
		request := entity.Request{
			ID:            1,
			DoctorId:      1,
			PatientId:     uint(user.ID),
			ApproveStatus: "waiting",
			ApprovedBy:    0,
		}
		user.Request = append(user.Request, request)
		fmt.Println("POST: request sent")
		fmt.Println(request)
		data.UpdateProfile(user)
		fmt.Println(user)
		http.Redirect(w, r, "/request", http.StatusSeeOther)

	} else {
		fmt.Println("Not sent")
		http.Redirect(w, r, "/request", http.StatusSeeOther)
	}
}
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		data.CheckErr(err)
		user, err := data.FetchPetient(id)
		data.CheckErr(err)
		err = tmpl.ExecuteTemplate(w, "admin.user.update.layout", user)
		data.CheckErr(err)
	} else if r.Method == http.MethodPost {
		user := entity.Petient{}
		id, _ := strconv.Atoi(r.FormValue("id"))
		var err error
		user, err = data.FetchPetient(id)
		data.CheckErr(err)
		fmt.Println(user.ID, user.Profile.FullName, "post")
		user.Profile.FullName = r.FormValue("full_name")
		user.Profile.Address = r.FormValue("address")
		user.Profile.Email = r.FormValue("email")
		user.Profile.Phone = r.FormValue("phone")
		user.Profile.BirthDate, _ = time.Parse(time.RFC3339, r.FormValue("birth_date"))
		fmt.Println(user.Profile.BirthDate)
		data.UpdateProfile(user)
		fmt.Println("Updated")
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	} else {
		fmt.Println("Not Updated")
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}
