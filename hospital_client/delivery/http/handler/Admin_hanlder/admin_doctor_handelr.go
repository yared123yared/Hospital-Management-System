package Admin_hanlder

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/web1_group_project/hospital_client/data/Admin"
	"github.com/web1_group_project/hospital_client/delivery/http/handler"
	"github.com/web1_group_project/hospital_client/entity"
	"github.com/web1_group_project/hospital_client/form"
	"github.com/web1_group_project/hospital_client/rtoken"
	"golang.org/x/crypto/bcrypt"
)

type AdminTempHandler struct {
	tmpl         *template.Template
	UserHandler  *handler.UserHandler
	LogedInAdmin *entity.Admin
	csrfSignKey  []byte
}

// NewAdminCategoryHandler initializes and returns new AdminCateogryHandler
func NewAdminTempHandler(T *template.Template, userHandler *handler.UserHandler, csKey []byte) *AdminTempHandler {
	return &AdminTempHandler{
		tmpl:        T,
		UserHandler: userHandler,
		csrfSignKey: csKey,
	}
}

func (ach *AdminTempHandler) AdminIndex(w http.ResponseWriter, r *http.Request) {
	ach.tmpl.ExecuteTemplate(w, "admin.home.layout", nil)

}
func (ach *AdminTempHandler) DoctorTempHandler(w http.ResponseWriter, r *http.Request) {
	doctors, _ := Admin.Doctors()
	ach.tmpl.ExecuteTemplate(w, "admin.manage.doctors.layout", doctors)

}
func (ach *AdminTempHandler) AddDoctorTempHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENtrting adding")
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		doctorRegistrationForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		doctorRegistrationForm.Required("full_name", "phone", "email", "user_password", "confirm_password")
		doctorRegistrationForm.MatchesPattern("email", form.EmailRX)
		doctorRegistrationForm.MatchesPattern("phone", form.PhoneRX)
		doctorRegistrationForm.MinLength("user_password", 8)
		doctorRegistrationForm.PasswordMatches("user_password", "confirm_password")
		doctorRegistrationForm.CSRF = token

		if !doctorRegistrationForm.Valid() {
			ach.tmpl.ExecuteTemplate(w, "admin.new.doctors.layout", doctorRegistrationForm)
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			fmt.Println("hash password")
			// doctorRegistrationForm.VErrors.Add("password", "Password Could not be stored")
			ach.tmpl.ExecuteTemplate(w, "admin.new.doctors.layout", doctorRegistrationForm)
			return
		}

		med := entity.Doctor{}

		med.User.FullName = r.FormValue("full_name")
		med.User.BirthDate = time.Now()
		med.User.Password = string(hashedPassword)
		med.User.Email = r.FormValue("email")
		med.User.Phone = r.FormValue("phone")
		med.User.Address = r.FormValue("address")
		med.User.Sex = r.FormValue("sex")
		med.User.RoleId = 3
		med.User.Description = r.FormValue("description")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()
		med.User.Image = fh.Filename

		Admin.WriteFile(&mf, fh.Filename)
		med.Department = r.FormValue("deparment")
		fmt.Println("Posting the user and jfgjfjh")
		Admin.PostDoctor(&med)
		http.Redirect(w, r, "/admin/doctors", http.StatusSeeOther)

	} else {
		if r.Method == http.MethodGet {
			doctorAddForm := struct {
				Values  url.Values
				VErrors form.ValidationErrors
				CSRF    string
			}{
				Values:  nil,
				VErrors: nil,
				CSRF:    token,
			}
			ach.tmpl.ExecuteTemplate(w, "admin.new.doctors.layout", doctorAddForm)

		}
		fmt.Println("temlating")
		fmt.Println("temlating2")

	}
}
func (ach *AdminTempHandler) UpdateDoctorTempHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		parms, _ := Admin.GetDoctor(uint(id))
		fmt.Println("updateing", parms)

		ach.tmpl.ExecuteTemplate(w, "admin.update.manage.doctors.layout", parms)

	} else if r.Method == http.MethodPost {
		token, err := rtoken.CSRFToken(ach.csrfSignKey)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		if r.Method == http.MethodPost {

			err := r.ParseForm()
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			doctorRegistrationForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
			doctorRegistrationForm.Required("full_name", "phone", "email", "user_password", "confirm_password")
			doctorRegistrationForm.MatchesPattern("email", form.EmailRX)
			doctorRegistrationForm.MatchesPattern("phone", form.PhoneRX)
			doctorRegistrationForm.MinLength("user_password", 8)
			doctorRegistrationForm.PasswordMatches("user_password", "confirm_password")
			doctorRegistrationForm.CSRF = token

			if !doctorRegistrationForm.Valid() {
				ach.tmpl.ExecuteTemplate(w, "admin.new.doctors.layout", doctorRegistrationForm)
				return
			}
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
			if err != nil {
				fmt.Println("hash password")
				doctorRegistrationForm.VErrors.Add("user_password", "Password Could not be stored")
				ach.tmpl.ExecuteTemplate(w, "admin.new.doctors.layout", doctorRegistrationForm)
				return
			}

			id, _ := strconv.Atoi(r.FormValue("id"))
			fmt.Println("id", id)
			medicine, _ := Admin.GetDoctor(uint(id))
			fmt.Println("medinen", medicine)

			med := entity.Doctor{}

			med.User.ID = medicine.User.ID
			fname := r.FormValue("fullname")

			med.User.FullName = fname
			med.User.BirthDate = time.Now()

			med.User.Password = string(hashedPassword)
			med.User.Email = r.FormValue("email")
			med.User.Phone = r.FormValue("phone")
			med.User.Address = r.FormValue("address")
			med.User.Sex = r.FormValue("sex")
			med.User.RoleId = 3
			med.User.Description = r.FormValue("description")
			var image, _, _ = r.FormFile("image")
			if image == nil {
				med.User.Image = r.FormValue("hidimage")
			} else {
				mf, fh, err := r.FormFile("image")
				if err != nil {
					panic(err)
				}
				defer mf.Close()

				med.User.Image = fh.Filename

				Admin.WriteFile(&mf, fh.Filename)

			}
			med.Department = r.FormValue("department")
			fmt.Println("TO be done", med)
			Admin.PutDoctor(&med)

			http.Redirect(w, r, "/adminDoctors", http.StatusSeeOther)

		} else {
			http.Redirect(w, r, "/adminDoctors", http.StatusSeeOther)
		}
	}
}

func (ach *AdminTempHandler) DeleteDoctorTempHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}
		fmt.Println(id, "ifhdsjkhfjdh")

		Admin.DeleteDoctor(uint(id))

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/adminDoctors", http.StatusSeeOther)
}
