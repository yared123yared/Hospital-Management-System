package handler

import (
	"fmt"
	entity2 "github.com/getach1/web1/web1_group_project/hospital_client/delivery/entity"
	"github.com/getach1/web1/web1_group_project/hospital_client/delivery/http/data"
	"html/template"
	"strconv"
	"time"
	"net/http"
)

type AdminPatientHandler struct {
	tmpl        *template.Template
}


func NewPatientHandler(T *template.Template) *AdminPatientHandler {
	return &AdminPatientHandler{tmpl: T}
}


func (ph *AdminPatientHandler)Profile(w http.ResponseWriter, _ *http.Request) {
	//petient:=Petient{1 ,"Getachew","Tebikew","prescription.png","Addis Ababa","xy@z.com","+1113444",time.Now()}
	petient := entity2.Petient{}
	var err error
	petient, err = data.FetchPetient(11)
	data.CheckErr(err)
	err = ph.tmpl.ExecuteTemplate(w, "patient.view.profile.update", petient)
	data.CheckErr(err)
}

func (ph *AdminPatientHandler)Doctors(w http.ResponseWriter, _ *http.Request) {
	doctors := []entity2.Doctor{}
	doctors, _ = data.FetchDoctors()
	err := ph.tmpl.ExecuteTemplate(w, "patient.view.doctor", doctors)
	data.CheckErr(err)
}
func (ph *AdminPatientHandler) Appointment(w http.ResponseWriter, _ *http.Request) {
	petient := entity2.Petient{}
	petient, _ = data.FetchPetient(11)
	err := ph.tmpl.ExecuteTemplate(w, "patient.view.appointment", petient)
	data.CheckErr(err)
}
func (ph *AdminPatientHandler) Prescription(w http.ResponseWriter, _ *http.Request) {
	petient := entity2.Petient{}
	petient, _ = data.FetchPetient(11)
	err := ph.tmpl.ExecuteTemplate(w, "patient.view.prescription", petient)
	fmt.Println(petient.Appointment)
	data.CheckErr(err)
}

func (ph *AdminPatientHandler) Request(w http.ResponseWriter, r *http.Request) {
	petient := entity2.Petient{}
	petient, _ = data.FetchPetient(11)
	err := ph.tmpl.ExecuteTemplate(w, "patient.view.request", petient)
	data.CheckErr(err)
}
func (ph *AdminPatientHandler) SendRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		fmt.Println("hello")
		id, err := strconv.Atoi(idRaw)
		data.CheckErr(err)
		user := entity2.Petient{}
		user, err = data.FetchPetient(id)
		data.CheckErr(err)
		user, _ = data.FetchPetient(id)
		err = ph.tmpl.ExecuteTemplate(w, "patient.view.request", user)
		data.CheckErr(err)
	} else if r.Method == http.MethodPost {
		user := entity2.Petient{}
		id, _ := strconv.Atoi(r.FormValue("id"))
		var err error
		user, err = data.FetchPetient(id)
		data.CheckErr(err)
		request := entity2.Request{
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
func (ph *AdminPatientHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		data.CheckErr(err)
		user, err := data.FetchPetient(id)
		data.CheckErr(err)
		err = ph.tmpl.ExecuteTemplate(w, "admin.user.update.layout", user)
		data.CheckErr(err)
	} else if r.Method == http.MethodPost {
		user := entity2.Petient{}
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
