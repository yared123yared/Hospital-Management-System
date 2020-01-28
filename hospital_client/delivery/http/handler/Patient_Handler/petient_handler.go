package Patient_Handler

import (
	//"context"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	//"strconv"
	"time"

	//"golang.org/x/crypto/bcrypt"

	"github.com/web1_group_project/hospital_client/data"
	petient_data "github.com/web1_group_project/hospital_client/data/Patient"
	"github.com/web1_group_project/hospital_client/delivery/http/handler"
	"github.com/web1_group_project/hospital_client/entity"
	"github.com/web1_group_project/hospital_client/form"
	"github.com/web1_group_project/hospital_client/rtoken"
)

type AdminPatientHandler struct {
	tmpl           *template.Template
	UserHandler    *handler.UserHandler
	LogedInPatient *entity.Petient
	csrfSignKey    []byte
}

type contextKey string

var ctxPatientSessionKey = contextKey("sign_in_patient_session")

func (ph *AdminPatientHandler) getPetient() {

	fmt.Println("i am about to fech data")

	patient, err := petient_data.FetchPetient2(ph.UserHandler.LoggedInUser.ID)
	if err != nil {
		panic(err)
	}

	ph.LogedInPatient = &patient
	return

}
func NewPatientHandler(T *template.Template, userHandler *handler.UserHandler, csKey []byte) *AdminPatientHandler {
	return &AdminPatientHandler{
		tmpl:        T,
		UserHandler: userHandler,
		csrfSignKey: csKey,
	}

}

func (ph *AdminPatientHandler) checkAdmin(rs int) bool {
	if rs == 1 {
		return true
	}
	return false
}

func (ph *AdminPatientHandler) Profile(w http.ResponseWriter, _ *http.Request) {

	petient := *ph.LogedInPatient
	fmt.Println(petient)
	err := ph.tmpl.ExecuteTemplate(w, "patient.view.profile", petient)
	petient_data.CheckErr(err)
}

func (ph *AdminPatientHandler) Doctors(w http.ResponseWriter, r *http.Request) {
	var doctors []entity.Doctor
	doctors, err := petient_data.FetchDoctors()
	petient_data.CheckErr(err)
	doctList := struct {
		User    entity.Petient
		Doctors []entity.Doctor
	}{
		User:    *ph.LogedInPatient,
		Doctors: doctors,
	}
	err = ph.tmpl.ExecuteTemplate(w, "patient.view.doctor", doctList)
	petient_data.CheckErr(err)
}

func (ph *AdminPatientHandler) Appointment(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("i am about to fech data")

	patient, err := petient_data.FetchPetient2(ph.UserHandler.LoggedInUser.ID)
	fmt.Println(patient)
	if err != nil {
		panic(err)
	}

	ph.LogedInPatient = &patient

	petient := *ph.LogedInPatient
	err = ph.tmpl.ExecuteTemplate(w, "patient.view.appointment", petient)
	petient_data.CheckErr(err)
}

func (ph *AdminPatientHandler) Prescription(w http.ResponseWriter, _ *http.Request) {
	petient := *ph.LogedInPatient
	err := ph.tmpl.ExecuteTemplate(w, "patient.view.prescription", petient)
	// fmt.Println(petient.Appointment)
	petient_data.CheckErr(err)
}

func (ph *AdminPatientHandler) Request(w http.ResponseWriter, _ *http.Request) {
	petient := *ph.LogedInPatient

	err := ph.tmpl.ExecuteTemplate(w, "patient.view.request", petient)
	petient_data.CheckErr(err)
}

// func (ph *AdminPatientHandler) SendRequest(w http.ResponseWriter, r *http.Request) {
// 	petient:=*ph.UserHandler.LoggedInUser

//   if r.Method == http.MethodGet {
//     err := ph.tmpl.ExecuteTemplate(w, "patient.view.request", petient)
//     petient_data.CheckErr(err)
//   } else if r.Method == http.MethodPost {
//     user:=*ph.UserHandler.LoggedInUser
//     docid,err:=strconv.Atoi(r.FormValue("doctor_id"))
//     petient_data.CheckErr(err)
//     adid,err:=strconv.Atoi(r.FormValue("admin_id"))
//     petient_data.CheckErr(err)
//     request := entity.Request{
//       ID:            0,
//       DoctorId: uint(docid),
//       PatientId:     uint(user.ID),
//       PatientName: user.FullName,
//       ApproveStatus: "waiting",
//       ApprovedBy:    uint(adid),
//     }
//    // user.Request = append(user.Request, request)
//     petient_data.UpdateProfile(user)
//     http.Redirect(w, r, "/request", http.StatusSeeOther)
//   } else {
//     fmt.Println("Not sent")
//     http.Redirect(w, r, "/request_new", http.StatusSeeOther)
//   }
// }

func (ph *AdminPatientHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("thise is the petient calsspatient/profile/updatessssssssssssssssssssssssssssssssssssssssssssssssssssss")
	token, err := rtoken.CSRFToken(ph.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		patient := *ph.LogedInPatient
		values := url.Values{}
		values.Add("full_name", patient.User.FullName)
		values.Add("email", patient.User.Email)
		values.Add("phone", patient.User.Phone)
		values.Add("image", patient.User.Image)
		values.Add("password", patient.User.Password)
		values.Add("sex", patient.User.Sex)
		upPatient := struct {
			CSRF    string
			Values  url.Values
			VErrors form.ValidationErrors
			Patient entity.Petient
		}{
			CSRF:    token,
			Values:  values,
			VErrors: form.ValidationErrors{},
			Patient: patient,
		}
		err = ph.tmpl.ExecuteTemplate(w, "patient.view.profile.update", upPatient)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Println("patient update method")
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		updatePatientForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		updatePatientForm.Required("full_name", "email", "phone", "description", "image", "birth_date")
		updatePatientForm.MinLength("description", 10)
		updatePatientForm.MatchesPattern("email", form.EmailRX)
		updatePatientForm.MatchesPattern("phone", form.PhoneRX)
		// updatePatientForm.MatchesPattern("full_name",form.NameRX)
		// updatePatientForm.MatchesPattern("address",form.AddressRX)
		updatePatientForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		fmt.Println(updatePatientForm.Values)
		fmt.Println(updatePatientForm.VErrors)

		if !updatePatientForm.Valid() {
			ph.tmpl.ExecuteTemplate(w, "patient.view.profile.update", updatePatientForm)
			return
		}
		petient := *ph.LogedInPatient

		user := &petient

		fmt.Println(user.ID, user.User.FullName, "post")
		user.User.FullName = r.FormValue("full_name")
		user.User.Address = r.FormValue("address")
		user.User.Email = r.FormValue("email")
		user.User.Phone = r.FormValue("phone")
		user.User.Address = r.FormValue("address")
		user.User.BirthDate, _ = time.Parse(time.RFC3339, r.FormValue("birth_date"))
		user.User.Image = r.FormValue("image")
		user.User.Description = r.FormValue("description")
		mf, fh, err := r.FormFile("image")
		if err == nil {
			user.User.Image = fh.Filename
			err = writeFile(&mf, user.User.Image)
		}
		if mf != nil {
			defer mf.Close()
		}
		petient_data.UpdateProfile(user)
		*ph.LogedInPatient = *user
		http.Redirect(w, r, "/patient/profile", http.StatusSeeOther)
		return
	}
}

func writeFile(mf *multipart.File, fname string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, "../ui", "assets", "img", fname)
	image, err := os.Create(path)
	if err != nil {
		return err
	}
	defer image.Close()
	io.Copy(image, *mf)
	return nil
}
func (ph *AdminPatientHandler) NewRequest(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ph.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		var doctor []entity.Doctor
		doctor, _ = petient_data.FetchDoctors()
		var admin []entity.Admin
		admin, _ = petient_data.FetchAdmins()
		labs, _ := data.Laboratorists()
		pharma, _ := data.Pharmacists()
		user := *ph.LogedInPatient
		req := struct {
			CSRF    string
			UserID  int
			Labor   []entity.Laboratorist
			Pharm   []entity.Pharmacist
			Doctors []entity.Doctor
			Admins  []entity.Admin
		}{
			CSRF:    token,
			UserID:  int(user.ID),
			Labor:   *labs,
			Pharm:   *pharma,
			Doctors: doctor,
			Admins:  admin,
		}

		err = ph.tmpl.ExecuteTemplate(w, "patient.new.request", req)
		petient_data.CheckErr(err)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Println("GEtetet request")
		user := *ph.LogedInPatient
		user, err := petient_data.FetchPetient(int(user.ID))
		petient_data.CheckErr(err)
		fmt.Println("GEtetet request")
		docid, err := strconv.Atoi(r.FormValue("doctor_id"))
		petient_data.CheckErr(err)
		adid, err := strconv.Atoi(r.FormValue("admin_id"))
		petient_data.CheckErr(err)
		petient_data.CheckErr(err)
		request := entity.Request{
			ID:            0,
			DoctorId:      uint(docid),
			PatientId:     uint(user.ID),
			PatientName:   user.User.FullName,
			ApproveStatus: "waiting",
			ApprovedBy:    uint(adid),
		}
		fmt.Println("GEtetet request")

		user.Request = append(user.Request, request)
		petient_data.UpdateProfile(&user)
		fmt.Println("New Request", request)
		http.Redirect(w, r, "/patient/request", http.StatusSeeOther)
	} else {
		fmt.Println("Not sent")
		http.Redirect(w, r, "/patient/request/new", http.StatusSeeOther)
	}

}
