package Admin_hanlder

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/web1_group_project/hospital_client/data/Admin"
	"github.com/web1_group_project/hospital_client/entity"
	"github.com/web1_group_project/hospital_client/form"
	"github.com/web1_group_project/hospital_client/rtoken"
	"golang.org/x/crypto/bcrypt"
)

func (ach *AdminTempHandler) LaboratoristTempHandler(w http.ResponseWriter, r *http.Request) {
	pharmacists, _ := Admin.Laboratorist()
	ach.tmpl.ExecuteTemplate(w, "admin.manage.laboratorists.layout", pharmacists)
}
func (ach *AdminTempHandler) LaboratoristNewTempHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENtrting adding")
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodPost {
		fmt.Println("Posting the method")
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
			ach.tmpl.ExecuteTemplate(w, "admin.new.manage.laboratorists.layout", doctorRegistrationForm)
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			fmt.Println("hash password")
			// doctorRegistrationForm.VErrors.Add("password", "Password Could not be stored")
			ach.tmpl.ExecuteTemplate(w, "admin.new.manage.laboratorists.layout.layout", doctorRegistrationForm)
			return
		}

		med := entity.Laboratorist{}

		med.User.FullName = r.FormValue("full_name")
		med.User.BirthDate = time.Now()
		med.User.Password = string(hashedPassword)
		med.User.Email = r.FormValue("email")
		med.User.Phone = r.FormValue("phone")
		med.User.Address = r.FormValue("address")
		med.User.Sex = r.FormValue("sex")
		med.User.RoleId = 5
		med.User.Description = r.FormValue("description")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()
		med.User.Image = fh.Filename

		Admin.WriteFile(&mf, fh.Filename)
		fmt.Println("Posting the user and jfgjfjh")
		Admin.PostLaboratorist(&med)
		http.Redirect(w, r, "/admin/laboratorists", http.StatusSeeOther)

	} else {
		if r.Method == http.MethodGet {
			fmt.Println("temlating")

			LaborAddForm := struct {
				Values  url.Values
				VErrors form.ValidationErrors
				CSRF    string
			}{
				Values:  nil,
				VErrors: nil,
				CSRF:    token,
			}
			ach.tmpl.ExecuteTemplate(w, "admin.new.manage.laboratorists.layout", LaborAddForm)

		}
		fmt.Println("temlating")
		fmt.Println("temlating2")
	}

}

func (ach *AdminTempHandler) UpdateLLaboratoristTempHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		parms, _ := Admin.GetLaboratorist(uint(id))
		fmt.Println("updateing", parms)

		ach.tmpl.ExecuteTemplate(w, "admin.update.manage.laboratorist.layout", parms)

	} else if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.FormValue("id"))
		fmt.Println("id", id)
		medicine, _ := Admin.GetPharmacist(uint(id))
		fmt.Println("medinen", medicine)

		med := entity.Laboratorist{}

		med.User.ID = medicine.User.ID
		fname := r.FormValue("fullname")

		med.User.FullName = fname
		med.User.BirthDate = time.Now()
		if r.FormValue("password") == r.FormValue("conpassword") {
			med.User.Password = r.FormValue("password")

		} else {
			ach.tmpl.ExecuteTemplate(w, "admin.update.laboratorist.layout", nil)
		}
		med.User.Password = r.FormValue("password")
		med.User.Email = r.FormValue("email")
		med.User.Phone = r.FormValue("phone")
		med.User.Address = r.FormValue("address")
		med.User.Sex = r.FormValue("sex")
		med.User.RoleId = 4
		med.User.Description = r.FormValue("description")
		image, _, _ := r.FormFile("image")
		if image == nil {
			med.User.Image = r.FormValue("hidimage")
			fmt.Println("Image is null")
		} else {
			mf, fh, err := r.FormFile("image")
			if err != nil {
				panic(err)
			}
			defer mf.Close()

			med.User.Image = fh.Filename

			Admin.WriteFile(&mf, fh.Filename)

		}
		fmt.Println("TO be done", med)
		Admin.PutLaboratorist(&med)

		http.Redirect(w, r, "/adminLaboratorists", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/adminLaboratorists", http.StatusSeeOther)
	}
}

func (ach *AdminTempHandler) DeleteLaboratoristTempHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}
		fmt.Println(id, "ifhdsjkhfjdh")

		Admin.DeleteLaboratorist(uint(id))

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/adminLaboratorists", http.StatusSeeOther)
}
