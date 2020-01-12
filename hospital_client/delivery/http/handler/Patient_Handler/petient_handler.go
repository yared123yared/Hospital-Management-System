package Patient_Handler

import (
	"fmt"
	data "github.com/getach1/web1/web1_group_project_old_new/hospital_client/data/Patient"
	"github.com/getach1/web1/web1_group_project_old_new/hospital_client/entity"

	"html/template"
	"net/http"
	"strconv"
	"time"
)

type AdminPatientHandler struct {
	tmpl *template.Template
}

func NewPatientHandler(T *template.Template) *AdminPatientHandler {
	return &AdminPatientHandler{tmpl: T}
}
func (ph *AdminPatientHandler) Profile(w http.ResponseWriter, _ *http.Request) {
	//petient:=Petient{1 ,"Getachew","Tebikew","prescription.png","Addis Ababa","xy@z.com","+1113444",time.Now()}
	petient := entity.Petient{}
	var err error
	petient, err = data.FetchPetient(1)
	data.CheckErr(err)
	err = ph.tmpl.ExecuteTemplate(w, "patient.view.profile", petient)
	data.CheckErr(err)
}

func (ph *AdminPatientHandler) Doctors(w http.ResponseWriter, _ *http.Request) {
	doctors := []entity.Doctor{}
	doctors,err:= data.FetchDoctors()
	data.CheckErr(err)
	user,err:=data.FetchPetient(2)
	doctList:=entity.DoctorsList{
		User:  user,
		Doctors: doctors,
	}
	err = ph.tmpl.ExecuteTemplate(w, "patient.view.doctor", doctList)
	data.CheckErr(err)
}
func (ph *AdminPatientHandler) Appointment(w http.ResponseWriter, _ *http.Request) {
	petient := entity.Petient{}
	petient, _ = data.FetchPetient(1)
	err := ph.tmpl.ExecuteTemplate(w, "patient.view.appointment", petient)
	data.CheckErr(err)
}
func (ph *AdminPatientHandler) Prescription(w http.ResponseWriter, _ *http.Request) {
	petient := entity.Petient{}
	petient, _ = data.FetchPetient(1)
	err := ph.tmpl.ExecuteTemplate(w, "patient.view.prescription", petient)
	fmt.Println(petient.Appointment)
	data.CheckErr(err)
}

func (ph *AdminPatientHandler) Request(w http.ResponseWriter, r *http.Request) {
	petient := entity.Petient{}
	petient, _ = data.FetchPetient(1)
	err := ph.tmpl.ExecuteTemplate(w, "patient.view.request", petient)
	data.CheckErr(err)
}
func (ph *AdminPatientHandler) NewRequest(w http.ResponseWriter, r *http.Request) {
	doctor := []entity.Doctor{}
	doctor, _ = data.FetchDoctors()
	admin := []entity.Admin{}
	admin,_= data.FetchAdmins()
	req:=entity.NewRequest{
		UserID:  2,
		Doctors: doctor,
		Admins:  admin,
	}
	err := ph.tmpl.ExecuteTemplate(w, "patient.new.request", req)
	data.CheckErr(err)
}
func (ph *AdminPatientHandler) SendRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		fmt.Println("hello")
		id, err := strconv.Atoi(idRaw)
		data.CheckErr(err)
		user := entity.Petient{}
		user, err = data.FetchPetient(id)
		data.CheckErr(err)
		user, _ = data.FetchPetient(id)
		err = ph.tmpl.ExecuteTemplate(w, "patient.view.request", user)
		data.CheckErr(err)
	} else if r.Method == http.MethodPost {
		user := entity.Petient{}
		id, _ := strconv.Atoi(r.FormValue("id"))
		var err error
		user, err = data.FetchPetient(id)
		data.CheckErr(err)
		docid,err:=strconv.Atoi(r.FormValue("doctor_id"))
		data.CheckErr(err)
		adid,err:=strconv.Atoi(r.FormValue("admin_id"))
		data.CheckErr(err)
		request := entity.Request{
			ID:            0,
			DoctorId: uint(docid),
			PatientId:     uint(user.ID),
			PatientName: user.Profile.FullName,
			ApproveStatus: "waiting",
			ApprovedBy:    uint(adid),
		}
		user.Request = append(user.Request, request)
		fmt.Println("POST: request sent")
		fmt.Println(request)

		data.UpdateProfile(user)

		fmt.Println(user)
		http.Redirect(w, r, "/request", http.StatusSeeOther)

	} else {
		fmt.Println("Not sent")
		http.Redirect(w, r, "/request/new", http.StatusSeeOther)
	}
}
func (ph *AdminPatientHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		data.CheckErr(err)
		user, err := data.FetchPetient(id)
		data.CheckErr(err)
		err = ph.tmpl.ExecuteTemplate(w, "patient.view.profile.update", user)
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
		http.Redirect(w, r, "/profile/update", http.StatusSeeOther)
	}
}
